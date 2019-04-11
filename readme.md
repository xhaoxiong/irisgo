#### 安装使用(仅支持go1.11以上版本)
##### 安装
```
go get -u -v github.com/xhaoxiong/irisgo

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
│
└ ─ ─web
    └ controllers
    │  Common.go
    │  TestController.go
    └ middleware
       jwt.go

```