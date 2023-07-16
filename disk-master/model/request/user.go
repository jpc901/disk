package request

type UserSignUpRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type UserSignInRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type UserInfoRequest struct {
	Username string `form:"username" binding:"required"`
}