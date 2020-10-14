package service

import (
	"test/repositories"
)

type TestService struct {
	repo *repositories.TestRepositories
}

func NewTestService() *TestService {
	return &TestService{repo: repositories.NewTestRepositories()}
}