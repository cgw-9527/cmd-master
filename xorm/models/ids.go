package models

type Ids struct {
	Uploadid   int `xorm:"not null INT(11)"`
	Deleteid   int `xorm:"not null INT(11)"`
	Shareid    int `xorm:"not null INT(11)"`
	Downloadid int `xorm:"not null INT(11)"`
	Blackid    int `xorm:"not null INT(11)"`
}
