package book

import (
	"database/sql"
	"library-api/dao"
	"library-api/model"
	timeconvert "library-api/util/time_convert"
)

type BookWithPariseAndFocus struct {
	Book   model.Book `json:"book"`
	Praise *bool      `json:"praise,omitempty"`
	Foucs  *bool      `json:"focus,omitempty"`
}

// 如果userID == -1 时, 不查询用户的priase和focus
func TryGetBooksWithPraiseAndFocus(userID int64) ([]BookWithPariseAndFocus, error) {
	list := make([]BookWithPariseAndFocus, 0)

	tx, err := dao.DB.Begin()
	if err != nil {
		return nil, err
	}

	var idNow = int64(0)

	var publishTime_ string
	for {
		m := BookWithPariseAndFocus{}
		row := tx.QueryRow(`SELECT id, name, author, comment_num, score, publish_time,
	 	cover_link, content_link, label_string,praise_cnt FROM book WHERE id > ?`, idNow)
		if err := row.Scan(&m.Book.ID, &m.Book.Name, &m.Book.Author, &m.Book.CommentNum,
			&m.Book.Score, &publishTime_, &m.Book.CoverLink,
			&m.Book.ContentLink, &m.Book.Label, &m.Book.PraiseCnt); err != nil {
			if err == sql.ErrNoRows {
				tx.Commit()
				return list, nil
			}
			return nil, err
		}
		idNow = m.Book.ID
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
				println("fuck")
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
		println("HEHE")
		m.Book.PublishTime = timeconvert.MustStrToTime(publishTime_)
		list = append(list, m)
	}
}
