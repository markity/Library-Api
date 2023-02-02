package comment

import (
	"database/sql"
	"library-api/dao"
)

func TryPraise(userID int64, commentID int64) (exists bool, ok bool, err error) {
	tx, err := dao.DB.Begin()
	if err != nil {
		return false, false, err
	}

	// 拿表锁
	row := tx.QueryRow("SELECT id FROM table_lock WHERE tablename='rela_praise_user_and_comment' FOR UPDATE")
	var discard int64
	err = row.Scan(&discard)
	if err != nil {
		tx.Rollback()
		return false, false, err
	}

	row = tx.QueryRow("SELECT praise_cnt FROM book WHERE id=? FOR UPDATE", commentID)
	praise := int64(0)
	err = row.Scan(&praise)
	if err != nil {
		if err == sql.ErrNoRows {
			tx.Commit()
			return false, false, nil
		}
		tx.Rollback()
		return false, false, err
	}

	r := tx.QueryRow("SELECT id FROM rela_praise_user_and_comment WHERE comment_id=? AND user_id=?", commentID, userID)
	if err := r.Scan(&discard); err != nil {
		if err == sql.ErrNoRows {
			_, err = tx.Exec("INSERT INTO rela_praise_user_and_comment(user_id, comment_id) VALUES(?,?)", userID, commentID)
			if err != nil {
				tx.Rollback()
				return false, false, err
			}
			_, err = tx.Exec("UPDATE comment SET praise_cnt=? WHERE id=?", praise+1, commentID)
			if err != nil {
				tx.Rollback()
				return false, false, err
			}
			if err != nil {
				tx.Rollback()
				return false, false, err
			}
			tx.Commit()
			return true, true, err
		} else {
			tx.Rollback()
			return false, false, err
		}
	}

	tx.Commit()
	return true, false, nil
}
