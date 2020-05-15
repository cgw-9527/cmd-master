package models

type Contentfilerecord struct {
	Owner    string `xorm:"not null VARCHAR(255)"`
	Hash     string `xorm:"not null VARCHAR(255)"`
	Filesize string `xorm:"not null VARCHAR(255)"`
	Time     string `xorm:"not null VARCHAR(255)"`
	Action   string `xorm:"not null VARCHAR(255)"`
}
