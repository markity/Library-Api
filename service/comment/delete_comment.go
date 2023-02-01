package comment

import (
	"database/sql"
	"library-api/dao"
)

func TryDeleteComment(commentID int64, userID int64) (exists bool, permissionDeny bool, err error) {
	tobeDeleted := make([]int64, 0)

	tx, err := dao.DB.Begin()
	if err != nil {
		return false, false, err
	}

	row := tx.QueryRow("SELECT id FROM comment WHERE id=? FOR UPDATE", commentID)
	var discard int64
	err = row.Scan(&discard)
	if err != nil {
		if err == sql.ErrNoRows {
			tx.Rollback()
			return false, false, nil
		} else {
			return false, false, err
		}
	}

	// 循环逐层扫描, 扫描并锁住所有相关评论
	tobeDeleted = append(tobeDeleted, commentID)
	var offset = 0
	var n = 1
	var tmp = 0
	for {
		if n == 0 {
			break
		}
		for i := 0; i < n; i++ {
			rows, err := tx.Query("SELECT id FROM comment WHERE parent=? FOR UPDATE", tobeDeleted[offset])
			if err != nil {
				tx.Rollback()
				return false, false, err
			}
			for rows.Next() {
				var id int64
				err := rows.Scan(&id)
				if err != nil {
					tx.Rollback()
					return false, false, err
				}
				tobeDeleted = append(tobeDeleted, id)
				tmp++
			}
			offset++
		}
		n = tmp
		tmp = 0
	}

	for _, v := range tobeDeleted {
		_, err := tx.Exec("DELETE FROM comment WHERE id=?", v)
		if err != nil {
			tx.Rollback()
			return false, false, err
		}
	}

	tx.Commit()
	return true, false, nil
}
