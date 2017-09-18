package main

type User struct {
	Id       int    `gorm:"primary_key;column:id"`
	UserId   int    `gorm:"column:userId"`
	Username string `gorm:"type:TEXT;column:userName"`
	IsWinner bool   `gorm:"column:is_winner"`
}

func (User) TableName() string {
	return "users"
}

type Info struct {
	Id      int    `gorm:"primary_key;column:id"`
	Admin   string `gorm:"type:TEXT;column:admin"`
	Active  bool   `gorm:"column:active"`
	IsReady bool   `gorm:"column:is_ready"`
}

func (Info) TableName() string {
	return "main"
}

type Sequence struct {
	Name int `gorm:"column:name"`
	Seq  int `gorm:"column:seq"`
}

func (Sequence) TableName() string {
	return "sqlite_sequence"
}
