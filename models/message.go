package models

type Message struct {
	Model
	Key    string `gorm:"unique_index;not null;" json:"key"`
	UserID int    `json:"user_id"`
	User   User   `json:"user"`
	//NoteID  int
	//Note    Note
	NoteKey string `sql:"index" json:"note_key"`
	Content string `json:"content"`
	Praise  int    `gorm:"default:0" json:"praise"`
}

func QueryMessageByKey(key string) (message Message, err error) {
	return message, db.Model(&Message{}).Where("`key` = ? ", key).Take(&message).Error
}

func QueryMessagesByNoteKey(noteKey string) (messages []*Message, err error) {
	return messages, db.Preload("User").Where("note_key = ? ", noteKey).Order("updated_at desc").Find(&messages).Error
}

func SaveMessage(message *Message) error {
	return db.Save(message).Error
}

func QueryMessageForNoteCount(key string) (count int, err error) {
	return count, db.Model(&Message{}).Where("note_key = ?", key).Count(&count).Error
}
func QueryMessageForNoteByPage(key string, page, limit int) (messages []*Message, err error) {
	return messages, db.Preload("User").Where("note_key = ?", key).Offset((page - 1) * limit).Limit(limit).Order("updated_at desc").Find(&messages).Error
}

func UpdateMessage4Praise(n *Message) error {
	return db.Model(&Message{}).Where("id = ?", n.ID).UpdateColumn("praise", n.Praise).Error
}
