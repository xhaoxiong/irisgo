package config

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type Config struct {
	Name string
}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	if err := c.initConfig(); err != nil {
		return err
	}
	initLog()
	return nil

}

func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name)
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}

	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("irisgo")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}
func initLog() {
	WithMaxAge := viper.GetDuration("log.with_max_age")
	WithRotationTime := viper.GetDuration("log.with_rotation_time")
	WithRotationCount := viper.GetInt("log.with_rotation_count")
	LogName := viper.GetString("log.logger_file")

	writer, err := rotatelogs.New(
		LogName+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(LogName),                // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(WithMaxAge),               // 文件最大保存时间
		rotatelogs.WithRotationTime(WithRotationTime),   // 日志切割时间间隔
		rotatelogs.WithRotationCount(WithRotationCount), //设置文件清理前最多保存的个数。
	)
	writer1 := os.Stdout
	if err!=nil{
		log.Fatalf("create file log failed:%v", err)
	}
	/**
	设置日志输入在文件
	writer2, err := os.OpenFile(LogName, os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0755)
	if err != nil {
		log.Fatalf("create file log.txt failed:%v", err)
	}*/


	//设置在输出日志中添加文件名和方法信息
	//logrus.SetReportCaller(true)

	//设置输出写入
	logrus.SetOutput(io.MultiWriter(writer, writer1))
	logrus.SetLevel(logrus.DebugLevel)
	//可设置json，txt，nested（需要引入github.com/antonfisher/nested-logrus-formatter）等格式
	//logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.SetFormatter(&nested.Formatter{
		FieldsOrder:           []string{"component", "category"},
		HideKeys:              false, //是否隐藏键值
		NoColors:              false, //是否显示颜色
		NoFieldsColors:        false,
		ShowFullLevel:         false,
		TrimMessages:          false,
		CallerFirst:           false,
		CustomCallerFormatter: nil,
		TimestampFormat:       time.RFC3339, //格式化时间
	})
}
