如何生成 protobuf go 文件以及生成gorm注解

一、生成protobuf go文件

1.1 安装protoc和protobuf-go

下载系统对应的可执行文件并放入 $GOPATH/bin/下。

https://github.com/protocolbuffers/protobuf/releases
https://github.com/protocolbuffers/protobuf-go/releases

1.2 在apis文件夹下执行命令。
protoc --go_out=. --plugin= types.proto
就会看到一个pb.go文件。

二、生成gorm注解
2.1 安装protoc-go-inject-tag
go install github.com/favadi/protoc-go-inject-tag@latest

go install github.com/favadi/protoc-go-inject-tag@latest

2.2 在apis 文件夹下执行命令
protoc-go-inject-tag -input="*.pb.go" // 不要改动
就会在pb.go文件中看到有gorm的后缀。
（注意：前提是，必须在 .proto 文件中加入// @gotags: 的注释，
然后运行上面的命令，才会在 .pb.go文件中出现gorm的注释）
//该命令是读取tag注释，根据tag注释生成注解


// package,option go_package都是什么意思？命令里是什么意思？参考下面的文档
package apis;

option go_package = "../apis";
https://www.cnblogs.com/zhangcaiwang/p/15755264.html
protocol buffer中的一些点（pacakge、go_package、proto依赖等）
protoc --proto_path=$GOPATH/src --proto_path=. --go_out=. ./*.proto
a. 上面的句编译语句中，--proto_path用于表示要编译的proto文件所依赖的其他proto文件的查找位置，可以使用-I来替代。如果没有指定则从当前目录中查找。
b. --go_out有两层含义，一层是输出的是go语言对应的文件；一层是指定生成的go文件的存放位置。
c. --go_out=plugins=grpc:helloworld，这里使用了grpc插件。如果proto文件想在rpc中使用，可以在proto中定义接口如下：
service SearchService {
rpc Search(SearchRequest) returns (SearchResponse);
}
helloworld表示生成的文件存放地址。
protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative ./update.proto
a. --go_opt表示生成go文件时候的目录选项，如上面写时表示生成的文件与proto在同一目录。
import、go_package、package
a. package主要是用于避免命名冲突的，不同的项目（project）需要指定不同的package。
b. import，如果proto文件需要使用在其他proto文件中已经定义的结构，可以使用import引入。
c. option go_package = "github.com/protocolbuffers/protobuf/examples/go/tutorialpb"; go_packge有两层意思，一层是表明如果要引用这个proto生成的文件的时候import后面的路径；一层是如果不指定--go_opt（默认值），生成的go文件存放的路径。
d. 需要注意的是package和go_package的含义。在官方给的文档中，package和go_package的最后一个单词不一样：

他们的含义分别是：package用于防止不同project之间定义了同名message结构的冲突，因为package名的一个作用是用于init方法中的注册：

而当go_package存在时，其最后一个单词是生成的go文件的package名字：

而当go_package不存在时，go文件件的package名字就变成了proto中package指定的名字了。



一个简单的 protobuf 文件定义如下:

response.proto

syntax = "proto3";

option go_package = "github.com/TripleCGame/apis/api;api";
import "google/protobuf/struct.proto";

message Response {
int32 code = 1;
google.protobuf.Struct data = 2;
string msg = 3;
}
syntax = “proto3”;—指定使用 proto3 语法

option go_package = "github.com/TripleCGame/apis/api;api";—前一个参数用于指定生成文件的位置，后一个参数指定生成的 .go 文件的 package 。具体语法如下：

option go_package = "{out_path};out_go_package";
注意：这里指定的 out_path 并不是绝对路径，只是相对路径或者说只是路径的一部分，和 protoc 的 --go_out 拼接后才是完整的路径。