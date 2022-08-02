# LOCAL

install:
	go mod download

run:
	go run ./cmd/dota2/main.go

dev:
	gin -p 3001 -a 3006 -i -d cmd/dota2/ run main.go

####################################

# DOCKER

docker-build:
	docker build -f docker/prd/Dockerfile -t renzotelenta.com/dota2:latest .

docker-run:
	docker run --rm -it \
	-p 3006:3006 \
	-v "${PWD}/cmd:/app/cmd" \
	-v "${PWD}/domain:/app/domain" \
	-v "${PWD}/drivers:/app/drivers" \
	--env-file ./.env \
	--name dota2 \
	renzotelenta.com/dota2:latest