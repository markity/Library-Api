package book

import (
	"database/sql"
	"library-api/dao"
)

func TryPraise(userID int64, bookID int64) (exists bool, ok bool, err error) {
	tx, err := dao.DB.Begin()
	if err != nil {
		return false, false, err
	}

	// 拿表锁
	row := tx.QueryRow("SELECT id FROM table_lock WHERE tablename='rela_praise_user_and_book' FOR UPDATE")
	var discard int64
	err = row.Scan(&discard)
	if err != nil {
		tx.Rollback()
		return false, false, err
	}

	row = tx.QueryRow("SELECT praise_cnt FROM book WHERE id=? FOR UPDATE", bookID)
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

	r := tx.QueryRow("SELECT id FROM rela_praise_user_and_book WHERE book_id=? AND user_id=?", bookID, userID)
	if err := r.Scan(&discard); err != nil {
		if err == sql.ErrNoRows {
			_, err = tx.Exec("INSERT INTO rela_praise_user_and_book(book_id, user_id) VALUES(?,?)", bookID, userID)
			if err != nil {
				tx.Rollback()
				return false, false, err
			}
			_, err = tx.Exec("UPDATE book SET praise_cnt=? WHERE id=?", praise+1, bookID)
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
