package userRepositoryTests

import (
	"testing"
	"twit/configs"
	"twit/models"
	"twit/repositories"
	"twit/servers"
	"twit/utils"

	"github.com/stretchr/testify/suite"
)

type RepositoryRegisterUserSuite struct {
	suite.Suite
	repository      repositories.UserRepository
	cleanupExecutor utils.TruncateTableExecutor
}

func (suite *RepositoryRegisterUserSuite) SetupTest() {
	repository := servers.SetupRepositories().UserRepository

	suite.repository = repository

	cleanupExecutor := utils.InitTruncateTableExecutor(configs.DB)
	suite.cleanupExecutor = cleanupExecutor
}

func (suite *RepositoryRegisterUserSuite) TearDownTest() {
	defer suite.cleanupExecutor.TruncateTable([]string{"users"})
}

func (suite *RepositoryRegisterUserSuite) TestRegisterUserEmptyEmail() {
	user := models.User{
		Password: "password",
		Username: "username",
	}
	err := suite.repository.RegisterUser(user)
	suite.Error(err)
}

func (suite *RepositoryRegisterUserSuite) TestRegisterUserEmptyPassword() {
	user := models.User{
		Email:    "email",
		Username: "username",
	}
	err := suite.repository.RegisterUser(user)
	suite.Error(err)
}

func (suite *RepositoryRegisterUserSuite) TestRegisterUserEmptyUsername() {
	user := models.User{
		Email:    "email",
		Password: "password",
	}
	err := suite.repository.RegisterUser(user)
	suite.Error(err)
}

func TestRepositoryRegisterUserSuite(t *testing.T) {
	suite.Run(t, new(RepositoryRegisterUserSuite))
}
