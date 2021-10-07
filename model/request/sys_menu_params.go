package request

type SysBaseMenuParams struct {
	MenuName       string `json:"menuName"`
	MenuIcon       string `json:"menuIcon"`
	MenuPath       string `json:"menuPath"`
	MenuParentId   int    `json:"menuParentId"`
	MenuParentName string `json:"menuParentName"`
}
