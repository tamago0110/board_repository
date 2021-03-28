package handler

import (
	"bytes"
	"encoding/json"
	"golang_api/auth"
	"golang_api/model"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	display model.Display
)

func TestPostBoard1(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	rowBody := map[string]interface{}{
		"title":   "test1",
		"content": "test1",
	}

	payload, payloadErr := json.Marshal(rowBody)
	assert.NoError(t, payloadErr)
	reqBody := bytes.NewBuffer(payload)

	board := r.Group("/board").Use(auth.Authz(tokenSecret, tokenIss))
	{
		board.POST("/item", PostBoard(database))
	}

	req, reqErr := http.NewRequest("POST", "/board/item", reqBody)
	assert.NoError(t, reqErr)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+jwtOldUser1.Token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	resBodyErr := json.Unmarshal(w.Body.Bytes(), &board1)
	assert.NoError(t, resBodyErr)
	assert.Equal(t, oldUser1.ID, board1.CreatedBy)
	assert.Equal(t, rowBody["title"], board1.Title)
	assert.Equal(t, rowBody["content"], board1.Content)
}

func TestPostBoard2(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	rowBody := map[string]interface{}{
		"title":   "test2",
		"content": "test2",
	}

	payload, payloadErr := json.Marshal(rowBody)
	assert.NoError(t, payloadErr)
	reqBody := bytes.NewBuffer(payload)

	board := r.Group("/board").Use(auth.Authz(tokenSecret, tokenIss))
	{
		board.POST("/item", PostBoard(database))
	}

	req, reqErr := http.NewRequest("POST", "/board/item", reqBody)
	assert.NoError(t, reqErr)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+jwtOldUser1.Token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	resBodyErr := json.Unmarshal(w.Body.Bytes(), &board2)
	assert.NoError(t, resBodyErr)
	assert.Equal(t, oldUser1.ID, board2.CreatedBy)
	assert.Equal(t, rowBody["title"], board2.Title)
	assert.Equal(t, rowBody["content"], board2.Content)
}

func TestGetMyBoard(t *testing.T) {
	var myBoards []model.Board

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	board := r.Group("/board").Use(auth.Authz(tokenSecret, tokenIss))
	{
		board.GET("/mine", GetMyBoard(database))
	}

	req, reqErr := http.NewRequest("GET", "/board/mine", nil)
	assert.NoError(t, reqErr)

	req.Header.Add("Authorization", "Bearer "+jwtOldUser1.Token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	resBodyErr := json.Unmarshal(w.Body.Bytes(), &myBoards)
	assert.NoError(t, resBodyErr)
	assert.Equal(t, 2, len(myBoards))
	assert.Equal(t, oldUser1.ID, myBoards[0].CreatedBy)
	assert.Equal(t, oldUser1.ID, myBoards[1].CreatedBy)
}

func TestPutBoard(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	rowBody := map[string]interface{}{
		"title":   "test3",
		"content": "test3",
	}

	payload, payloadErr := json.Marshal(rowBody)
	assert.NoError(t, payloadErr)
	reqBody := bytes.NewBuffer(payload)

	board := r.Group("/board").Use(auth.Authz(tokenSecret, tokenIss))
	{
		board.PUT("/item/:id", PutBoard(database))
	}

	req, reqErr := http.NewRequest("PUT", "/board/item/"+strconv.Itoa(int(board1.ID)), reqBody)
	assert.NoError(t, reqErr)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+jwtOldUser1.Token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	resBodyErr := json.Unmarshal(w.Body.Bytes(), &board1)
	assert.NoError(t, resBodyErr)
	assert.Equal(t, oldUser1.ID, board1.CreatedBy)
	assert.Equal(t, rowBody["title"], board1.Title)
	assert.Equal(t, rowBody["content"], board1.Content)
}

func TestRejectBoard(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	rowBody := map[string]interface{}{
		"board_id":    board2.ID,
		"rejected_by": oldUser3.ID,
	}

	payload, payloadErr := json.Marshal(rowBody)
	assert.NoError(t, payloadErr)
	reqBody := bytes.NewBuffer(payload)

	board := r.Group("/board").Use(auth.Authz(tokenSecret, tokenIss))
	{
		board.POST("/reject", RejectBoard(database))
	}

	req, reqErr := http.NewRequest("POST", "/board/reject", reqBody)
	assert.NoError(t, reqErr)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+jwtOldUser3.Token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	resBodyErr := json.Unmarshal(w.Body.Bytes(), &display)
	assert.NoError(t, resBodyErr)
	assert.Equal(t, oldUser3.ID, display.RejectedBy)
	assert.Equal(t, rowBody["board_id"], display.BoardID)
}

func TestListBoard(t *testing.T) {
	var boardsForOldUser2 []model.Board

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	board := r.Group("/board").Use(auth.Authz(tokenSecret, tokenIss))
	{
		board.GET("/item", ListBoard(database))
	}

	req, reqErr := http.NewRequest("GET", "/board/item", nil)
	assert.NoError(t, reqErr)

	req.Header.Add("Authorization", "Bearer "+jwtOldUser2.Token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	resBodyErr := json.Unmarshal(w.Body.Bytes(), &boardsForOldUser2)
	assert.NoError(t, resBodyErr)
	assert.Equal(t, 0, len(boardsForOldUser2))
}

func TestListBoard2(t *testing.T) {
	var boardsForOldUser3 []model.Board

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	board := r.Group("/board").Use(auth.Authz(tokenSecret, tokenIss))
	{
		board.GET("/item", ListBoard(database))
	}

	req, reqErr := http.NewRequest("GET", "/board/item", nil)
	assert.NoError(t, reqErr)

	req.Header.Add("Authorization", "Bearer "+jwtOldUser3.Token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	resBodyErr := json.Unmarshal(w.Body.Bytes(), &boardsForOldUser3)
	assert.NoError(t, resBodyErr)
	assert.Equal(t, 1, len(boardsForOldUser3))
}

func TestDeleteBoard(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	board := r.Group("/board").Use(auth.Authz(tokenSecret, tokenIss))
	{
		board.DELETE("/item/:id", DeleteBoard(database))
	}

	req, reqErr := http.NewRequest("DELETE", "/board/item/"+strconv.Itoa(int(board1.ID)), nil)
	assert.NoError(t, reqErr)

	req.Header.Add("Authorization", "Bearer "+jwtOldUser1.Token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestDeleteBoard2(t *testing.T) {
	var boards []model.Board

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	board := r.Group("/board").Use(auth.Authz(tokenSecret, tokenIss))
	{
		board.DELETE("/item/:id", DeleteBoard(database))
	}

	req, reqErr := http.NewRequest("DELETE", "/board/item/"+strconv.Itoa(int(board2.ID)), nil)
	assert.NoError(t, reqErr)

	req.Header.Add("Authorization", "Bearer "+jwtOldUser1.Token)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	findBoardErr := database.Where("created_by = ?", oldUser1.ID).Find(&boards).Error
	assert.NoError(t, findBoardErr)
	assert.Equal(t, 0, len(boards))
}
