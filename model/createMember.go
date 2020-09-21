package model

import "RandomChatting_Server/database"

// CreateMember - 멤버 생성
func CreateMember(name, pw string) error {
	Member := &database.Member{Name: name, Pw: pw}
	err := database.DB.Create(Member).Error
	return err
}
