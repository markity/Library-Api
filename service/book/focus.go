package book

import (
	"database/sql"
	"library-api/dao"
)

func TryFocus(userID int64, bookID int64) (bool, error) {
	tx, err := dao.DB.Begin()
	if err != nil {
		return false, err
	}

	// 拿表锁
	_, err = tx.Exec("SELECT id FROM table_lock WHERE tablename='rela_focus_user_and_book' FOR UPDATE")
	if err != nil {
		tx.Rollback()
		return false, err
	}

	r := tx.QueryRow("SELECT id FROM rela_focus_user_and_book WHERE book_id=? AND user_id=?", bookID, userID)
	var discard int64
	if err := r.Scan(&discard); err != nil {
		if err == sql.ErrNoRows {
			_, err = tx.Exec("INSERT INTO rela_focus_user_and_book(book_id, user_id) VALUES(?,?)", userID, bookID)
			if err != nil {
				tx.Rollback()
				return false, err
			}
			tx.Commit()
			return true, nil
		} else {
			tx.Rollback()
			return false, err
		}
	}

	tx.Commit()
	return false, nil
}
