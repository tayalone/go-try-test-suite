//go:build integration
// +build integration

package e2e

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"go-try-test-suite/router"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type PingTestSuite struct {
	suite.Suite
	Router *gin.Engine
}

func (suite *PingTestSuite) SetupSuite() {
	suite.Router = router.SetUpRouter()
}

// // Make sure that VariableThatShouldStartAtFive is set to five
// // before each test
// func (suite *PingTestSuite) SetupTest() {
// 	fmt.Println("--------- SetupTest --------------")
// 	suite.VariableThatShouldStartAtFive = 5
// }

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *PingTestSuite) TestSuccess() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
	suite.Router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
	suite.JSONEq(`{"message":"pong"}`, w.Body.String())
}

func TestRoutePingSuite(t *testing.T) {
	suite.Run(t, new(PingTestSuite))
}
