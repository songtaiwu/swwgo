# 一、回顾grpc

我们先简单看下如何构建一个grpc服务的，梳理出里面的主要关键点。

## 1、准备工作

需要使用grpc，我们得先安装把proto文件能转为代码的工具，即protoc，去官网下载zip包解压得到protoc.exe。

因为要使用的go语言，需要安装protoc执行时的插件，支持把proto转为go语言代码。为了使用grpc，还要安装protoc工具关于grpc的插件。两个插件都可以通过go get的方式安装。

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

最后我们就可以通过如下两个命令，把一个proto文件转为两个go代码文件。

```go
protoc --go_out=. *.proto
protoc --go-grpc_out=. *.proto
```

得到的文件名格式

- xxxx.pb.go
- xxxx_grpc.pb.go



## 2、如何实现server

以我们定义的article.proto为例，执行完protoc命令后，得到如下文件

- article.pb.go
- article_grpc.pb.go

在article_grpc.pb.go中实现了接口ArticleServer

```go
// ArticleServer is the server API for Article service.
// All implementations must embed UnimplementedArticleServer
// for forward compatibility
type ArticleServer interface {
	ArticleAdd(context.Context, *AddRequest) (*AddResponse, error)
	mustEmbedUnimplementedArticleServer()
}
```

在article_grpc.pb.go中还提供好了把ArticleServer接口实现注册到grpc服务注册器中的方法

```go
func RegisterArticleServer(s grpc.ServiceRegistrar, srv ArticleServer) {
	s.RegisterService(&Article_ServiceDesc, srv)
}
```

对于业务开发人员，我们要做的工作就是如下几步：

1. 创建一个结构体并实现proto中的ArticleServer接口，如上，只要实现ArticleAdd这个方法即可。方法中的代码逻辑就是我们具体的业务处理。
2. 创建tcp监听，通过google.golang.org/grpc包提供的 grpc.NewServer()方法创建出服务（含注册功能），即提供了RisterService（）来注册我们实现的服务。
3. 调用proto中提供好的方法，把我们的ArticleServer接口实现注册到grpc服务中。
4. 调用grpc服务的Server()启动监听。

```golang
func main() {
    // tpc监听
	lis, err := net.Listen("tcp", "0.0.0.0:1001")
	if err != nil {
		log.Fatalln("failed to listen")
	}
	// 创建grpc服务
	grpcServer := grpc.NewServer()
    // 注册我们实现的服务
	proto.RegisterArticleServer(grpcServer, new(ArticleServer))
    // 启动服务
	grpcServer.Serve(lis)
}
```



# 二、如何集成到go-kit

## 1、从基础概念推导

### 1.1 必须实现proto中的Server接口

要提供grpc服务，我们离不开google.golang.org/grpc包提供的基础能力。要创建grpc服务`grpcServer := grpc.NewServer()`, grpcServer服务提供了方法`func (s *Server) RegisterService(sd *ServiceDesc, ss interface{})`来注入具体的service，这个service都是业务开发人员按照服务接口定义实现的具体struct。从这个角度看，这些都是最底层的方式，无法避开的。

比如我们的article.proto构建出article_grpc.pb.go之后，里面有个接口如下，我们必须实现它，它既是我们业务的处理逻辑。

```go
type ArticleServer interface {
	ArticleAdd(context.Context, *AddRequest) (*AddResponse, error)
	mustEmbedUnimplementedArticleServer()
}
```

### 1.2 必须有go-kit的service

go-kit的思想是想在transport层处理各类协议并把负责数据的encode和decode即代码中定义的业务数据与协议提供的数据转换。那么对于grpc协议，go-kit需要提供能从grpc协议中拿到数据并把业务数据返回grpc通信的能力。transport会把真正的处理交给endpoint接口去处理，只要提供给它`context.Context`和数据 `interface{}`即可。endpoint会把处理能力最终交给go-kit的service层，go-kit的service层不在乎协议，它就是最朴实的业务处理，有自己的请求响应，有自己的业务逻辑代码，如果一个服务既提供http也提供grpc，对于service层的方法都是一个。

