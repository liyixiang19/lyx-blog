package controllers

import (
	"bytes"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/jinzhu/gorm"
	"lyx-blog/models"
	"lyx-blog/syserrors"
)

type NoteController struct {
	BaseController
}


///note
// @router /new [get]
func (this *NoteController)Index()  {
	this.Data["key"] = this.UUID()
	this.TplName = "note_new.html"
}

func (this *NoteController) NextPrepare() {
	this.MustLogin()
	if this.User.Role  != 0 {
		this.Abort500(errors.New("权限不足"))
	}
}


///note
// @router /save/:key [post]
func (this *NoteController)Save()  {
	key := this.Ctx.Input.Param(":key")
	title := this.GetMsgString("title", "请输入标题")
	content := this.GetMsgString("content", "请输入内容")

	var n models.Note
	note, err := models.QueryNoteByKey(key)
	if err != nil {
		//如果没有该key值的文章
		if err == gorm.ErrRecordNotFound {
			n = models.Note{
				Key: key,
				Title: title,
				Content: content,
				UserID: int(this.User.ID),
				User: this.User,
			}
		}else {
			this.Abort500(syserrors.NewError("保存失败", err))
		}
	}else {
		//如果有这个文章的记录，修改文章内容
		note.Title = title
		note.Content = content
		n = *note
	}
	//摘要
	n.Summary, _ = getSummary(content)
	if err = models.SaveNote(&n); err != nil {
		this.Abort500(syserrors.NewError("保存失败", err))
	}
	this.JSONOkH("保存成功", H{})

}



func getSummary(content string) (string, error) {
	var buf bytes.Buffer
	buf.Write([]byte(content))
	doc, err := goquery.NewDocumentFromReader(&buf)
	if err != nil {
		return "", err
	}
	str := doc.Find("body").Text()
	strRune := []rune(str)
	if len(strRune) > 400 {
		strRune = strRune[:400]
	}
	return string(strRune) + "...", nil
}
