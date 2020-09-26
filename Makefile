generate:
	rm -rf ./gen
	find ./services -type f -iname "*.proto" -exec \
	protoc --proto_path=${GOPATH}/src --proto_path=./ --go_out=plugins=grpc:./ --govalidators_out=. {} \;

test:
	go test -cover ./...

up:
	docker-compose up -d --build

down:
	docker-compose down

get_project_dependencies:
	#todo add installation of protoc if it doesn't exist
	GO111MODULE=on go get github.com/mwitkow/go-proto-validators/protoc-gen-govalidators
	GO111MODULE=on go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate

curl:
	$(eval proto_cmd :=$(shell find ./services -type f -iname "*.proto" -exec printf " --proto {}" \;))
	grpc-client-cli $(proto_cmd) \
	--protoimports ${GOPATH}/src 0.0.0.0:8080

create_pg_migrate:
	migrate create -ext sql -dir migrations/pg -seq $(migrate_name)

up_pg_migrate:
	migrate -database postgresql://gateway:gateway@localhost:5432/gateway?sslmode=disable -path ./migrations/pg up

down_pg_migrate:
	migrate -database postgresql://gateway:gateway@localhost:5432/gateway?sslmode=disable -path ./migrations/pg down

build:
	docker build -t ultimatesuperwebdevcorp/customer:$(tag) .

push:
	docker push ultimatesuperwebdevcorp/customer:$(tag)