# grpc-tcp-gateway
grpc tpc/ws（tls/wss）协议网关

## 相关仓库

```
https://github.com/generalzgd/grpc-tcp-gateway
https://github.com/generalzgd/protoc-gen-grpc-tcpgw
https://github.com/generalzgd/grpc-tcp-gateway-proto
```

### Proto文件生成
```shell script
protoc -Iiproto --go_out=plugins=grpc:./iproto ./iproto/backendsvr.proto
protoc -Iiproto --go_out=plugins=grpc:./iproto ./iproto/tcpgate.proto
protoc -Iiproto --grpc-tcpgw_out=logtostderr=true:./iproto ./iproto/tcpgate.proto
```

## Schema

```
请参考 https://github.com/generalzgd/protoc-gen-grpc-tcpgw
```

## 特点

```
1. 客户端不用关心后端服务有哪些，只需知道网关地址。由网关根据包头信息自动路由到后端服务并返回对应数据。
2. 支持双向数据发送. 
3. 支持服务端主动下发消息给客户端
4. 同时支持protobuf和json两种协议格式
5. 对比grpc-ecosystem/grpc-gateway
5.1 ecosystem需要为每个后端服务都注册一个网关地址和端口，客户端需要关心对应服务的网关和端口。
5.2 ecosystem只支持http的短连接访问，不支持双向数据发送。
6. 支持路由转发给不同的后端服务
7. grpc转发支持后端服务发现和均衡负载
```

## 性能点

```
稍后补上
```

## PS

```
目前该项目处于试运行阶段，尚有不足之处。恳请广大网友提点迷津。
```

![usecase](./uml/usecase.png)