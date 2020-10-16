package model

import "RandomChatting_Server/database"

// CheckDupID - 멤버 중복 확인
func CheckDupID(id string) error {
	Member := &database.Member{}
	err := database.DB.Where("ID = ?", id).Find(Member).Error
	return err
}
