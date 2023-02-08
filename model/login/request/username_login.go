package request

type LoginUser struct {
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
}
