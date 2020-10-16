package model

import "RandomChatting_Server/database"

// CreateMember - 멤버 생성
func CreateMember(id, name, pw string) error {
	Member := &database.Member{ID: id, Name: name, Pw: pw}
	err := database.DB.Create(Member).Error
	return err
}
