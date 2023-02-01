package errorcodes

// 用户名被占用
var ErrorUsernameOccupiedCode = 10100
var ErrorUsernameOccupiedMsg = "occupied user name"

// 登录失败
var ErrorUserInfoWrongCode = 10200
var ErrorUserInfoWrongMsg = "login failed, check password and username again"

// 修改密码失败, 原密码错误
var ErrorWrongOldPasswordCode = 10300
var ErrorWrongOldPasswordMsg = "wrong old password"
