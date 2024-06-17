package errno

const (
	SuccessCode    = 0
	ServiceErrCode = 10000 + iota
	ParamErrCode
	ReadFileErrCode
)
