migration-up:
	go run migrations/main.go up

migration-down:
	go run migrations/main.go down

migration-drop:
	go run migrations/main.go drop

migration-create:
	migrate create -ext sql -dir ./migrations/data -seq $(name)

devup:
	docker-compose build && docker-compose up

build:
	docker build -t bearners-backend .

deploy:
	heroku container:push bearners-backend -a amazingtalker && heroku container:release bearners-backend -a amazingtalker

install:
	brew install golang-migrate &
	go get -u github.com/onsi/ginkgo/ginkgo &
	go get -u github.com/onsi/gomega/... &
	go get -u github.com/golang/mock/mockgen
	go get -u github.com/spf13/cobra/cobra

deploy-to-ecr:
	docker build -t bearners-backend . && 
	docker tag bearners-backend\:latest 502266988632.dkr.ecr.us-east-1.amazonaws.com/bearners-backend\:latest && 
	docker push 502266988632.dkr.ecr.us-east-1.amazonaws.com/bearners-backend\:latest

test:
	go test -cover ./...

vet:
	go vet ./

start-prod:
	wget https://zi-wei-dou-shu-config.s3-ap-southeast-1.amazonaws.com/env.yaml --output-document=$WORKDIR/configs/yaml/release-config.yml
	zi-wei-dou-shu-gin