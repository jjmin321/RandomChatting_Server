# 대소고인을 위한 1대1 랜덤채팅 서비스

너를 만나 

### 9월 15일
웹 소켓 서버를 여는 것을 성공적으로 띄웠다.
하지만 연결하면 http: connection has been hijacked 라는 에러가 난다.

### 9월 16일
웹소켓 클라이언트-서버 연결이 성공적으로 됐다. 
하지만 연결하면 http: connection has been hijacked 라는 에러가 난다.
132번째 줄 수정해야함 - 접속했습니다 메세지는 testsendtoClient말고 메서드를 만들어야함 (진우 : 님이 입장하셨습니다. 이런식으로 되서)

### 9월 17일 
연결하면 http: connection has been hijacked 라는 에러(경고문을) 없애야함.