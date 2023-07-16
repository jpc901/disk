package model

type User struct {
	Uid 		 int64  `json:"uid"`
	Username     string	`json:"username"`
	Password	 string `json:"-"`
	SignUpAt     string	`json:"signUpAt"`
	LastActiveAt string `json:"lastActiveAt"`
	Status       int	`json:"status"`
}