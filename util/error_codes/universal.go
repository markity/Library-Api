package errorcodes

// 此文件定义了通用错误100xx

type BasicErrorResp struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
}

// 成功
var ErrorOKCode = 10000
var ErrorOKMsg = "success"

// 服务暂时不可用
var ErrorServiceNotAvailabelCode = 10001
var ErrorServiceNotAvailabelMsg = "service not available temporarily"

// 鉴权失败
var ErrorInvalidUserTokenCode = 10002
var ErrorInvalidUserTokenMsg = "invalid identity token"

// 不合法的入参
var ErrorInvalidInputParametersCode = 10003
var ErrorInvalidInputParametersMsg = "invalid parameters"
