package model

import "RandomChatting_Server/database"

// FindMember - 멤버 존재 여부 확인
func FindMember(name string) error {
	Member := &database.Member{}
	err := database.DB.Where("Name = ?", name).Find(Member).Error
	return err
}
