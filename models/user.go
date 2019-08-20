package models

type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type JwtToken struct {
	Token string `json:"token"`
}

var admin *User

func Register(user *User) {
	admin = user
}

func Login(user *User) bool {
	if admin == nil {
		return false
	}
	return admin.UserName == user.UserName && admin.Password == user.Password
}
