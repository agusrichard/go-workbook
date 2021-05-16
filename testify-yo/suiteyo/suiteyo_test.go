package suiteyo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type YoTestSuite struct {
	suite.Suite
	StartingVariale int64
}

func (suite *YoTestSuite) SetupTest() {
	fmt.Println("SetupTest")
	suite.StartingVariale = 21
}

func (suite *YoTestSuite) TearDownTest() {
	fmt.Println("TearDown")
	suite.StartingVariale = 0
}

func (suite *YoTestSuite) TestOne() {
	suite.Equal(int64(21), suite.StartingVariale)
}

func (suite *YoTestSuite) TestTwo() {
	suite.StartingVariale = 99
	suite.Equal(int64(99), suite.StartingVariale)
}

func TestYoTestSuite(t *testing.T) {
	suite.Run(t, new(YoTestSuite))
}
