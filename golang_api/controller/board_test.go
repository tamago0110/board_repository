package controller

import (
	"github.com/stretchr/testify/assert"
	"golang_api/model"
	"strconv"
	"testing"
)

var (
	display                  model.Display
	notDisplayUsersToUser2   []model.NotDisplayUser
	notDisplayBoardsToUser2  []model.NotDisplayBoard
	notDisplayUsersToUser3   []model.NotDisplayUser
	notDisplayBoardsToUser3  []model.NotDisplayBoard
	notDisplayUserIDToUser2  []string
	notDisplayBoardIDToUser2 []uint
	notDisplayUserIDToUser3  []string
	notDisplayBoardIDToUser3 []uint
	query1                   string
	args1                    []interface{}
	queryErr1                error
	query2                   string
	args2                    []interface{}
	queryErr2                error
	boards                   []model.Board
)

func TestCreateBoard(t *testing.T) {
	board1.Title = "board1"
	board1.Content = "test board1"
	board2.Title = "board2"
	board2.Content = "test board2"

	err1 := CreateBoard(&oldUser1, &board1, database)
	err2 := CreateBoard(&oldUser1, &board2, database)
	assert.NoError(t, err1)
	assert.NoError(t, err2)
}

func TestSelectBoardWhereId(t *testing.T) {
	var board model.Board
	err := SelectBoardWhereId(board1.ID, &board, database)
	assert.NoError(t, err)
	assert.Equal(t, board1.ID, board.ID)
	assert.Equal(t, board1.CreatedBy, board.CreatedBy)
	assert.Equal(t, board1.Title, board.Title)
	assert.Equal(t, board1.Content, board.Content)
}

func TestCreateDisplayOnUser(t *testing.T) {
	display.BoardID = board2.ID
	err := CreateDisplayOnUser(&oldUser3, &display, database)
	assert.NoError(t, err)
}

func TestCreateDisplayOnBoard(t *testing.T) {
	err := CreateDisplayOnBoard(&board2, &display, database)
	assert.NoError(t, err)
}

func TestSelectNotDisplayUsers(t *testing.T) {
	err1 := SelectNotDisplayUsers(oldUser2.ID, &notDisplayUsersToUser2, database)
	err2 := SelectNotDisplayUsers(oldUser3.ID, &notDisplayUsersToUser3, database)
	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.Equal(t, oldUser1.ID, notDisplayUsersToUser2[0].Producer)
	assert.Equal(t, 0, len(notDisplayUsersToUser3))
}

func TestSelectNotDisplayBoards(t *testing.T) {
	err1 := SelectNotDisplayBoards(oldUser2.ID, &notDisplayBoardsToUser2, database)
	err2 := SelectNotDisplayBoards(oldUser3.ID, &notDisplayBoardsToUser3, database)
	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.Equal(t, 0, len(notDisplayBoardsToUser2))
	assert.Equal(t, board2.ID, notDisplayBoardsToUser3[0].BoardID)
}

func TestCreateDisplayBoardsQuery(t *testing.T) {
	for _, v := range notDisplayUsersToUser2 {
		notDisplayUserIDToUser2 = append(notDisplayUserIDToUser2, v.Producer)
	}
	notDisplayUserIDToUser2 = append(notDisplayUserIDToUser2, oldUser2.ID)
	for _, v := range notDisplayBoardsToUser2 {
		notDisplayBoardIDToUser2 = append(notDisplayBoardIDToUser2, v.BoardID)
	}
	if len(notDisplayBoardIDToUser2) == 0 {
		notDisplayBoardIDToUser2 = append(notDisplayBoardIDToUser2, 0)
	}
	query1, args1, queryErr1 = CreateDisplayBoardsQuery(notDisplayUserIDToUser2, notDisplayBoardIDToUser2)

	for _, v := range notDisplayUsersToUser3 {
		notDisplayUserIDToUser3 = append(notDisplayUserIDToUser3, v.Producer)
	}
	notDisplayUserIDToUser3 = append(notDisplayUserIDToUser3, oldUser2.ID)
	for _, v := range notDisplayBoardsToUser3 {
		notDisplayBoardIDToUser3 = append(notDisplayBoardIDToUser3, v.BoardID)
	}
	if len(notDisplayBoardIDToUser3) == 0 {
		notDisplayBoardIDToUser3 = append(notDisplayBoardIDToUser3, 0)
	}
	query2, args2, queryErr2 = CreateDisplayBoardsQuery(notDisplayUserIDToUser3, notDisplayBoardIDToUser3)

	assert.NoError(t, queryErr1)
	assert.NoError(t, queryErr2)
}

func TestSelectDisplayBoards(t *testing.T) {
	var displayBoardsToUser2 []model.Board
	var displayBoardsToUser3 []model.Board
	err1 := SelectDisplayBoards(query1, args1, &displayBoardsToUser2, database)
	err2 := SelectDisplayBoards(query2, args2, &displayBoardsToUser3, database)
	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.Equal(t, 0, len(displayBoardsToUser2))
	assert.Equal(t, 1, len(displayBoardsToUser3))
	assert.Equal(t, board1.ID, displayBoardsToUser3[0].ID)
	assert.Equal(t, board1.CreatedBy, displayBoardsToUser3[0].CreatedBy)
	assert.Equal(t, board1.Title, displayBoardsToUser3[0].Title)
	assert.Equal(t, board1.Content, displayBoardsToUser3[0].Content)
}

func TestSelectBoardsWhereCreatorID(t *testing.T) {
	err := SelectBoardsWhereCreatorID(oldUser1.ID, &boards, database)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(boards))
}

func TestSelectOldBoardWhereStrID(t *testing.T) {
	searchID := strconv.Itoa(int(board1.ID))
	var board model.Board
	err := SelectOldBoardWhereStrID(searchID, &board, database)
	assert.NoError(t, err)
	assert.Equal(t, board1.CreatedBy, board.CreatedBy)
	assert.Equal(t, board1.Title, board.Title)
	assert.Equal(t, board1.Content, board.Content)
}

func TestUpdateBoard(t *testing.T) {
	var newBoard model.Board
	newBoard.Title = "new"
	newBoard.Content = "new"
	err := UpdateBoard(&board1, &newBoard, database)
	assert.NoError(t, err)
	assert.Equal(t, board1.ID, newBoard.ID)
	assert.Equal(t, board1.CreatedBy, newBoard.CreatedBy)
}

func TestRemoveBoard(t *testing.T) {
	err1 := RemoveBoard(&board1, database)
	err2 := RemoveBoard(&board2, database)
	assert.NoError(t, err1)
	assert.NoError(t, err2)
}
