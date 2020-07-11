package controllers

import (
	"lyx-blog/models"
	"lyx-blog/syserrors"
)

type PraiseController struct {
	BaseController
}

func (this *PraiseController) NestPrepare() {
	this.MustLogin()
}

// @router /:type/:key [post]
func (this *PraiseController) Parse() {
	key := this.Ctx.Input.Param(":key")
	ttype := this.Ctx.Input.Param(":type")
	var (
		praise  int = 0
		user_id int = int(this.User.ID)
		err     error
	)
	this.Dao.Begin()
	switch ttype {
	case "message":
		var message models.Message
		if message, err = models.QueryMessageByKey(key); err != nil {
			this.Dao.Rollback()
			this.Abort500(syserrors.NewError("点赞失败", err))
		}
		message.Praise = message.Praise + 1
		if err := models.UpdateMessage4Praise(&message); err != nil {
			this.Dao.Rollback()
			this.Abort500(syserrors.NewError("点赞失败", err))
		}
		praise = message.Praise
	case "note":
		var note *models.Note
		if note, err = models.QueryNoteByKey(key); err != nil {
			this.Dao.Rollback()
			this.Abort500(syserrors.NewError("点赞失败", err))
		}
		note.Praise = note.Praise + 1
		if err := models.UpdateNote4Praise(note); err != nil {
			this.Dao.Rollback()
			this.Abort500(syserrors.NewError("点赞失败", err))
		}
		praise = note.Praise
	default:
		this.Dao.Rollback()
		this.Abort500(syserrors.NewError("未知类型", nil))
	}

	p := models.PraiseLog{
		Key:    key,
		UserID: user_id,
		Type:   ttype,
	}
	var pp models.PraiseLog
	if pp, err = models.QueryPraiseLog(key, user_id, ttype); err != nil {
		pp = p
	} else {
		if pp.Flag {
			this.Dao.Rollback()
			this.Abort500(syserrors.HasPraiseError{})
		}
	}
	pp.Flag = true
	if err := models.SavePraiseLog(&pp); err != nil {
		this.Dao.Rollback()
		this.Abort500(syserrors.NewError("点赞失败", err))
	}
	this.Dao.Commit()
	this.JSONOkH("点赞成功！", H{
		"praise": praise,
	})
}
