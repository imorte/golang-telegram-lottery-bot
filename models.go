package main

type Users struct {
	Id       int    `gorm:"primary_key;column:id"`
	UserId   int    `gorm:"column:userId"`
	Username string `gorm:"type:TEXT;column:userName"`
	IsWinner bool   `gorm:"column:is_winner"`
}

func (Users) TableName() string {
	return "users"
}

type Info struct {
	Id        int    `gorm:"primary_key;column:id"`
	Admin     string `gorm:"type:TEXT;column:admin"`
	Active    bool   `gorm:"column:is_winner"`
	TimeStart string `gorm:"type:TEXT;column:time_start"`
	TimeEnd   string `gorm:"type:TEXT;column:time_end"`
}

func (Info) TableName() string {
	return "main"
}
