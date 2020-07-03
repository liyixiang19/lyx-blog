package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Note struct {
	gorm.Model
	Key     string `gorm:"unique_index;not null;"`
	UserID  int
	User    User
	Title   string `gorm:"type:text"`
	Summary string `gorm:"type:text"`
	Content string `gorm:"type:text"`
	Visit   int    `gorm:"default:0"`
	Praise  int    `gorm:"default:0"`
}

func QueryNoteByKey(key string) (note *Note, err error) {
	note = new(Note)
	err = db.Where("key = ?" , key).Take(&note).Error
	return note, err
}

func SaveNote(note *Note)error {
	return db.Save(note).Error
}

func QueryNoteByPage(page int, limit int) (note []*Note, err error) {
	return note, db.db.Model(&Note{}).Where("title like ?", fmt.Sprintf("%%%s%%", title)).Offset((page - 1)
	* limit).Limit(limit).Order("updated_at DESC").Find(&note).Error
}

