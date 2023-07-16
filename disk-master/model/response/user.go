package response

import "disk-master/model"

type UserLogin struct {
	Token string `json:"token"`
	model.User
}