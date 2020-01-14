package user

type CreateRequest struct {
	UserName    string `json:"user_name"`
	UserPwd     string `json:"user_pwd"`
	UserEmail   string `json:"user_email"`
	UserPhone   string `json:"user_phone"`
	Sex         int8   `json:"sex"`
	Address     string `json:"address"`
	IsAvailable int8   `json:"is_available"`
	Remark      string `json:"remark"`
	RoleId      uint   `json:"role_id"`
}

type CreateResponse struct {
	Username string `json:"username"`
}

type GetResponse struct {
	UserName    string `json:"user_name"`
	UserEmail   string `json:"user_email"`
	UserPhone   string `json:"user_phone"`
	Sex         int8   `json:"sex"`
	Address     string `json:"address"`
	IsAvailable int8   `json:"is_available"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	UID   uint   `json:"uid"`
	Token string `json:"token"`
}

type LoginTotalResponse struct {
	Total uint `json:"total"`
}
