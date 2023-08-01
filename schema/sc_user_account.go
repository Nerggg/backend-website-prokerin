package schema

type RegisterUserBodyReq struct {
	Email    string `validate:"required" json:"email"`
	NickName string `validate:"required" json:"nick_name"`
	Password string `validate:"required" json:"password"`
	Instansi string `json:"reuqired"`
}

type LoginBodyReq struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}
type SuccesLogin struct {
	AccessToken string `json:"access_token"`
}
