##### 更新日志
```
v1.0.0 默认初始化基础的api模板 gorm+iris restful接口
v1.1.1-v1.1.5 修改iris版本更新代码，添加日志组件
v1.1.1-v1.1.6 将jinzhu/gorm修改为gorm.io/gorm，并且添加casbin权限初始化
```

##### <a href="#">DEMO后台管理(iris-react-admin)正在开发中预计12月份第一个基础版本</a> 

## V1.0.0
#### 安装使用(仅支持go1.11以上版本推荐使用go1.13版本，irisv12要求go1.13)
##### 安装
```
go get -u -v github.com/xhaoxiong/irisgo

也可以直接clone编译，最终目的是编译好直接跑就好了。
```
##### 使用
```
>~ irisgo new [项目名称(默认为irisApp)] 
>~/irisApp go mod init [项目名]
>~/irisApp go build -v
windows下
>~/irisApp irisApp.exe 运行
linux下
>~/irisApp ./irisApp 
此项目生成对应目录，运行后跑在127.0.0.1:8080,请求访问将返回
{
  "code": 10000,
  "message": "success"
}

只有将config.yaml中数据库配置写好才能操作models.GetDB()操作gorm 否则会报错

设置代理:
set GOPROXY="https://goproxy.cn"  export GOPROXY="https://goproxy.cn"
set GOPROXY="https://goproxy.io"  export GOPROXY="https://goproxy.io"
set GOPROXY="http://mirrors.aliyun.com/goproxy/"  export GOPROXY="http://mirrors.aliyun.com/goproxy/"
```
##### <a href="https://gorm.io/zh_CN/docs/">gorm中文官网</a>
##### <a href="https://learnku.com/docs/iris-go/10">iris中文文档翻译1</a>
##### <a href="https://www.studyiris.com/doc/index.html">iris中文文档翻译2</a>
#####  <a href="https://github.com/iris-contrib/swagger">引入swagger点击此处(需采用v12版本)</a>

##### V1.1.0-v1.16
#### 安装使用(仅支持go1.11以上版本推荐使用go1.13版本，irisv12要求go1.13)
##### 安装(同v1.0.0)
##### 使用(修改了命令行参数，及其对应的代码结构，具体操作如下)
```
Mac: 
命令使用如下:
不带-name参数将默认生成irisApp
>~ ./irisgo -n test
默认生成api模板脚手架

Windows:
>~ irisgo.exe -n test 

```


##### 生成的目录结构如下

```
App
│ main.go
└ ─ ─conf[配置信息]
│      config.yaml 
└ ─ ─conifg[热加载与viper的启用]
│      config.go
└ ─ ─models
│      init.go(数据库初始化)
└ ─ ─repositories[持久层]
│      TestRepo.go
└ ─ ─service
└ ─ -route
│    	route.go
└ ─ ─web
    └ controllers
    │  Common.go
    │  TestController.go
    └ middleware
       jwt.go
       logrus.go

```

