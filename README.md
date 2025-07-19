升级 go.mod 中的 Kitex 版本到最新（推荐）
```shell
go get github.com/cloudwego/kitex@latest
go mod tidy
```
把 Kitex 库升级到和命令行工具兼容的版本。

```shell
kitex -module hanyuMQ -service client_operations ./operations.thrift
kitex -module hanyuMQ -service raftoperations ./raftoperations.thrift
```