基于上面go-kit思路，很显然，我们先得实现go-kit所定义的service层的业务代码，这个可不是xxxx_grpc.pb.go中定义的接口实现，而是纯粹的业务处理，不关乎协议的。等实现好service，就得考虑怎么提供给endpoint层。  方式是一样的，就是endpoint层都是函数类型的实现，通过闭包方式把service传进去。要把service层业务代码提供给endpoint层，之前做http的服务时，是对于一个service的方法（一个业务处理）转为一个endpoint，是利用闭包把service注入到其中，让调用这个endpoint时，实际还得调用里面service的具体业务处理。这里我们仍然是要这么实现，这些endpoint闭包也是不关心协议的，他们就是如下的函数类型：

```go
func(ctx context.Context, request interface{}) (response interface{}, err error)
```

### 1.3 endpoint层提供给transport层

transport要把协议中要处理的方法，用go-kit提供的方式处理，而且处理方法里面套上三个核心元素，即endpoint、decode、encode。

比如go-kit集成http，net/http提供了一个路由和一个handler的映射，handler接口定义如下：

```go
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

go-kit就提供了自己的handler，它会包装把三个元素包装成一个handler然后处理一个路由：

```go
// 使用kit创建handler
	handler := httpTransport.NewServer(
		endpointArticleAdd,
		my_tansport.DecodeArticleAddRequest,
		my_tansport.EncodeArticleAddResponse,
	)
```

参照这个思路，grpc也是一样，grpc用的handler是什么，go-kit也需要提供能把三个元素包装成这个handler的方法。

但是对于grpc，我们上面说明，得有个实现grpc Server的接口实现，比如`ArticleServer`接口实现，实现里面的方法`ArticleAdd(context.Context, *AddRequest) (*AddResponse, error)`。grpc协议过来，也是先要交给`ArticleServer`接口实现来处理的，它好比是http路由处理必须提供的handler，这里是grpc，必须提供grpc Server接口实现。

那么go-kit就要提供方式，来重写这个handler，即重写我们实现的grpc Server接口的结构体中的每个方法。

我们需要三个元素：

- endpoint，它跟其他协议比如http的endpoint都是一样的，都是`func(ctx context.Context, request interface{}) (response interface{}, err error)`
- decode, 它是从配合ArticleServer这种grpc Server接口实现的，它拿到数据是proto中定义的比如*AddRequest，然解析为go-kit service层需要的数据。
- encode，它也是配合ArticleServer这种grpc Server接口实现的，它从endpoint拿到数据，转为proto中要响应的数据格式，比如*AddResponse。

这里就发现，如何把endpoint转为 Server实现才是**核心问题**。

Server接口实现,, 我们比如叫  `articleServer` 里面就是proto中`ArticleServer`接口定义的，需要实现里面的一个个方法。这些方法实现不是自己直接写业务逻辑了，而要要交给go-kit的transport层定义的handler，这样它才能继续交给endpoint，这样就协议解耦了。

那么`articleServer`中的方法怎么由transport层接管呢？go-kit提供了包`github.com/go-kit/kit/transport/grpc`提供了这个能力。grpc包提供的如下方法

```go
grpc.NewServer(
			endpoint,
			DecodeArticleAddRequest,
			EncodeArticleAddResponse,
			)
```

它可以把endpoint闭包和encode、decode方法转为一个Handler接口实现。我们看下`github.com/go-kit/kit/transport/grpc`包Handler接口定义，它用于处理grpc协议的通用方式。

```go
// Handler which should be called from the gRPC binding of the service
// implementation. The incoming request parameter, and returned response
// parameter, are both gRPC types, not user-domain.
type Handler interface {
	ServeGRPC(ctx context.Context, request interface{}) (context.Context, interface{}, error)
}
```

比如我们一个微服务中定义了多个proto，我们得到多个xxxx_grpc.pb.go文件，对应的就有多个XXXXServer接口。我们要实现这些XXXXServer接口，就会有多个xxxxServer实现。这些实现要注入到grpc服务中，最终就能基于tcp连接提供服务。

每一个xxxxServer中每个业务处理方法，都不再是自己填入业务处理，而是要使用一个Handler接口实现，Handler实现服务从协议中encode、decode数据然后交给endpoint处理。一个xxxxServer比如有十几个业务方法，对应的就得有十几个Handler实现，每个Handler实现绑定一个endpoint、一个decode、一个encode。









