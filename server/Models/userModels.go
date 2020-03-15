package Models

import (
	"fmt"
	Mysql "server/Databases"
)

type User struct {
	Uid      int    `gorm:"primary_key;auto_increment" json:"uid"`
	Name     string `gorm:"default:NULL"json:"name"`
	Password string `gorm:"default:NULL"json: "password"`
}

// 设置Test的表名为`test`
func (User) TableName() string {
	return "users"
}

func (this *User) Login() (u *User, err error) {
	var user User
	// result := Mysql.DB.Create(&this)
	// id = this.Id
	// if result.Error != nil {
	// 	err = result.Error
	// 	return
	// }
	err = Mysql.DB.Where(&User{Name: this.Name, Password: this.Password}).First(&user).Error
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	u = &user
	return
}
