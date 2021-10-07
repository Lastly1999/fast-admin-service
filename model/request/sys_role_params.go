package request

type SysRoleParams struct {
	RoleId   uint   `json:"roleId"`
	RoleName string `json:"roleName"`
	Status   *bool  `json:"status"`
	Describe string `json:"describe"`
}

type SysRoleMenuParams struct {
	PermissionId []uint `json:"permissionId"`
}
