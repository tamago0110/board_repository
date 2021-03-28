package handler

import (
	"bytes"
	"encoding/json"
	"golang_api/auth"
	"golang_api/model"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	newUser           model.User
	newUserPro        model.Profile
	jwtNewUser        auth.JWT
	newUserToOldUser1 model.Lead
	imageDir          string
	imgDirPath        string
)

func TestSignup(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	rowBody := map[string]interface{}{
		"email":    "test@gmail.com",
		"password": "test",
	}
	payload, payloadErr := json.Marshal(rowBody)
	assert.NoError(t, payloadErr)
	reqBody := bytes.NewBuffer(payload)

	r.POST("/signup", Signup(database))
	req, reqErr := http.NewRequest("POST", "/signup", reqBody)
	assert.NoError(t, reqErr)

	req.Header.Add("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	resBodyErr := json.Unmarshal(w.Body.Bytes(), &newUser)
	assert.NoError(t, resBodyErr)
	assert.Equal(t, rowBody["email"], newUser.Email)
}

func TestLogin(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	rowBody := map[string]interface{}{
		"email":    "test@gmail.com",
		"password": "test",
	}
	payload, payloadErr := json.Marshal(rowBody)
	assert.NoError(t, payloadErr)
	reqBody := bytes.NewBuffer(payload)

	r.POST("/login", Login(database, tokenSecret, tokenIss))

	req, reqErr := http.NewRequest("POST", "/login", reqBody)
	assert.NoError(t, reqErr)

	req.Header.Add("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	resBodyErr := json.Unmarshal(w.Body.Bytes(), &jwtNewUser)
	assert.NoError(t, resBodyErr)

	authInfo, authErr := auth.Parse(jwtNewUser.Token, tokenSecret)
	assert.NoError(t, authErr)

	assert.Equal(t, tokenIss, authInfo.Iss)
	assert.Equal(t, newUser.ID, authInfo.UserID)
}

func TestPostProfile(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	rowBody := map[string]interface{}{
		"name": "test",
	}
	payload, payloadErr := json.Marshal(rowBody)
	assert.NoError(t, payloadErr)
	reqBody := bytes.NewBuffer(payload)

	user := r.Group("/user").Use(auth.Authz(tokenSecret, tokenIss))
	{
		user.POST("/profile", PostProfile(database))
	}

	req, reqErr := http.NewRequest("POST", "/user/profile", reqBody)
	assert.NoError(t, reqErr)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+jwtNewUser.Token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	resBodyErr := json.Unmarshal(w.Body.Bytes(), &newUserPro)
	assert.NoError(t, resBodyErr)
	assert.Equal(t, newUser.ID, newUserPro.UserID)
	assert.Equal(t, rowBody["name"], newUserPro.Name)
}

func TestGetMyProfile(t *testing.T) {
	var resultPro model.Profile

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	user := r.Group("/user").Use(auth.Authz(tokenSecret, tokenIss))
	{
		user.GET("/me", GetMyProfile(database))
	}

	req, reqErr := http.NewRequest("GET", "/user/me", nil)
	assert.NoError(t, reqErr)

	req.Header.Add("Authorization", "Bearer "+jwtNewUser.Token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	resBodyErr := json.Unmarshal(w.Body.Bytes(), &resultPro)
	assert.NoError(t, resBodyErr)
	assert.Equal(t, newUserPro.ID, resultPro.ID)
	assert.Equal(t, newUserPro.UserID, resultPro.UserID)
	assert.Equal(t, newUserPro.Name, resultPro.Name)
	assert.Equal(t, newUserPro.Image, resultPro.Image)
}

func TestGetProfile(t *testing.T) {
	var resultPro model.Profile

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	user := r.Group("/user").Use(auth.Authz(tokenSecret, tokenIss))
	{
		user.GET("/profile/:uuid", GetProfile(database))
	}

	req, reqErr := http.NewRequest("GET", "/user/profile/"+newUser.ID, nil)
	assert.NoError(t, reqErr)

	req.Header.Add("Authorization", "Bearer "+jwtNewUser.Token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	resBodyErr := json.Unmarshal(w.Body.Bytes(), &resultPro)
	assert.NoError(t, resBodyErr)
	assert.Equal(t, newUserPro.ID, resultPro.ID)
	assert.Equal(t, newUserPro.UserID, resultPro.UserID)
	assert.Equal(t, newUserPro.Name, resultPro.Name)
	assert.Equal(t, newUserPro.Image, resultPro.Image)
}

func TestPostLead(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	rowBody := map[string]interface{}{
		"producer": oldUser1.ID,
	}
	payload, payloadErr := json.Marshal(rowBody)
	assert.NoError(t, payloadErr)
	reqBody := bytes.NewBuffer(payload)

	user := r.Group("/user").Use(auth.Authz(tokenSecret, tokenIss))
	{
		user.POST("/lead", PostLead(database))
	}

	req, reqErr := http.NewRequest("POST", "/user/lead", reqBody)
	assert.NoError(t, reqErr)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+jwtNewUser.Token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	resBodyErr := json.Unmarshal(w.Body.Bytes(), &newUserToOldUser1)
	assert.NoError(t, resBodyErr)
	assert.Equal(t, newUser.ID, newUserToOldUser1.Consumer)
	assert.Equal(t, oldUser1.ID, newUserToOldUser1.Producer)
}

func TestListProfile(t *testing.T) {
	var leadPros []model.Profile

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	user := r.Group("/user").Use(auth.Authz(tokenSecret, tokenIss))
	{
		user.GET("/profile", ListProfile(database))
	}

	req, reqErr := http.NewRequest("GET", "/user/profile", nil)
	assert.NoError(t, reqErr)

	req.Header.Add("Authorization", "Bearer "+jwtOldUser1.Token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	resBodyErr := json.Unmarshal(w.Body.Bytes(), &leadPros)
	assert.NoError(t, resBodyErr)
	assert.Equal(t, newUserPro.ID, leadPros[0].ID)
	assert.Equal(t, newUserPro.UserID, leadPros[0].UserID)
	assert.Equal(t, newUserPro.Name, leadPros[0].Name)
	assert.Equal(t, newUserPro.Image, leadPros[0].Image)
}

func TestPutProfile(t *testing.T) {
	var newUserOldPro model.Profile
	newUserOldPro.ID = newUserPro.ID
	newUserOldPro.UserID = newUserPro.UserID
	newUserOldPro.Name = newUserPro.Name
	newUserOldPro.Image = newUserPro.Image

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	imageDir = "http://server/image/"
	imgDirPath = "../image/"

	user := r.Group("/user").Use(auth.Authz(tokenSecret, tokenIss))
	{
		user.PUT("/me", PutProfile(database, imageDir, imgDirPath))
	}

	body := &bytes.Buffer{}
	mw := multipart.NewWriter(body)
	nameField, nameFieldErr := mw.CreateFormField("name")
	if assert.NoError(t, nameFieldErr) {
		_, err := nameField.Write([]byte("newName1"))
		assert.NoError(t, err)
	}
	mwCloseErr := mw.Close()
	assert.NoError(t, mwCloseErr)

	req, reqErr := http.NewRequest("PUT", "/user/me", body)
	assert.NoError(t, reqErr)

	req.Header.Add("Content-Type", mw.FormDataContentType())
	req.Header.Add("Authorization", "Bearer "+jwtNewUser.Token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	resBodyErr := json.Unmarshal(w.Body.Bytes(), &newUserPro)
	assert.NoError(t, resBodyErr)
	assert.Equal(t, newUserOldPro.ID, newUserPro.ID)
	assert.Equal(t, newUserOldPro.UserID, newUserPro.UserID)
	assert.Equal(t, "newName1", newUserPro.Name)
	assert.Equal(t, newUserOldPro.Image, newUserPro.Image)
}

func TestPutProfile2(t *testing.T) {
	var newUserOldPro model.Profile
	newUserOldPro.ID = newUserPro.ID
	newUserOldPro.UserID = newUserPro.UserID
	newUserOldPro.Name = newUserPro.Name
	newUserOldPro.Image = newUserPro.Image

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	user := r.Group("/user").Use(auth.Authz(tokenSecret, tokenIss))
	{
		user.PUT("/me", PutProfile(database, imageDir, imgDirPath))
	}

	newFileName := "test1.png"

	body := &bytes.Buffer{}
	mw := multipart.NewWriter(body)
	nameField, nameFieldErr := mw.CreateFormField("name")
	if assert.NoError(t, nameFieldErr) {
		_, err := nameField.Write([]byte("newName2"))
		assert.NoError(t, err)
	}
	imgField, imgFieldErr := mw.CreateFormFile("image", newFileName)
	if assert.NoError(t, imgFieldErr) {
		_, err := imgField.Write([]byte(newFileName))
		assert.NoError(t, err)
	}
	mwCloseErr := mw.Close()
	assert.NoError(t, mwCloseErr)

	req, reqErr := http.NewRequest("PUT", "/user/me", body)
	assert.NoError(t, reqErr)

	req.Header.Add("Content-Type", mw.FormDataContentType())
	req.Header.Add("Authorization", "Bearer "+jwtNewUser.Token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	resBodyErr := json.Unmarshal(w.Body.Bytes(), &newUserPro)
	assert.NoError(t, resBodyErr)
	assert.Equal(t, newUserOldPro.ID, newUserPro.ID)
	assert.Equal(t, newUserOldPro.UserID, newUserPro.UserID)
	assert.Equal(t, "newName2", newUserPro.Name)
	assert.Equal(t, imageDir+newUser.ID+"_"+newFileName, newUserPro.Image)
}

func TestPutProfile3(t *testing.T) {
	var newUserOldPro model.Profile
	newUserOldPro.ID = newUserPro.ID
	newUserOldPro.UserID = newUserPro.UserID
	newUserOldPro.Name = newUserPro.Name
	newUserOldPro.Image = newUserPro.Image

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	user := r.Group("/user").Use(auth.Authz(tokenSecret, tokenIss))
	{
		user.PUT("/me", PutProfile(database, imageDir, imgDirPath))
	}

	newFileName := "test2.png"

	body := &bytes.Buffer{}
	mw := multipart.NewWriter(body)
	nameField, nameFieldErr := mw.CreateFormField("name")
	if assert.NoError(t, nameFieldErr) {
		_, err := nameField.Write([]byte("newName3"))
		assert.NoError(t, err)
	}
	imgField, imgFieldErr := mw.CreateFormFile("image", newFileName)
	if assert.NoError(t, imgFieldErr) {
		_, err := imgField.Write([]byte(newFileName))
		assert.NoError(t, err)
	}
	mwCloseErr := mw.Close()
	assert.NoError(t, mwCloseErr)

	req, reqErr := http.NewRequest("PUT", "/user/me", body)
	assert.NoError(t, reqErr)

	req.Header.Add("Content-Type", mw.FormDataContentType())
	req.Header.Add("Authorization", "Bearer "+jwtNewUser.Token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	resBodyErr := json.Unmarshal(w.Body.Bytes(), &newUserPro)
	assert.NoError(t, resBodyErr)
	assert.Equal(t, newUserOldPro.ID, newUserPro.ID)
	assert.Equal(t, newUserOldPro.UserID, newUserPro.UserID)
	assert.Equal(t, "newName3", newUserPro.Name)
	assert.Equal(t, imageDir+newUser.ID+"_"+newFileName, newUserPro.Image)

	newImage, newImgOpenErr := os.Open(imgDirPath + newUser.ID + "_" + newFileName)
	assert.NoError(t, newImgOpenErr)
	newImgCloseErr := newImage.Close()
	assert.NoError(t, newImgCloseErr)

	_, oldImgOpenErr := os.Open(imgDirPath + newUser.ID + "_test1.png")
	assert.Error(t, oldImgOpenErr)
}
