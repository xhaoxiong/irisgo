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

方法一:
设置代理:https://goproxy.io/ 


方法二:
如遇到翻墙不能get的包请在.mod文件中加入以下


replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.34.0
	go.opencensus.io => github.com/census-instrumentation/opencensus-go v0.19.0
	go.uber.org/atomic => github.com/uber-go/atomic v1.3.2
	go.uber.org/multierr => github.com/uber-go/multierr v1.1.0
	go.uber.org/zap => github.com/uber-go/zap v1.9.1
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20181001203147-e3636079e1a4
	golang.org/x/lint => github.com/golang/lint v0.0.0-20181026193005-c67002cb31c3
	golang.org/x/net => github.com/golang/net v0.0.0-20180826012351-8a410e7b638d
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20180821212333-d2e6202438be
	golang.org/x/sync => github.com/golang/sync v0.0.0-20181108010431-42b317875d0f

	golang.org/x/sys => github.com/golang/sys v0.0.0-20190312061237-fead79001313
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/time => github.com/golang/time v0.0.0-20180412165947-fbb02b2291d2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20181219222714-6e267b5cc78e
	google.golang.org/api => github.com/googleapis/google-api-go-client v0.0.0-20181220000619-583d854617af
	google.golang.org/appengine => github.com/golang/appengine v1.3.0
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20181219182458-5a97ab628bfb
	google.golang.org/grpc => github.com/grpc/grpc-go v1.17.0
	gopkg.in/alecthomas/kingpin.v2 => github.com/alecthomas/kingpin v2.2.6+incompatible
	gopkg.in/mgo.v2 => github.com/go-mgo/mgo v0.0.0-20180705113604-9856a29383ce

	gopkg.in/tomb.v1 => github.com/go-tomb/tomb v1.0.0-20141024135613-dd632973f1e7
	gopkg.in/vmihailenco/msgpack.v2 => github.com/vmihailenco/msgpack v2.9.1+incompatible
	gopkg.in/yaml.v2 => github.com/go-yaml/yaml v0.0.0-20181115110504-51d6538a90f8
	labix.org/v2/mgo => github.com/go-mgo/mgo v0.0.0-20160801194620-b6121c6199b7
	launchpad.net/gocheck => github.com/go-check/check v0.0.0-20180628173108-788fd7840127
)


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