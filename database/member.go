package database

import "time"

// Member 멤버 관리 테이블
type Member struct {
	Idx       uint      `gorm:"primary_key; auto_increment:true" json:"idx"`
	UserID    string    `gorm:"type:varchar(255);unique;not null" json:"user_id"`
	Pw        string    `gorm:"type:varchar(255);not null" json:"pw"`
	Nickname  string    `gorm:"type:varchar(255);not null" json:"nickname"`
	Image     string    `gorm:"type:varchar(255);not null" json:"image"`
	Megaphone uint      `gorm:"not null" sql:"DEFAULT:0" json:"megaphone"`
	JoinedAt  time.Time `gorm:"not null" sql:"DEFAULT:current_timestamp" json:"joined_at"`
}
