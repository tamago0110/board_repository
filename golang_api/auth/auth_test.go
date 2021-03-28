package auth

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	testUserID string
	testSecret string
	testIss    string
	testToken  string
	err        error
)

func TestCreateToken(t *testing.T) {
	testUserID = "boardGolangTypescriptReactHooksReduxDocker"
	testSecret = "golangGolangGolandJetBrains"
	testIss = "Gopher"
	testToken, err = CreateToken(testUserID, testSecret, testIss)
	assert.NoError(t, err)
}

func TestParse(t *testing.T) {
	var auth *Auth
	auth, err = Parse(testToken, testSecret)
	assert.NoError(t, err)
	assert.Equal(t, testUserID, auth.UserID)
	assert.Equal(t, testIss, auth.Iss)
}

func TestAuthz(t *testing.T) {
	var dummy Dummy

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	r.Use(Authz(testSecret, testIss))
	r.GET("/dummy", dummyFunc)

	req, reqErr := http.NewRequest("GET", "/dummy", nil)
	assert.NoError(t, reqErr)

	req.Header.Add("Authorization", "Bearer "+testToken)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	resBodyErr := json.Unmarshal(w.Body.Bytes(), &dummy)
	assert.NoError(t, resBodyErr)
	assert.Equal(t, testUserID, dummy.UserID)
}

func TestAuthz2(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	r.Use(Authz(testSecret, testIss))
	r.GET("/dummy", dummyFunc)

	req, reqErr := http.NewRequest("GET", "/dummy", nil)
	assert.NoError(t, reqErr)

	r.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Code)
}

func TestAuthz3(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	r.Use(Authz(testSecret, testIss))
	r.GET("/dummy", dummyFunc)

	req, reqErr := http.NewRequest("GET", "/dummy", nil)
	assert.NoError(t, reqErr)

	req.Header.Add("Authorization", "Token "+testToken)
	r.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Code)
}

func TestAuthz4(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	r.Use(Authz(testSecret, testIss))
	r.GET("/dummy", dummyFunc)

	req, reqErr := http.NewRequest("GET", "/dummy", nil)
	assert.NoError(t, reqErr)

	invalidSecretKey := "invalid"
	invalidSecretToken, invalidSecretErr := CreateToken(testUserID, invalidSecretKey, testIss)
	assert.NoError(t, invalidSecretErr)

	req.Header.Add("Authorization", "Bearer "+invalidSecretToken)
	r.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Code)
}

func TestAuthz5(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	r.Use(Authz(testSecret, testIss))
	r.GET("/dummy", dummyFunc)

	req, reqErr := http.NewRequest("GET", "/dummy", nil)
	assert.NoError(t, reqErr)

	invalidIss := "invalid"
	invalidIssToken, invalidIssErr := CreateToken(testUserID, testSecret, invalidIss)
	assert.NoError(t, invalidIssErr)

	req.Header.Add("Authorization", "Bearer "+invalidIssToken)
	r.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Code)
}
