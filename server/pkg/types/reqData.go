package types

type LoginReqData struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=40"`
}

type RegisterReqData struct {
	Name     string `json:"name" validate:"required,min=3,max=15"`
	About    string `json:"about" validate:"required,min=3,max=40"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=40"`
}

type UserVerificationReqData struct {
	Uid              string `json:"uid" validate:"required,len=20"`
	VerificationCode string `json:"verification_code" validate:"required,len=6"`
}

type EditUserProfileReqData struct {
	Name       string `json:"name,string"`
	About      string `json:"about,string"`
	ProfilePic string `json:"profile_pic,string"`
}

type ChangePasswordReqData struct {
	OldPassword string `json:"old_password,string"`
	NewPassword string `json:"new_password,string"`
}
