# 基本思想解析
## 1、核心的分层概念
我们使用这个框架，首要的就是要了解service、endpoint、transport这三个概念并编写对应代码，然后可以拼成一个服务供访问。

## 2、我们先来聊聊transport层怎么来的？
我们要搭建的微服务体系，不确定用户会用使用什么协议，比如http、grpc、thrift等等，这些得交给一个层去实现，这个就是transport层。

它可以对多种协议进行处理的实现，但是他与endpoint层之间就用规范的接口标准去定义，这样层与层之间解耦合并有清晰的接口关系。

transport底层负责把协议中数据提取出来，交给endpoint层处理，然后拿到endpoint层返回的数据，再通过协议封装返回。 
总结下来，transport层就是要确认用什么协议，然后对每个路由应用一个endpoint层的接口实现。

不同协议返回的数据不一样，http的和rpc的数据有自己的格式，不过给到endpoint和从endpoint拿的数据都是一样的，
这样transport层还得实现对每个路由有关的一对DecodeRequest和EncodeResponse函数。

总结下来transport层就是两件事，一是对接具体协议，二是编解码数据对接enpoint层。

## 3、再来聊聊endpoint层怎么来的？
上面说到接口，这个接口如果是我们怎么设计呢，我们需要处理过来的路由，一个路由一个处理器。

处理器就是一个函数，参数是从协议里解析的，不确定格式我们就用万能的interface{}，返回值也不好说，那就也用万能的interface{}，这样一来，
我们可以定义出这样一个接口就叫Point吧，Point接口就一个方法如上所说的，就是func(context.Context, interface{}) (interface{}, error)。
transport层处理完协议把不同路由的请求就调用不同的Point接口实现即可。

考虑去实现Point接口，对一个路由我们创立一个对应结构体实现Point接口，里面的方法就是具体业务实现。这样代码写的很繁琐，既然接口就一个方法，
那不如直接定义一个函数类型得了，这样就出来了EndPoint函数类型，对一个路由我们实现一个函数类型对象即可。如下就是Endpoint函数类型：
```go
type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)
```

## 4、再来聊聊service层怎么来的？
其实上面说到的两层即endpoint和transport就已经可以把代码规划层次然后各司其职了， 把所有业务逻辑都写到一个个enpoint这个函数类型的实现中。

但在实际开发中，我们得用很多中间件，得用redis、用mysql、用其他依赖服务等等，对应的项目中就是一个个struct并构建之间复杂的关系。

类似的我们可能会构建出一套MVC分层的代码来实现业务。对一个具体的请求，我们使用Controller层与之对接。基于这个背景，go-kit让我们实现一个业务层，
这里就叫service层。在service层实现业务处理的接口，endpoint也不关心你的具体实现，就是定义好接口方法，它处理的时候按照标准和service层交互就好了。

## 5、最后看看怎么把这三层整合起来呢？
首先endpoint层依赖service层，我们定义一个接口比如就叫IService，我们实例化endpoint的时候就提供一个具体的IService接口实现。

上面我们说到Endpoint实现不是一个标准struct，如果是我们可以给他加上IService依赖，实例化的时候提供实现即可。现在我们的每个EndPoint已经
是一个函数类型实现，想要往里填入对IService接口的依赖，可以考虑使用闭包。闭包可以让函数类型的实现依赖外部的变量，而且在使用时在加载。
因此，我们本来要实现一个个Endpoint函数类型的实现，现在要改为实现一个个闭包，它让一个函数类型对象和外部变量绑定了。
```go
func MakeEndPointArticleAdd(svc my_service.IService) endpoint.Endpoint {
   return func(ctx context.Context, request interface{}) (response interface{}, err error) {
      req := request.(my_service.ArticleAddRequest)
      res := svc.ArticleAdd(ctx, req)
      return res, nil
   }
}
```

在说下transport层和endpoint层，endpoint本身就是一种函数类型了，我们实现一个transport层的处理器时，就可以提供一个Endpoint函数类型实现。
比如对于http协议的处理，transport要提供一个Server结构体，它包含了Endpoint、DecodeRequest、EncodeResponse，
我们实例化它时提供好我们的endpoint层的具体实现接口。
```go
handler := httpTransport.NewServer(
   endpointArticleAdd,
   my_tansport.DecodeArticleAddRequest,
   my_tansport.EncodeArticleAddResponse,
)
```