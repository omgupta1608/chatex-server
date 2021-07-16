package types

type User struct {
	Uid        string `json:"uid,omitempty"`
	Name       string `json:"name,omitempty"`
	About      string `json:"about,omitempty"`
	Email      string `json:"email,omitempty"`
	Password   string `json:"password,omitempty"`
	ProfilePic string `json:"profile_pic,omitempty"`
}
