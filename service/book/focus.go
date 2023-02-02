package book

import (
	"database/sql"
	"library-api/dao"
)

func TryFocus(userID int64, bookID int64) (exists bool, ok bool, err error) {
	tx, err := dao.DB.Begin()
	if err != nil {
		return false, false, err
	}

	// 拿表锁
	_, err = tx.Exec("SELECT id FROM table_lock WHERE tablename='rela_focus_user_and_book' FOR UPDATE")
	if err != nil {
		tx.Rollback()
		return false, false, err
	}

	row := tx.QueryRow("SELECT id FROM book WHERE id=? FOR UPDATE", bookID)
	discard := int64(0)
	err = row.Scan(&discard)
	if err != nil {
		if err == sql.ErrNoRows {
			tx.Commit()
			return false, false, nil
		}
		tx.Rollback()
		return false, false, err
	}

	r := tx.QueryRow("SELECT id FROM rela_focus_user_and_book WHERE book_id=? AND user_id=?", bookID, userID)
	if err := r.Scan(&discard); err != nil {
		if err == sql.ErrNoRows {
			_, err = tx.Exec("INSERT INTO rela_focus_user_and_book(book_id, user_id) VALUES(?,?)", userID, bookID)
			if err != nil {
				tx.Rollback()
				return false, false, err
			}
			tx.Commit()
			return true, true, nil
		} else {
			tx.Rollback()
			return false, false, err
		}
	}

	tx.Commit()
	return true, false, nil
}
