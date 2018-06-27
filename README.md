# deomapi
first  go deomapi

-----------------------------------------------------------------
运行main文件

GET   localhost:8080/u        查看数据
POST  localhost:8080/user     添加数据
PUL   localhost:8080/user/id  更改数据
DELETE localhost:8080/user/id 删除数据

首先将testws项目用tomcat运行起来，访问index.html页面进行ws通信。或通过HbBuilder工具直接 ctrl+r运行项目

-----------------------------------------------------------------
curd使用postman操作：


-----------------------------------------------------------------
查询
URL localhost:8080/u 

method:GET
-----------------------------------------------------------------
创建  
URL  localhost:8080/user

method:POST
在boyd中选择x-wwww-form-urlencoded

key         value
id          1
name        tom
age         10
gender      boy
email       111@222

-----------------------------------------------------------------
更新  
URL   localhost:8080/user/1

method:PUL
在boyd中选择x-wwww-form-urlencoded

key         value
id          1
name        jerry
age         10
gender      gerl
email       111@222
-----------------------------------------------------------------
删除  
URL   localhost:8080/user/1

method:DELETE

-----------------------------------------------------------------
