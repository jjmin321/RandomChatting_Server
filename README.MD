# 대소고 랜덤채팅 RandomChatting 

💬 대구소프트웨어고등학교 학생들이 함께 사용할 수 있는 채팅 서비스입니다

## 기능 Function
- 서비스 이용을 위한 로그인 및 회원가입 기능 
- 닉네임을 통한 익명성이 보장되는 전체 채팅과 랜덤 채팅 기능 
- 전체 채팅방, 랜덤 채팅방에 접속되어 있는 유저 목록 확인 기능
- 전체 채팅, 랜덤 채팅방을 클릭을 통한 스위치 기능 (연결과 채팅 기록은 유지됨)
- 원하는 채팅을 전송하고 상대방의 채팅을 볼 수 있는 기능
- 자동으로 랜덤채팅 방 배정 및 채팅방 나가기 기능 

## 기술 Stack
|                      | Web     | Server        | 
|:--------------------:|:---------------:|:------------------:|
| Developer | 제정민 | 제정민       | 
| Develop Language | VueJs| Go| 
| Develop Tool     | Visual Studio Code  | Visual Studio Code | 

## 시스템 구상도

![image](https://user-images.githubusercontent.com/52072077/97800503-0676c680-1c79-11eb-9156-7c5c303a30d5.png)

## 사용 기술

### 프론트엔드
![front-end](https://user-images.githubusercontent.com/52072077/97800589-c6fcaa00-1c79-11eb-8295-9b0450432e3a.png)

VueJS, Scss, Axios를 사용하여 프론트엔드를 제작하였습니다

### 백엔드

![back-end](https://user-images.githubusercontent.com/52072077/97800688-a1bc6b80-1c7a-11eb-8946-20e0e87b9c4b.png)

Go, Echo, PostgreSQL 등을 사용하여 백엔드를 제작하였고 GCP와 Ubuntu를 사용해 배포하였습니다.