### 前置工作
- 安装 protobuf
- 安装PHP相关扩展（`grpc`、`protobuf`）
- 生成grpc_php_plugin

    - [官方地址](https://github.com/grpc/grpc/tree/master/src/php)
    - 实际操作
  ```shell
  $ git clone -b RELEASE_TAG_HERE https://github.com/grpc/grpc
  $ cd grpc
  $ git submodule update --init
  $ mkdir -p cmake/build
  $ cd cmake/build
  $ cmake ../..
  $ make protoc grpc_php_plugin
  $ mv ./grpc_php_plugin /opt/homebrew/bin/
  ```
- 生成相应语言的文件命令
  ##### GO
  ```shell
   protoc -I. --php_out=plugins=grpc:./php pb/user.proto
  ```
  ##### PHP
  ```shell
  cd php-server
  
  protoc --proto_path=../pb --php_out=. --grpc_out=. --plugin=protoc-gen-grpc=/opt/homebrew/bin/grpc_php_plugin ../pb/user.proto
  ```

### 启动GRPC服务
```shell
➜  grpc-demo go run server.go 
2022/01/24 09:55:49 receive user index request: page 1 page_size 12
2022/01/24 09:55:49 receive user view request: uid 1
2022/01/24 09:55:49 receive user post request: name 12312 password 324523 age 12
```
### 启动GRPC客户端

##### GO GRPC客户端
  ```shell
  ➜  grpc-demo go run client/client.go 
  2022/01/24 09:56:46 user index success: success
  u1 11
  u2 12
  2022/01/24 09:56:46 user view success: success
  james 28
  2022/01/24 09:56:46 user post success: success
  2022/01/24 09:56:46 user delete success: success
  ```
##### PHP GRPC客户端
  ```shell
  ➜  php-server php  client.php
  success
  success
  success
  ```
