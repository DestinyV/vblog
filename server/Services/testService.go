package Services

import (
	"server/Models"
)

type Test struct {
	Id      int    `json:"id"`
	Testcol string `json:"testcol"`
}

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json: "password"`
}

func (this *Test) Insert() (id int, err error) {
	var testModel Models.Test
	testModel.Id = this.Id
	testModel.Testcol = this.Testcol
	id, err = testModel.Insert()
	return
}

func (user *User) Login() {

}
