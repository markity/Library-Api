package book

import (
	"database/sql"
	"library-api/dao"
	timeconvert "library-api/util/time_convert"
)

func TrySearchLabelBooks(userID int64, label string) ([]BookWithPariseAndFocus, error) {
	tx, err := dao.DB.Begin()
	if err != nil {
		return nil, err
	}

	list := make([]BookWithPariseAndFocus, 0)
	var publishTime_ string

	var idNow int64 = 0
	for {
		m := BookWithPariseAndFocus{}
		r := tx.QueryRow(`SELECT book.id,book.name,author,comment_num,score,publish_time,
	cover_link,content_link,label_string,praise_cnt FROM book,label,rela_book_and_label WHERE 
	rela_book_and_label.label_id=label.id AND rela_book_and_label.book_id=book.id AND label.name=? AND book.id>?`, label, idNow)
		err := r.Scan(&m.Book.ID, &m.Book.Name, &m.Book.Author, &m.Book.CommentNum, &m.Book.Score,
			&publishTime_, &m.Book.CoverLink, &m.Book.ContentLink, &m.Book.Label, &m.Book.PraiseCnt)
		if err != nil {
			if err == sql.ErrNoRows {
				return list, nil
			} else {
				return nil, err
			}
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
		list = append(list, m)
		idNow = m.Book.ID
	}

}
