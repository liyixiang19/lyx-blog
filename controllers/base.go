package controllers

import (
	"errors"
	"github.com/astaxie/beego"
	uuid "github.com/satori/go.uuid"
	"lyx-blog/models"
	"lyx-blog/syserrors"
)

const SESSION_USER_KEY = "SESSION_USER_KEY"

type BaseController struct {
	beego.Controller
	User    models.User
	IsLogin bool
	Dao     *models.DB
}

type NestPrepare interface {
	NextPrepare()
}

func (this *BaseController) Prepare() {
	this.Data["Path"] = this.Ctx.Request.RequestURI
	// 验证用户是否登陆
	this.IsLogin = false
	u, ok := this.GetSession(SESSION_USER_KEY).(*models.User) //断言
	if ok {
		this.User = *u
		this.Data["User"] = *u
		this.IsLogin = true
	}
	this.Data["isLogin"] = this.IsLogin
	//判断子controller是否实现接口 NestPreparer
	if a, ok := this.AppController.(NestPrepare); ok {
		a.NextPrepare()
	}
}

func (this *BaseController) Abort500(err error) {
	this.Data["error"] = err
	this.Abort("500")
}

func (this *BaseController) GetMsgString(key string, msg string) string {
	s := this.GetString(key)
	if len(s) == 0 {
		this.Abort500(errors.New(msg))
	}
	return s
}

func (this *BaseController) MustLogin() {
	if !this.IsLogin {
		this.Abort500(syserrors.NoUserError{})
	}
}

type H map[string]interface{}

type ResultJsonValue struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Action string      `json:"action,omitempty"`
	Count  int         `json:"count,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

func (this *BaseController) JSONOk(msg string, actions ...string) {
	var action string
	if len(actions) > 0 {
		action = actions[0]
	}
	this.Data["json"] = &ResultJsonValue{
		Code:   0,
		Msg:    msg,
		Action: action,
	}
	this.ServeJSON()
}

func (this *BaseController) JSONOkH(msg string, maps H) {
	if maps == nil {
		maps = H{}
	}
	maps["code"] = 0
	maps["msg"] = msg
	this.Data["json"] = maps
	this.ServeJSON()
}

func (this *BaseController) UUID() string {
	u := uuid.NewV4()
	return u.String()
}
