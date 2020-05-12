/**
 * @Author: xiaoxiao
 * @Description:
 * @File:  repo
 * @Version: 1.0.0
 * @Date: 2020/5/12 2:11 下午
 */
package mvc_template

var repositories = `package repositories

import (
	"{{.Appname}}/models"
	"github.com/jinzhu/gorm"
)

type TestRepositories struct {
	db *gorm.DB
}

func NewTestRepositories() *TestRepositories {
	return &TestRepositories{db: models.DB.Mysql}
}`
