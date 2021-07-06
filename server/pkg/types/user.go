package types

type User struct {
	Uid        string `json:"uid,string"`
	Name       string `json:"name,string"`
	About      string `json:"about,string"`
	Email      string `json:"email,string"`
	Password   string `json:"password,string"`
	ProfilePic string `json:"profile_pic,string"`
}
