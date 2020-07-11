package controllers

import (
	"lyx-blog/models"
	"lyx-blog/syserrors"
)

type MessageController struct {
	BaseController
}

///message
// @router /new/?:key [post]
func (this *MessageController) NewMessage() {
	this.MustLogin()
	noteKey := this.Ctx.Input.Param(":key")
	content := this.GetString("content", "请输入内容!")
	key := this.UUID()
	message := &models.Message{
		Key:     key,
		NoteKey: noteKey,
		User:    this.User,
		UserID:  int(this.User.ID),
		Content: content,
	}
	err := models.SaveMessage(message)
	if err != nil {
		this.Abort500(syserrors.NewError("保存失败!", err))
	}
	this.JSONOkH("保存成功!", H{"data": message})
}

// @router /count [get]
func (this *MessageController) Count() {
	count, err := models.QueryMessageForNoteCount("")
	if err != nil {
		this.Abort500(syserrors.NewError("查询失败", err))
	}
	this.JSONOkH("查询成功！", H{
		"count": count,
	})
}

// @router /query [get]
func (this *MessageController) Query() {
	pageno, err := this.GetInt("pageno", 1)
	if err != nil || pageno < 1 {
		pageno = 1
	}
	limit, err := this.GetInt("limit", 10)
	if err != nil || limit < 5 {
		limit = 10
	}

	datas, err := models.QueryMessageForNoteByPage("", pageno, limit)
	if err != nil {
		this.Abort500(syserrors.NewError("查询失败", err))
	}
	this.JSONOkH("查询成功！", H{
		"data": datas,
	})
}
