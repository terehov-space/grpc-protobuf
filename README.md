Before run and testing:

``cd auth-service && go mod tidy``

then

``cd user-service service && go mod tidy``

To connect to local database change file in user-service:

``user-service/database/connection.go``

change to your own data:
``dsn := "postgresql://ovr_user:ovr_pass@localhost:5429/performance"``

---

After changing proto/user-service.proto you need to regenerate grpc autogen files:

``cd user-service service && protoc --proto_path=../proto/ --go_out=. --go-grpc_out=. ../proto/user-service.proto``

``cd auth-service && protoc --proto_path=../proto/ --go_out=. --go-grpc_out=. ../proto/user-service.proto``