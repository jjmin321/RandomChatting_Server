package model

import "RandomChatting_Server/database"

// CheckDupName - 멤버 중복 확인
func CheckDupName(name string) error {
	Member := &database.Member{}
	err := database.DB.Where("Name = ?", name).Find(Member).Error
	return err
}
