package utils

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CommonUtilsTest struct {
	suite.Suite
}

func TestCommonUtils(t *testing.T) {
	suite.Run(t, new(CommonUtilsTest))
}

func (t *CommonUtilsTest) SetupTest() {
}

func (t *CommonUtilsTest) TestIsDuplicatedStringTrue() {
	testIsDuplicatedStringTrue(t.T(), []string{"1", "1", "2"})
	testIsDuplicatedStringTrue(t.T(), []string{"1", "1"})
}

func testIsDuplicatedStringTrue(t *testing.T, check []string) {
	actual := IsDuplicatedString(check)

	assert.True(t, actual)
}

func (t *CommonUtilsTest) TestIsDuplicatedStringFalse() {
	testIsDuplicatedStringFalse(t.T(), []string{"1", "2", "3"})
	testIsDuplicatedStringFalse(t.T(), []string{})
}

func testIsDuplicatedStringFalse(t *testing.T, check []string) {
	actual := IsDuplicatedString(check)

	assert.False(t, actual)
}
