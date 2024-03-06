N_THREADS=4

_common_folders:
	mkdir -p configs/graphite
	mkdir -p configs/grafana_config
.PHONY: _common_folders

build_image:
	docker rmi aes_go -f || true
	docker build -t aes_go -f ./docker/Dockerfile .

setup: _common_folders build_image

run_image:
	docker run -it --rm --name aes_go aes_go

deploy:
	N_THREADS=$(N_THREADS) docker stack deploy -c docker/docker-compose.yml aes_go

remove:
	docker stack rm aes_go

logs:
	docker service logs aes_go_app


# local_run:
# go mod download
# go get aes_go
#	go run src/main.go
# diff input.txt deciphered.txt