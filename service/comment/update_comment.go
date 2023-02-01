package comment

import (
	"database/sql"
	"library-api/dao"
)

func TryUpdateComment(userID int64, commendID int64,
	updateAnnoymous bool, annoymous bool,
	updateContent bool, content string) (exists bool, permissionDenied bool, err error) {
	tx, err := dao.DB.Begin()
	if err != nil {
		return false, false, err
	}

	// 先锁住, 查看是否有此条目
	r := tx.QueryRow("SELECT user_id FROM comment WHERE id=? FOR UPDATE", commendID)
	var commentUserID int64
	err = r.Scan(&commentUserID)
	if err != nil {
		// 根本没找到这个条目
		if err == sql.ErrNoRows {
			tx.Commit()
			return false, false, nil
		} else {
			tx.Rollback()
			return false, false, err
		}
	}

	// 没权限
	if commentUserID != userID {
		tx.Commit()
		return true, false, nil
	}

	// 插入
	args := make([]interface{}, 0)
	s := "UPDATE comment SET "
	if updateAnnoymous {
		args = append(args, annoymous)
		s += "annoymous=?,"
	}
	if updateContent {
		args = append(args, content)
		s += "content=?,"
	}
	s = s[:len(s)-1]
	s += " WHERE id=?"
	args = append(args, commendID)
	_, err = tx.Exec(s, args...)
	if err != nil {
		tx.Rollback()
		return false, false, err
	}

	err = tx.Commit()
	if err != nil {
		return false, false, err
	}

	return true, false, nil

}
