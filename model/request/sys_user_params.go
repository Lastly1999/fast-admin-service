package request

type Login struct {
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
	CodeAuth string `json:"codeAuth"`
	Code     string `json:"code"`
}
