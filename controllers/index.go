package controllers

import (
	"lyx-blog/models"
	"lyx-blog/syserrors"
)

type IndexController struct {
	BaseController
}

// @router / [get]
func (this *IndexController) Index() {
	limit := 3
	//page
	page, err := this.GetInt("page", 1)
	if err != nil || page <= 0 {
		page = 1
	}

	title := this.GetString("title")
	//得到页面数据
	notes, err := models.QueryNoteByPage(title, page, limit)
	if err != nil {
		this.Abort500(err)
	}
	this.Data["notes"] = notes

	//页数
	count, err := models.QueryNotesCount(title)
	if err != nil {
		this.Abort500(err)
	}
	totalpage := count / limit
	if count%limit != 0 {
		totalpage++
	}
	this.Data["totalpage"] = totalpage
	this.Data["page"] = page
	this.Data["title"] = title
	this.TplName = "index.html"
}

// @router /message [get]
func (this *IndexController) GetMessage() {
	this.TplName = "message.html"
}

// @router /about [get]
func (this *IndexController) GetAbout() {
	this.TplName = "about.html"
}

// @router /user [get]
func (this *IndexController) GetUser() {
	this.TplName = "user.html"
}

// @router /reg [get]
func (this *IndexController) GetReg() {
	this.TplName = "register.html"
}

// @router /details/:key [get]
func (this *IndexController) GetDetails() {
	key := this.Ctx.Input.Param(":key")
	note, err := models.QueryNoteByKey(key)
	if err != nil {
		this.Abort500(syserrors.NewError("文章不存在", err))
	}
	this.Data["note"] = note

	messages, err := models.QueryMessagesByNoteKey(key)
	if err != nil {
		this.Abort500(syserrors.NewError("评论不存在", err))
	}
	this.Data["messages"] = messages
	this.TplName = "details.html"
}

// @router /comment/:key [get]
func (this *IndexController) GetComment() {
	key := this.Ctx.Input.Param(":key")
	note, err := models.QueryNoteByKey(key)
	if err != nil {
		this.Abort500(syserrors.NewError("文章不存在", err))
	}
	this.Data["note"] = note
	this.TplName = "comment.html"
}
