package request

type SysRoleParams struct {
	RoleId   string `json:"roleId"`
	RoleName string `json:"roleName"`
	Status   *bool  `json:"status"`
	Describe string `json:"describe"`
}

type SysRoleMenuParams struct {
	RoeId        string `json:"roleId"`
	PermissionId []uint `json:"permissionId"`
}

type SysRoleDefaultParams struct {
	RoleId string `json:"roleId"`
	UserId int    `json:"userId"`
}
