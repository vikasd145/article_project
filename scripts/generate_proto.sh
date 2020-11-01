# run this file from project root path
protoc -I=./pkg/proto --go_out=./pkg/gen ./pkg/proto/article_db.proto