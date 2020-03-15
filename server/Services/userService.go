package Services

import (
	"server/Models"
)

type User struct {
	Username string `json:"username"`
	Password string `json: "password"`
}

func (this *User) Login() (u *Models.User, err error) {
	user := &Models.User{
		Name:     this.Username,
		Password: this.Password,
	}
	u, err = user.Login()
	return
}
