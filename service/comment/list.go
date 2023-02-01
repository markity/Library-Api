package comment

import (
	"database/sql"
	"library-api/dao"
	"library-api/model"
	timeconvert "library-api/util/time_convert"
)

func TryListCommentsUnderBook(userID int64, bookID int64) (res []*model.CommentWithParise, exists bool, err error) {
	res = make([]*model.CommentWithParise, 0)
	tx, err := dao.DB.Begin()
	if err != nil {
		return nil, false, err
	}

	// 先查看书是否存在
	r := tx.QueryRow("SELECT id FROM book WHERE id=?", bookID)
	var discard int64
	err = r.Scan(&discard)
	if err != nil {
		if err == sql.ErrNoRows {
			tx.Commit()
			return nil, false, nil
		} else {
			tx.Rollback()
			return nil, false, err
		}
	}

	// 先扫顶层评论
	var idNow = int64(0)
	for {
		m := model.CommentWithParise{}
		var publishTime_ string
		row := tx.QueryRow(`SELECT id,book_id,user_id,username,avatar_link,praise_cnt,content,
			parent,annoymous,publish_time FROM comment WHERE parent IS NULL AND book_id=? AND id>?`, bookID, idNow)
		err := row.Scan(&m.Comment.ID, &m.Comment.BookID, &m.Comment.SenderUserID, &m.Comment.Username, &m.Comment.AvatarLink,
			&m.Comment.PraiseCnt, &m.Comment.Content, &m.Comment.Parent,
			&m.Comment.Anonymous, &publishTime_)
		if err != nil {
			if err == sql.ErrNoRows {
				break
			}
			tx.Rollback()
			return nil, false, err
		}

		m.Comment.PublishTime = timeconvert.MustStrToTime(publishTime_)

		// 如果是匿名字段, 就隐藏用户名
		if m.Comment.Anonymous {
			m.Comment.Username = nil
		}

		// 如果用户登录则补全praise字段
		if userID != -1 {
			row = tx.QueryRow("SELECT id FROM rela_praise_user_and_comment WHERE user_id=? AND comment_id=?", userID, m.Comment.ID)
			if err := row.Scan(&discard); err != nil {
				if err == sql.ErrNoRows {
					b := false
					m.Praise = &b
				} else {
					tx.Rollback()
					return nil, false, err
				}
			} else {
				b := true
				m.Praise = &b
			}
		} else {
			m.Praise = nil
		}

		idNow = m.Comment.ID
		m.Comment.SonComments = make([]*model.CommentWithParise, 0)

		err = toolScanSonComments(tx, &m, userID)
		if err != nil {
			tx.Rollback()
			return nil, false, err
		}
		res = append(res, &m)
	}

	tx.Commit()
	return res, true, nil
}

func toolScanSonComments(tx *sql.Tx, msg *model.CommentWithParise, userID int64) error {
	var idNow int64
	for {
		m := model.CommentWithParise{}
		var publishTime_ string
		row := tx.QueryRow(`SELECT id,book_id,user_id,username,avatar_link,praise_cnt,content,
			parent,annoymous,publish_time FROM comment WHERE parent=? AND id>?`, msg.Comment.ID, idNow)
		err := row.Scan(&m.Comment.ID, &m.Comment.BookID, &m.Comment.SenderUserID, &m.Comment.Username, &m.Comment.AvatarLink,
			&m.Comment.PraiseCnt, &m.Comment.Content, &m.Comment.Parent,
			&m.Comment.Anonymous, &publishTime_)
		if err != nil {
			if err == sql.ErrNoRows {
				break
			}
			tx.Rollback()
			return err
		}

		m.Comment.PublishTime = timeconvert.MustStrToTime(publishTime_)

		// 如果是匿名字段, 就隐藏用户名
		if m.Comment.Anonymous {
			m.Comment.Username = nil
		}

		// 如果用户登录则补全praise字段
		if userID != -1 {
			row = tx.QueryRow("SELECT id FROM rela_praise_user_and_comment WHERE user_id=? AND comment_id=?", userID, m.Comment.ID)
			var discard int64
			if err := row.Scan(&discard); err != nil {
				if err == sql.ErrNoRows {
					b := false
					m.Praise = &b
				} else {
					tx.Rollback()
					return err
				}
			} else {
				b := true
				m.Praise = &b
			}
		} else {
			m.Praise = nil
		}

		idNow = m.Comment.ID
		m.Comment.SonComments = make([]*model.CommentWithParise, 0)
		msg.Comment.SonComments = append(msg.Comment.SonComments, &m)
	}

	// 进入递归, 开始扫更子层的东西
	for _, v := range msg.Comment.SonComments {
		err := toolScanSonComments(tx, v, userID)
		if err != nil {
			return err
		}
	}

	return nil
}
