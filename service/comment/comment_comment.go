package comment

import (
	"database/sql"
	"library-api/dao"
	"time"
)

func TryAddCommentComment(commentID int64, userID int64, content string, annoymous bool) (int64, error) {
	tx, err := dao.DB.Begin()
	if err != nil {
		return -1, err
	}

	// 确认有这个评论
	r := tx.QueryRow("SELECT book_id FROM comment WHERE id=?", commentID)
	var bookID int64
	err = r.Scan(&bookID)
	if err != nil {
		if err == sql.ErrNoRows {
			tx.Commit()
			return -1, nil
		} else {
			tx.Rollback()
			return -1, err
		}
	}

	var username string
	var avatarLink *string
	r = tx.QueryRow("SELECT username,avatar_link FROM user WHERE id = ?", userID)
	err = r.Scan(&username, &avatarLink)
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	result, err := tx.Exec(`INSERT INTO comment(book_id,user_id,username,avatar_link,praise_cnt,content,
		parent,annoymous,publish_time,created_at) VALUES (?,?,?,?,?,?,?,?,?,?)`,
		bookID, userID, username, avatarLink, 0, content, commentID, annoymous, time.Now(), time.Now())
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	last, _ := result.LastInsertId()
	return last, nil
}
