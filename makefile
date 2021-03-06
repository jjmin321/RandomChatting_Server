# dockerfile 빌드
.PHONY: build 
build:
	@docker build --tag randomchatting_server .

# 서버 컨테이너 실행, 접속 
.PHONY: run
run:
	@docker run -i -t -p 8080:8080/tcp --name server randomchatting_server

# 실행된 dockerfile 컨테이너, 이미지 삭제
.PHONY: rm 
rm:
	@docker rm server
	@docker rmi randomchatting_server

# docker-compose.yml 서비스 시작
.PHONY: compose-up
compose-up:
	@docker-compose --env-file docker.env -f docker-compose.yml up

# docker-compose.yml 서비스 삭제
.PHONY: compose-down
compose-down:
	@docker-compose --env-file docker.env -f docker-compose.yml down

# 데이터베이스 컨테이너 접속
.PHONY: postgresql
pg:
	@docker exec -it postgres psql -Ujejeongmin
