# gfcommon

# 简介

基于goframe，对业务开发通用的一些规范做了封装。

## 鉴权

- [context.go](auth%2Fcontext.go) 鉴权上下文
- [middleware.go](auth%2Fmiddleware.go) 鉴权中间件

## 缓存

- [cache.go](cache%2Fcache.go) 缓存
- [dbCache.go](cache%2FdbCache.go) 数据库缓存工具库

## 命令

- [cmd.go](cmd%2Fcmd.go) 服务命令
- [middleware.go](cmd%2Fmiddleware.go) 服务命令全局中间件
- [http.go](cmd%2Fhttp.go) http服务
- [grpc.go](cmd%2Fgrpc.go) grpc服务
- [consume.go](cmd%2Fconsume.go) 消费者服务
- [xxljob.go](cmd%2Fxxljob.go) xxljob服务

## 定时任务

## 日志

- [formatJson.go](logging%2FformatJson.go) 日志json化处理
- [httpMiddleware.go](logging%2FhttpMiddleware.go) http日志中间件
- [grpcUnary.go](logging%2FgrpcUnary.go) grpc日志中间件
- [httpAccess.go](logging%2FhttpAccess.go) http请求日志
- [grpcAccess.go](logging%2FgrpcAccess.go) grpc请求日志
- [error.go](logging%2Ferror.go) 错误日志

## 业务指标

- 业务指标日志

## 消息队列

## 响应

- [http.go](response%2Fhttp.go) http响应格式

