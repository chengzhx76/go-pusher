package mode

type CommonError struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
