package request

type Login struct {
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
	CodeAuth string `json:"codeAuth"`
	Code     string `json:"code"`
}

type SystemUserParams struct {
	Id         uint     `json:"id"`
	UserName   string   `json:"userName"`
	PassWord   string   `json:"passWord"`
	UserAvatar string   `json:"userAvatar"`
	NikeName   string   `json:"nikeName"`
	RoleId     string   `json:"roleId"`
	RoleIds    []string `json:"roleIds"`
}

type SystemUserRoleParams struct {
	UserId  uint     `json:"id" binding:"required"`
	RoleIds []string `json:"roleIds" binding:"required"`
}
