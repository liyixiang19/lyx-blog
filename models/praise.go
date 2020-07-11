package models

type PraiseLog struct {
	Model
	Key    string `sql:"index"`
	UserID int    `sql:"index"`
	Type   string `sql:"index"`
	Flag   bool
}

func UpdateNote4Praise(n *Note) error {
	return db.Model(&Note{}).Where("id = ?", n.ID).UpdateColumn("praise", n.Praise).Error
}

func QueryPraiseLog(key string, user_id int, ttype string) (parselog PraiseLog, err error) {
	return parselog, db.Where("`key` = ? and user_id =? and type = ? ", key, user_id, ttype).Take(&parselog).Error
}

func SavePraiseLog(p *PraiseLog) error {
	return db.Save(&p).Error
}
