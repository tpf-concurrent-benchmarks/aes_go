
_common_folders:
	mkdir -p configs/graphite
	mkdir -p configs/grafana_config
.PHONY: _common_folders

build_image:
	docker rmi aes_go -f || true
	docker build -t aes_go -f ./docker/Dockerfile .

dummy_file:
	mkdir -p data
	echo "Hello World!" > data/input.txt

setup: _common_folders build_image

run_image:
	docker run -it --rm --name aes_go aes_go

deploy:
	docker stack deploy -c docker/docker-compose.yml aes_go

remove:
	docker stack rm aes_go

logs:
	docker service logs aes_go_app -f

# cd src
# go mod download
#	go get aes_go
#	go run main.go
# diff input.txt deciphered.txt