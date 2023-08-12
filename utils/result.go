package utils

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Success 成功
func (ctx *Result) Success(data interface{}) *Result {
	ctx.Code = 200
	ctx.Message = "SUCCESS"
	ctx.Data = data
	return ctx
}

// Error 错误
func (ctx *Result) Error(code int, message string) *Result {
	ctx.Code = code
	ctx.Message = message
	return ctx
}
