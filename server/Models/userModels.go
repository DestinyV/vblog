package Models

type User struct {
	Id       int    `gorm:"primary_key;auto_increment" json:"id"`
	Name     string `gorm:"default:NULL"json:"name"`
	Password string `gorm:"default:NULL"json: "password"`
}

// 设置Test的表名为`test`
func (User) TableName() string {
	return "user"
}

func (this *Test) Login() {
	// result := Mysql.DB.Create(&this)
	// id = this.Id
	// if result.Error != nil {
	// 	err = result.Error
	// 	return
	// }
	// return
}
