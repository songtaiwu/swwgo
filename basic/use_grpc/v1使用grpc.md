# 一、准备工具
对于go语言项目怎么用protobuf，官方有完整的资料说明，地址
https://developers.google.cn/protocol-buffers/docs/gotutorial

## 1、获取编译工具（protoc）
首先需要一个编译工具，对应的就是一个可执行文件比如protoc.exe，protoc运行是把.proto文件转换为 c++统一格式的内存数据，
然后由语言插件把内存数据转换为各自语言的代码文件，比如go语言就提供了protoc-gen-go插件。

要获取protoc，去官网下载zip包解压即可使用，windows、linux版本都有。参照protocol-buffers官网的下载说明，
要去github上release page下载需要的版本，地址看上面官网获取。

windows上，我们下载得到protoc-21.9-win64.zip，解压后把bin目录下的protoc.exe放到PATH下。我选择把可执行文件
复制到已经配置好的PATH路径下，比如C:\Program Files\Go\bin下。

## 2、获取编译工具go语言的插件
protoc-gen-go语言插件用go get就能下载，可执行文件会被放到GOPATH/bin下。
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```
有了protoc和protoc-gen-go, 使用protoc命令，就能把一个.proto文件转换为一个 .go文件。
```shell
protoc --go_out=. *.proto
```
## 3、获取编译工具的grpc插件
对于旧版本的protoc工具，只要安装完了protoc工具，就能使用protoc命令，把一个定义了grpc的.protoc文件转换为.go文件，
命令要用protoc --go_out=plugins=grpc:. hello.proto。

但是新版本google.golang.org/protobuf/ 不再支持gRPC服务定义，如果想要生成grpc代码需要使用新插件 protoc-gen-go-grpc。

安装插件，也是用go get下载，可执行文件会被放到GOPATH/bin下。
```shell
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```
使用命令
```shell
protoc --go-grpc_out=. *.proto
```

# 二、proto转pb
先按照protobuf官方使用要求，安装protoc工具，安装protoc-gen-go插件，安装protoc-gen-go-grpc插件。

1、编写proto文件， 这里的例子是article.proto， 具体内容参照我的代码。

2、然后进入项目目录 /use_grpc/v1下，即article.proto所在目录下，执行如下两个命令
```shell
protoc --go_out=. *.proto
protoc --go-grpc_out=. *.proto
```

3、会在项目目录下生成一个proto目录，里面得到两个文件
* article.pb.go
* article_grpc.pb.go

# 三、实现server
我们需要创建一个struct，它要实现article_grpc.pb.go文件中定义的接口ArticleServer，具体的方式，就是我们的业务逻辑实现。

# 四、实现client
实现client非常简单，直接使用文件中定义的NewArticleClient()就可以。