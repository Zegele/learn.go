如何生成 protobuf go 文件以及生成gorm注解

一、生成protobuf go文件

1.1 安装protoc和protobuf-go

下注系统对应的可执行文件并放入 $GOPATH/bin/下。

https://github.com/protocolbuffers/protobuf/releases
https://github.com/protocolbuffers/protobuf-go/releases

1.2 在apis文件夹下执行命令。
proto --go_out=. --plugin= types.proto
就会看到一个pb.go文件。

二、生成gorm注解
2.1 安装proto-go-inject-tag
go install github.com/favadi/protoc-go-inject-tag@latest

go install github.com/favadi/protoc-go-inject-tag@latest

2.2 在apis 文件夹下执行命令
protoc-go-inject-tag -input="*.pb.go"
就会在pb.go文件中看到有gorm的后缀。
（注意：前提是，必须在 .proto 文件中加入// @gotags: 的注释，
然后运行上面的命令，才会在 .pb.go文件中出现gorm的注释）