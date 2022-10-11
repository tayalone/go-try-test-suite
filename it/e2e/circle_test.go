//go:build integration
// +build integration

package e2e

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"go-try-test-suite/router"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type CircleTestSuite struct {
	suite.Suite
	Router *gin.Engine
}

func (suite *CircleTestSuite) SetupSuite() {
	suite.Router = router.SetUpRouter()
}

func (suite *CircleTestSuite) TestSuccess() {
	w := httptest.NewRecorder()
	body := strings.NewReader(`{"r": 7}`)
	req, _ := http.NewRequest(http.MethodPost, "/circle", body)
	suite.Router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
	suite.JSONEq(`{"desc": {
		  "area": 153.93804002589985,
		  "perimeter": 43.982297150257104
		},
		"message": "OK"
	  }`, w.Body.String())
}

func (suite *CircleTestSuite) TestBadReq() {
	w := httptest.NewRecorder()
	body := strings.NewReader(`{"r": "7"}`)
	req, _ := http.NewRequest(http.MethodPost, "/circle", body)
	suite.Router.ServeHTTP(w, req)

	suite.Equal(http.StatusBadRequest, w.Code)
	suite.JSONEq(`{
		"error":"json: cannot unmarshal string into Go struct field circleBody.r of type float64"
	  }`, w.Body.String())
}

func TestRouteCircleSuite(t *testing.T) {
	suite.Run(t, new(CircleTestSuite))
}
