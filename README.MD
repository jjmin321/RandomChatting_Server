# 대소고인을 위한 1대1 랜덤채팅 서비스

### 9월 15일
웹 소켓 서버를 여는 것을 성공적으로 띄웠다.
하지만 연결하면 http: connection has been hijacked 라는 에러가 난다.

### 9월 16일
웹소켓 클라이언트-서버 연결이 성공적으로 됐다. 
하지만 연결하면 http: connection has been hijacked 라는 에러가 난다.
132번째 줄 수정해야함 - 접속했습니다 메세지를 보내기 위해서 testsendtoClient이외의 메서드를 만들어야함 (현재 정민 : 님이 입장하였습니다 로 됨.)

### 9월 17일 
echo: http: response.Write on hijacked connection from github.com/labstack/echo.(*Response).Write (response.go:75) 에러를 성공적으로 해결했다.
126번째 줄에 testsendToJoinMsgToClient 메서드를 호출하는 코드를 생성해야 함.
testsendToJoinMsgToClient 메서드를 만들어야함. 

### 9월 18일 
162번째 줄 수정