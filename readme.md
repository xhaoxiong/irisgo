#### 安装使用(仅支持go1.11以上版本)
##### 安装
```
go get github.com/xhaoxiong/irisgo

```
##### 使用
```
>~ irisgo new [项目名称(默认为irisApp)] 
>~/irisApp go mod init [项目名]
>~/irisApp go build -v

此项目仅仅是生成对应目录，仅能编译通过，只有将config.yanl中数据库配置写好才能运行
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
│
└ ─ ─service


```