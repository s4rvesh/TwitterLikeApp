package model

type Users struct {
	Name     string `json:"Name"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type Tweets struct {
	Tid  int
	Uid  int
	Data string
}

//Followers data
type Followers struct {
	Uid int
	Fid int
}
