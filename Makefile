generate:
	rm -rf ./gen
	cd ${GOPATH} && find ./github.com/Ultimate-Super-WebDev-Corp/gateway/services -type f -iname "*.proto" -exec \
	protoc --proto_path=${GOPATH}/src --proto_path=./ --go_out=plugins=grpc:./ --govalidators_out=. {} \;

test:
	go test -cover ./...

up:
	docker-compose up -d --build

down:
	docker-compose down

create_pg_migrate:
	migrate create -ext sql -dir migrations/pg -seq $(migrate_name)

up_pg_migrate:
	migrate -database postgresql://gateway:gateway@localhost:5432/gateway?sslmode=disable -path ./migrations/pg up

down_pg_migrate:
	migrate -database postgresql://gateway:gateway@localhost:5432/gateway?sslmode=disable -path ./migrations/pg down

build:
	docker build -t ultimatesuperwebdevcorp/gateway:$(tag) .

push:
	docker push ultimatesuperwebdevcorp/gateway:$(tag)
