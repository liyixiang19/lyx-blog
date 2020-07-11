package controllers

import (
	"errors"
	"fmt"
	"lyx-blog/models"
	"lyx-blog/syserrors"
	"strings"
)

type UserController struct {
	BaseController
}

// @router /login [post]
func (this *UserController) Login() {
	email := this.GetMsgString("email", "邮箱不能空")
	pwd := this.GetMsgString("password", "密码不能空")
	//this.Ctx.WriteString("zheli")
	user, err := models.QueryByEmailAndPwd(email, pwd)
	if err != nil {
		this.Abort500(syserrors.NewError("用户名或密码错误", err))
	}
	fmt.Println(user)
	this.SetSession(SESSION_USER_KEY, user)
	this.JSONOk("登陆成功", "/")
}

// @router /logout [get]
func (this *UserController) Logout() {
	this.MustLogin()
	this.DelSession(SESSION_USER_KEY)
	this.Redirect("/", 302)
}

// @router /reg [post]
func (this *UserController) Reg() {
	name := this.GetMsgString("name", "昵称不能为空")
	email := this.GetMsgString("email", "邮箱不能为空")
	password := this.GetMsgString("password", "密码不能为空")
	password2 := this.GetMsgString("password2", "确认密码不能为空")
	if strings.Compare(password, password2) != 0 {
		this.Abort500(errors.New("两次输入的密码不一样"))
	}

	if user, err := models.QueryUserByName(name); err != nil && user.ID > 0 {
		this.Abort500(errors.New("用户昵称已存在"))
	}
	if user, err := models.QueryUserByEmail(name); err != nil && user.ID > 0 {
		this.Abort500(errors.New("用户邮箱已存在"))
	}

	if err := models.SaveUser(&models.User{
		Name:   name,
		Email:  email,
		Pwd:    password,
		Avatar: "/static/images/info-img.png",
		Role:   1,
	}); err != nil {
		this.Abort500(syserrors.NewError("用户保存失败", err))
	}
	this.JSONOk("注册成功", "/user")
}
