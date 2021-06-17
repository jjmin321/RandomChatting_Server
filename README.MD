# 대소고 랜덤채팅 RandomChatting 

💬 대구소프트웨어고등학교 학생들이 함께 사용할 수 있는 채팅 서비스입니다

![dgchat-mockup](https://user-images.githubusercontent.com/52072077/112093760-5e3cac80-8bdd-11eb-8224-f940176b7b8e.png)

## 기술 Stack
|                      | Web     | Server        | 
|:--------------------:|:---------------:|:------------------:|
| Developer | 제정민 | 제정민       | 
| Develop Language | VueJs| Go| 
| Develop Tool     | Visual Studio Code  | Visual Studio Code | 

## 느낀점

처음으로 혼자서 만든 웹 서비스였습니다. 실제로 사용자들이 사용할 서비스라는 점을 고려하면서 개발하다보니 개발을 끝내고 나서도 추가, 보완해야할 부분들이 많았습니다.

개발부터 배포까지 하면서 힘들었지만 이 프로젝트를 진행하면서 정말 많은 기술들을 접하였고 배우게 된 것 같습니다.


## 소개

대구소프트웨어고등학교 학생들이 함께 사용할 수 있는 채팅 서비스입니다. 

채팅 서비스를 만들어보고 싶었고, 닉네임을 통한 익명성이 보장되는 전체 채팅과 랜덤채팅을 구상하였습니다.

제가 원하는 기능들을 추가하고자 개인으로 제작하게 되었습니다.

## 기능

### 채팅방 입장

<img width="1439" alt="채팅방입장" src="https://user-images.githubusercontent.com/52072077/112093738-554bdb00-8bdd-11eb-8725-20cdded7546d.png">

채팅방 입장 전 화면입니다. 

한 계정으로 다중 접속을 시도할 시 접속되어 있는 모든 계정의 연결이 종료됩니다. 

### 랜덤 채팅

<img width="1440" alt="랜덤채팅" src="https://user-images.githubusercontent.com/52072077/112093682-3fd6b100-8bdd-11eb-8758-bd0974ebbfca.png">

랜덤 채팅은 1대1 채팅으로 이루어지며 상대방이 자동으로 배정됩니다. 

왼쪽 채팅방 목록에서 클릭을 통해 랜덤 채팅방과 전체 채팅방을 이동할 수 있습니다. 

### 전체 채팅

<img width="1438" alt="전체채팅" src="https://user-images.githubusercontent.com/52072077/112093685-3fd6b100-8bdd-11eb-98d0-b98c587fe423.png">

전체 채팅은 접속되어 있는 모든 유저들과 채팅이 이루어지며 오른쪽 사람 목록을 통해 접속되어 있는 사람들의 목록을 확인할 수 있습니다.

### 추가적인 기능

<img width="1437" alt="안내사항" src="https://user-images.githubusercontent.com/52072077/112093686-41a07480-8bdd-11eb-9a79-9323ce465459.png">

채팅 길이를 100자로 제한하였으며, 도배 방지를 위해 채팅을 친 후 1초동안 채팅을 보내지 못하게 막았습니다.

또한 같은 유저가 두 번 이상 입장하는 것을 막기 위하여 다중 접속을 시도할 시 그 유저의 IP를 통해 연결을 강제로 종료시킵니다.

### 그 외

[jjmin321/RandomChatting_Web](https://github.com/jjmin321/randomchatting_web)

[도커로 배포하기 2020-12-06](https://jjmin321.github.io/development/도커로-배포하기/)


