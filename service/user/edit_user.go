package user

import "library-api/dao"

func TryEditUserInfo(userID int64, doAvatar bool, avartar *string, doBirthday bool, birthday *string, doEmail bool, email *string,
	doGender bool, gender *int, doIntro bool, intro *string, doPhone bool, phone *string) error {

	args := make([]interface{}, 0)
	s := "UPDATE user SET "
	i := 0
	if doAvatar {
		i++
		args = append(args, avartar)
		s += "avartar=?,"
	}
	if doBirthday {
		i++
		args = append(args, birthday)
		s += "birthday=?,"
	}
	if doEmail {
		i++
		args = append(args, email)
		s += "email=?,"
	}
	if doGender {
		i++
		args = append(args, gender)
		s += "genger=?,"
	}
	if doIntro {
		i++
		args = append(args, intro)
		s += "introduction=?,"
	}
	if doPhone {
		i++
		args = append(args, phone)
		s += "phone = ?,"
	}

	if i != 0 {
		s = s[:len(s)-1]
		s += " WHERE id = ?"
		args = append(args, userID)
	} else {
		return nil
	}

	_, err := dao.DB.Exec(s, args...)
	if err != nil {
		return err
	}

	return nil
}
