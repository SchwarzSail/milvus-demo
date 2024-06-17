package errno

var (
	Success     = NewErrNo(SuccessCode, "Success")
	ServiceErr  = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	ParamErr    = NewErrNo(ParamErrCode, "Wrong Parameter has been given")
	ReadFileErr = NewErrNo(ReadFileErrCode, "Read file error")
)
