package models

type User struct {
	Model
	Name   string `gorm:"unique_index" json:"name"`
	Email  string `gorm:"unique_index" json:"email"`
	Avatar string `json:"avatar"`
	Pwd    string `json:"-"`
	Role   int    `gorm:"default:0" json:"role"` // 0 管理员 1正常用户
	Editor string `json:"editor"`
}

func QueryByEmailAndPwd(email, pwd string) (user *User, err error) {
	user = new(User)
	return user, db.Where("email=? and pwd=?", email, pwd).Take(&user).Error
}

func QueryUserByName(name string) (user *User, err error) {
	user = new(User)
	return user, db.Where("name= ?", name).Take(&user).Error
}

func QueryUserByEmail(email string) (user *User, err error) {
	user = new(User)
	return user, db.Where("email= ?", email).Take(&user).Error
}

func SaveUser(user *User) error {
	return db.Save(user).Error
}
