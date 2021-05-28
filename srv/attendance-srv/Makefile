run: 
	go run main.go

generateProto:
	protoc --micro_out=. --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/attendance/attendance.proto


createdb:
	docker exec -it database createdb --username=postgres --owner=postgres attendance-db

removedb:
	docker exec -it database dropdb --username=postgres attendance-db
