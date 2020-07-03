package main

import (
	"encoding/gob"
	"lyx-blog/models"
	_ "lyx-blog/routers"
	"github.com/astaxie/beego"
	"strings"
	_ "lyx-blog/models"
)

func main() {
	initTemplate()
	initSeesion()
	beego.Run()
}

func initSeesion() {
	gob.Register(&models.User{})
	//https://beego.me/docs/mvc/controller/session.md
	//beego.SetStaticPath("assert", "assert")
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "blog-key"
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "data/session"
}

func initTemplate() {
	beego.AddFuncMap("equrl", func(x, y string) bool{
		x1 := strings.Trim(x, "/")
		y1 := strings.Trim(y, "/")
		return strings.Compare(x1, y1) == 0
	})
}

