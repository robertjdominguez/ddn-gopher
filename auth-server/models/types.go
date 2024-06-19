package models

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Request struct {
	Token string `json:"token"`
}

type DecodedToken struct {
	Username string
	UserId   float64
	IsValid  bool
}
