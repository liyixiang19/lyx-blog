package controllers

import "lyx-blog/models"

type IndexController struct {
	BaseController
}


// @router / [get]
func (this *IndexController)Index() {
	limit := 10
//page
	page, err := this.GetInt("page", 1)
	if err != nil || page <= 0 {
		page = 1
	}
	models.QueryNoteByPage(page, limit)
	this.TplName = "index.html"
}

// @router /message [get]
func (this *IndexController)GetMessage() {
	this.TplName = "message.html"
}

// @router /about [get]
func (this *IndexController)GetAbout() {
	this.TplName = "about.html"
}

// @router /user [get]
func (this *IndexController)GetUser() {
	this.TplName = "user.html"
}

// @router /reg [get]
func (this *IndexController)GetReg() {
	this.TplName = "register.html"
}
