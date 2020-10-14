/**
 * @Author: xiaoxiao
 * @Description:
 * @File:  jwt
 * @Version: 1.0.0
 * @Date: 2020/5/12 2:12 下午
 */
package api_template

var jwt = `
package middleware

import (
"github.com/dgrijalva/jwt-go"
jwtmiddleware "github.com/iris-contrib/middleware/jwt"
"github.com/kataras/iris/v12"
"fmt"
"time"
)

var JwtAuthMiddleware = jwtmiddleware.New(jwtmiddleware.Config{
	ValidationKeyGetter: validationKeyGetterFuc,
	SigningMethod:       jwt.SigningMethodHS256,
	Expiration:          true,
	Extractor:           extractor,
}).Serve

const jwtKey = "{{.Appname}}"

var validationKeyGetterFuc = func(token *jwt.Token) (interface{}, error) {
	return []byte(jwtKey), nil
}

var extractor =func(ctx iris.Context) (string, error) {
	authHeader := ctx.GetHeader("token")
	if authHeader == "" {
		return "", nil // No error, just no token
	}

	return authHeader, nil
}

//注册jwt中间件
func GetJWT() *jwtmiddleware.Middleware {
	jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
		//这个方法将验证jwt的token
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			//自己加密的秘钥或者说盐值
			return []byte(jwtKey), nil
		},
		//加密的方式
		SigningMethod: jwt.SigningMethodHS256,
		//验证未通过错误处理方式
		//ErrorHandler: func(context.Context, string)
		ErrorHandler: func(ctx iris.Context, e error) {
			ctx.Next()
		},
	})
	return jwtHandler
}

//生成token
func GenerateToken(msg string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"msg": msg,                                                      //openid
		"iss": "iris_{{.Appname}}",                                      //签发者
		"iat": time.Now().Unix(),                                        //签发时间
		"jti": "9527",                                                   //jwt的唯一身份标识，主要用来作为一次性token,从而回避重放攻击。
		"exp": time.Now().Add(10 * time.Hour * time.Duration(1)).Unix(), //过期时间
	})

	tokenString, _ := token.SignedString([]byte(jwtKey))
	fmt.Println("签发时间：", time.Now().Unix())
	fmt.Println("到期时间：", time.Now().Add(10*time.Hour*time.Duration(1)).Unix())
	return tokenString
}
`

var logrus = `package middleware

import (
	"{{.Appname}}/models"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"runtime"
	"strings"
	"time"
)


func LogMiddle(ctx iris.Context) {
	appName := viper.GetString("app_name")
	traceLogger := models.NewLogger()
	traceLogger.Proto = ctx.Request().Proto
	traceLogger.Referer = ctx.Request().Referer()
	traceLogger.ReqTime = time.Now()
	traceLogger.Length = ctx.Request().ContentLength
	traceLogger.UserAgent = ctx.Request().UserAgent()
	traceLogger.ReqUri = ctx.Request().RequestURI
	traceLogger.ReqMethod = ctx.Request().Method
	logrus.AddHook(NewContextHook(appName, traceLogger))
	logrus.Info("这是一条测试")
	ctx.Next()
}

/**
添加一个自定义的钩子，实现在日志中增加一个字段appName和请求的基本信息，
及其日志rotate策略
*/

type AppHook struct {
	Skip        int
	AppName     string
	TraceLogger *models.TraceLogger
}

func NewContextHook(appName string, logger *models.TraceLogger) *AppHook {
	return &AppHook{AppName: appName, TraceLogger: logger, Skip: 5}
}

func (h *AppHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *AppHook) Fire(entry *logrus.Entry) error {
	entry.Data["app"] = h.AppName
	entry.Data["proto"] = h.TraceLogger.Proto
	entry.Data["method"] = h.TraceLogger.ReqMethod
	entry.Data["uri"] = h.TraceLogger.ReqUri
	entry.Data["user-agent"] = h.TraceLogger.UserAgent
	entry.Data["length"] = h.TraceLogger.Length
	entry.Data["referer"] = h.TraceLogger.Referer
	entry.Data["line"] = findCaller(h.Skip)
	return nil
}

func getCaller(skip int) (string, int) {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "", 0
	}
	n := 0
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			n++
			if n >= 2 {
				file = file[i+1:]
				break
			}
		}
	}
	return file, line
}

func findCaller(skip int) string {
	file := ""
	line := 0
	for i := 0; i < 10; i++ {
		file, line = getCaller(skip + i)
		if !strings.HasPrefix(file, "logrus") {
			break
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}`
