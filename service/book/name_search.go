package book

import (
	"database/sql"
	"library-api/dao"
	timeconvert "library-api/util/time_convert"
)

func TryGetBookByName(name string, userID int64) (*BookWithPariseAndFocus, error) {
	m := BookWithPariseAndFocus{}
	var publishTime_ string

	tx, err := dao.DB.Begin()
	if err != nil {
		return nil, err
	}

	row := tx.QueryRow(`SELECT id, name, author, comment_num, score, publish_time,
	 	cover_link, content_link, label_string,praise_cnt FROM book WHERE name=?`, name)
	if err := row.Scan(&m.Book.ID, &m.Book.Name, &m.Book.Author, &m.Book.CommentNum,
		&m.Book.Score, &publishTime_, &m.Book.ContentLink,
		&m.Book.ContentLink, &m.Book.Label, &m.Book.PraiseCnt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	if userID != -1 {
		row := tx.QueryRow("SELECT id FROM rela_focus_user_and_book WHERE user_id=? AND book_id=?", userID, m.Book.ID)
		var discard int64
		if err := row.Scan(&discard); err != nil {
			if err == sql.ErrNoRows {
				b := false
				m.Foucs = &b
			} else {
				return nil, err
			}
		} else {
			b := true
			m.Foucs = &b
		}
		row = tx.QueryRow("SELECT id FROM rela_praise_user_and_book WHERE user_id=? AND book_id=?", userID, m.Book.ID)
		if err := row.Scan(&discard); err != nil {
			if err == sql.ErrNoRows {
				b := false
				m.Praise = &b
			} else {
				return nil, err
			}
		} else {
			b := true
			m.Praise = &b
		}
	}
	m.Book.PublishTime = timeconvert.MustStrToTime(publishTime_)

	return &m, nil
}
