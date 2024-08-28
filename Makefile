run:
	modd -f ./.modd/modd.conf

tidy:
	go mod tidy

migrate-up:
	go run main.go migrate


migrate-down:
	go run main.go migrate -d down

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. \
--go-grpc_opt=paths=source_relative pb/auth/*.proto