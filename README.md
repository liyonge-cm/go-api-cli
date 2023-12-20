# go-cli
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

## cli使用
1. 按规则创建库表

2. 修改配置文件
mysql: mysql连接配置

frame:
  out_path: 项目文件所在位置
  prj_name: 项目名称
  json_case: api入参处参的json格式：camel-驼峰，默认下划线

api:
  tables: 
    - user 指定生成API的表

3. 创建项目
go run main.go -g frame

4. 生成API
go run main.go -g api

## 启动生成的新项目
1. 进入项目文件: cd xxx/prj-aiee
2. 下载依赖包: go mod tidy
3. 修改配置config，数据库连接地址等
4. 启动: go run main.go
