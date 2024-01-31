# github.com/liyonge-cm/go-api-cli

Author:liyonge(aiee)

根据表结构直接自动化生成各个表的CRUD，增加(Create)、读取(Read)、更新(Update)和删除(Delete)。

## API规则
首先建立好规则，方便生成逻辑统一的CRUD：

### 1.数据库表结构规则
1. id为int类型自增数值
2. 都有创建、更新时间，字段名固定为created_at,updated_at，且为int64类型，存值为时间戳
3. 都有状态字段，字段名固定为status，int类型，初始值为0，删除为-1

### 2.CRUD规则
1. Create时，除了id、status,created_at,updated_at，其他字段都作为入参，入库时初始status为0，创建、更新时间为当前时间戳
2. Read时，带分页功能，分页参数统一为：limit（每页条数）、offset（第几页，从1开始），响应所有表字段原始数据
3. Update时，必传id参数，除status,created_at,updated_at外的其他字段都可更改，更新时间自动设置为当前时间戳
4. Delete时，必传id参数，status设置为-1，更新时间自动设置为当前时间戳

### 3. api 风格
1. API method统一为post，body以json入参。如
```shell
curl -X POST "http://localhost:8080/user/getList" -H  "accept: application/json" -d "{\"limit\": 10,\"offset\": 1}" 
```
2. 以表名作为group，crud分别为create,get,getList,update,delete。
如有一个表名为user，生成API请求地址分别为：
- http://localhost:8080/user/create 
- http://localhost:8080/user/get 
- http://localhost:8080/user/getList 
- http://localhost:8080/user/update 
- http://localhost:8080/user/delete 


## cli使用

首先下载源码，编译可执行文件
```shell
go build
```

1. 按规则创建库表

2. 修改配置文件
```yml
mysql: mysql连接配置

frame:
  out_path: 项目文件所在位置
  prj_name: 项目名称
  json_case: api入参处参的json格式：camel-驼峰，默认下划线

api:
  tables: 
    - user 指定要生成API的表名
```

3. 创建项目

执行编辑文件创建项目
```shell
go-api-cli -g frame
```
或直接运行main文件
```shell
go run main.go -g frame
```

4. 生成API

执行编辑文件生成API
```shell
go-api-cli -g api
```
或直接运行main文件
```shell
go run main.go -g api
```

## 启动生成的新项目
```shell
# 1. 进入项目文件: 
cd xxx/prj-aiee-api

# 2. 下载依赖包: 
go mod tidy

# 3. 修改配置config，数据库连接地址等

# 4. 启动: 
go run main.go

# 5. API调用，user为表名
curl -X POST "http://localhost:8080/user/getList" -H  "accept: application/json" -d "{}" 

```

生成的项目代码保持简洁易懂，方便根据项目实际需求二次开发。