tag = dev

generate:
	protoc --proto_path=${GOPATH}/src --proto_path=./ --go_out=./ --go-grpc_out=./ --govalidators_out=. ./services/*/*.proto

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
	grpc-client-cli \
	--proto ./services/*/*.proto \
	--protoimports ${GOPATH}/src 0.0.0.0:8080

create_migrate:
	migrate create -ext sql -dir migrations/$(service) -seq $(migrate_name)

up_migrate_local: #todo move to docker-compose
	migrate -database postgresql://customer:customer@localhost:5432/customer?sslmode=disable -path ./migrations/customer up
