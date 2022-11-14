# 本地运行jaeger
```shell
docker run -d --name jaeger \
-e COLLECTOR_ZIPKIN_HTTP_PORT=9411  \
-p 5775:5775/udp \
-p 6831:6831/udp \
-p 6832:6832/udp \
-p 5778:5778 \
-p 16686:16686 \
-p 14268:14268 \
-p 9411:9411 \
--restart=always \
jaegertracing/all-in-one:1.15
```

ui的访问方式：http://192.168.71.131:16686/search

