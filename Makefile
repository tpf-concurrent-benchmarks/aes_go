
build:
	cd src && go build -o ../bin/main.exe
	docker rmi aes_go -f || true
	docker build -t aes_go -f ./docker/Dockerfile .

run_local:
	go run src/main.go