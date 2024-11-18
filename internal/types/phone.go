package types

type PhoneReq struct {
	Phone     string `json:"phone"`
	Code      string `json:"code"`
	AutoLogin bool   `json:"auto_login"`
	// 不需要前端传参
	UserIP    string `json:"-"`
	UserAgent string `json:"-"`
}
type PhoneResp struct {
	Atoken    string `json:"atoken"`
	Rtoken    string `json:"rtoken"`
	LoginId   string `json:"login_id"`
	UserAgent string `json:"user_agent"`
	Ip        string `json:"ip"`
}
