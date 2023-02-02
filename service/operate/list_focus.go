package operate

import (
	"library-api/dao"
	"library-api/model"
)

func TryListFocus(userID int64) (res []*model.FocusListItem, err error) {
	res = make([]*model.FocusListItem, 0)
	rows, err := dao.DB.Query(`SELECT book.id,book.name,book.publish_time,book.content_link FROM rela_focus_user_and_book, user, book WHERE user.id = rela_focus_user_and_book.user_id AND
	rela_focus_user_and_book.book_id=book.id`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		m := model.FocusListItem{}
		var publishTime_ string
		err := rows.Scan(&m.BookID, &m.Name, &publishTime_, &m.ContentLink)
		if err != nil {
			rows.Close()
			return nil, err
		}
		res = append(res, &m)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return res, nil
}
