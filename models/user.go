package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"unique_index"`
	Email string `gorm:"unique_index"`
	Pwd string
	Avatar string
	Role int // 0代表管理员，1代表普通用户,默认值是1
}

func QueryByEmailAndPwd(email, pwd string)(user *User, err error)  {
	user = new(User)
	return user,db.Where("email=? and pwd=?", email,pwd).Take(&user).Error
}

func QueryUserByName(name string) (user *User, err error)  {
	user = new(User)
	return user, db.Where("name= ?", name).Take(&user).Error
}


func QueryUserByEmail(email string) (user *User, err error)  {
	user = new(User)
	return user, db.Where("email= ?", email).Take(&user).Error
}

func SaveUser(user *User) error {
	return db.Save(user).Error
}