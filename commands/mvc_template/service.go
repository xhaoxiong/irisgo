/**
 * @Author: xiaoxiao
 * @Description:
 * @File:  service
 * @Version: 1.0.0
 * @Date: 2020/5/12 2:11 下午
 */
package mvc_template

var service = `package service

import (
	"{{.Appname}}/repositories"
)

type TestService struct {
	repo *repositories.TestRepositories
}

func NewTestService() *TestService {
	return &TestService{repo: repositories.NewTestRepositories()}
}`
