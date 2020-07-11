package models

import (
	"fmt"
)

type Note struct {
	Model
	Key     string `gorm:"unique_index;not null;"`
	UserID  int
	User    User
	Title   string
	Summary string `gorm:"type:text"`
	Content string `gorm:"type:text"`
	Source  string `gorm:"type:text" json:"source"`
	Editor  string `gorm:"varchar(40)'" `
	Files   string `gorm:"type:text"`
	Visit   int    `gorm:"default:0"`
	Praise  int    `gorm:"default:0"`
}

func QueryNoteByKey(key string) (note *Note, err error) {
	note = new(Note)
	err = db.Where("key = ?", key).Take(&note).Error
	return note, err
}

func SaveNote(note *Note) error {
	return db.Save(note).Error
}

func QueryNotesCount(title string) (count int, err error) {
	return count, db.Model(&Note{}).Where("title like ?", fmt.Sprintf("%%%s%%",
		title)).Count(&count).Error
}

func QueryNoteByPage(title string, page int, limit int) (note []*Note, err error) {
	//like "%like%"
	return note, db.Where("title like ?", fmt.Sprintf("%%%s%%",
		title)).Offset((page - 1) * limit).Limit(limit).Order("updated_at DESC").Find(&note).Error
}

func QueryNoteByKeyAndUserId(key string, userid int) (note Note, err error) {
	return note, db.Where("`key` = ? and user_id = ?", key, userid).Take(&note).Error
}

func DelNoteByUserIDAndKey(key string, userid int) error {
	return db.Delete(&Note{}, "key=? and user_id = ?", key, userid).Error
}
