


Утилита клиента gRPC - Evans:
https://github.com/ktr0731/evans/blob/master/README.md

команда установка:
go install github.com/ktr0731/evans@latest

вызываем Evans путем до прото-файла:
evans proto/message.proto -p 8080


Создаем протобаф:
protoc --go_out=. --go-grpc_out=. proto/message.proto