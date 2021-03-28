package handler

import (
	"fmt"
	"golang_api/controller"
	"golang_api/model"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func PostBoard(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var creator model.User
		var board model.Board

		if err := c.ShouldBindJSON(&board); err != nil {
			c.String(http.StatusBadRequest, "Request failed: "+err.Error())
			return
		}

		userID, ok := c.Get("userID")
		if !ok {
			c.String(http.StatusUnauthorized, "You are not authorized user.")
			return
		}
		strUserID := fmt.Sprintf("%v", userID)

		if err := controller.SelectUserWhereId(strUserID, &creator, db); err != nil {
			c.String(http.StatusBadRequest, "Request failed: "+err.Error())
			return
		}

		if err := controller.CreateBoard(&creator, &board, db); err != nil {
			c.String(http.StatusBadRequest, "Request failed: "+err.Error())
			return
		}

		c.JSON(http.StatusOK, board)
	}
}

func RejectBoard(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var rejectedBy model.User
		var board model.Board
		var display model.Display

		if err := c.ShouldBindJSON(&display); err != nil {
			c.String(http.StatusBadRequest, "Request failed: "+err.Error())
			return
		}

		userID, ok := c.Get("userID")
		if !ok {
			c.String(http.StatusUnauthorized, "You are not authorized user.")
			return
		}
		strUserID := fmt.Sprintf("%v", userID)

		if err := controller.SelectUserWhereId(strUserID, &rejectedBy, db); err != nil {
			c.String(http.StatusBadRequest, "Request failed: "+err.Error())
			return
		}

		if err := controller.SelectBoardWhereId(display.BoardID, &board, db); err != nil {
			c.String(http.StatusBadRequest, "Request failed: "+err.Error())
			return
		}

		if err := controller.CreateDisplayOnUser(&rejectedBy, &display, db); err != nil {
			c.String(http.StatusBadRequest, "Request failed: "+err.Error())
			return
		}

		if err := controller.CreateDisplayOnBoard(&board, &display, db); err != nil {
			c.String(http.StatusBadRequest, "Request failed: "+err.Error())
			return
		}

		c.JSON(http.StatusOK, display)
	}
}

func ListBoard(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var notDisplayUsers []model.NotDisplayUser
		var notDisplayBoards []model.NotDisplayBoard
		var boards []model.Board

		userID, ok := c.Get("userID")
		if !ok {
			c.String(http.StatusUnauthorized, "You are not authorized user.")
			return
		}
		strUserID := fmt.Sprintf("%v", userID)

		if err := controller.SelectNotDisplayUsers(strUserID, &notDisplayUsers, db); err != nil {
			c.String(http.StatusBadRequest, "Request is failed A: "+err.Error())
			return
		}
		var notDisplayUserID []string
		for _, v := range notDisplayUsers {
			notDisplayUserID = append(notDisplayUserID, v.Producer)
		}
		notDisplayUserID = append(notDisplayUserID, strUserID)

		if err := controller.SelectNotDisplayBoards(strUserID, &notDisplayBoards, db); err != nil {
			c.String(http.StatusBadRequest, "Request is failed A: "+err.Error())
			return
		}
		var notDisplayBoardID []uint
		for _, v := range notDisplayBoards {
			notDisplayBoardID = append(notDisplayBoardID, v.BoardID)
		}
		if len(notDisplayBoardID) == 0 {
			notDisplayBoardID = append(notDisplayBoardID, 0)
		}

		query, args, queryErr := controller.CreateDisplayBoardsQuery(notDisplayUserID, notDisplayBoardID)
		if queryErr != nil {
			c.String(http.StatusBadRequest, "Request is failed A: "+queryErr.Error())
			return
		}

		if err := controller.SelectDisplayBoards(query, args, &boards, db); err != nil {
			c.String(http.StatusBadRequest, "Request is failed A: "+err.Error())
			return
		}

		c.JSON(http.StatusOK, boards)
	}
}

func GetMyBoard(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var boards []model.Board

		userID, ok := c.Get("userID")
		if !ok {
			c.String(http.StatusUnauthorized, "You are not authorized user.")
			return
		}
		strUserID := fmt.Sprintf("%v", userID)

		if err := controller.SelectBoardsWhereCreatorID(strUserID, &boards, db); err != nil {
			c.String(http.StatusBadRequest, "Request is failed A: "+err.Error())
			return
		}

		c.JSON(http.StatusOK, boards)
	}
}

func PutBoard(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var oldBoard model.Board
		var newBoard model.Board

		userID, ok := c.Get("userID")
		if !ok {
			c.String(http.StatusUnauthorized, "You are not authorized user.")
			return
		}
		strUserID := fmt.Sprintf("%v", userID)

		searchID := c.Param("id")

		if err := controller.SelectOldBoardWhereStrID(searchID, &oldBoard, db); err != nil {
			c.String(http.StatusBadRequest, "Request is failed A: "+err.Error())
			return
		}

		if oldBoard.CreatedBy != strUserID {
			c.String(http.StatusUnauthorized, "Request failed: You can't edit this board")
			return
		}

		if err := c.ShouldBindJSON(&newBoard); err != nil {
			c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
			return
		}

		if err := controller.UpdateBoard(&oldBoard, &newBoard, db); err != nil {
			c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
			return
		}

		c.JSON(http.StatusOK, newBoard)
	}
}

func DeleteBoard(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var board model.Board

		userID, ok := c.Get("userID")
		if !ok {
			c.String(http.StatusUnauthorized, "You are not authorized user.")
			return
		}
		strUserID := fmt.Sprintf("%v", userID)

		searchID := c.Param("id")

		if err := controller.SelectOldBoardWhereStrID(searchID, &board, db); err != nil {
			c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
			return
		}

		if board.CreatedBy != strUserID {
			c.String(http.StatusUnauthorized, "Request failed: You can't delete this board")
			return
		}

		if err := controller.RemoveBoard(&board, db); err != nil {
			c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Delete successfully"})
	}
}
