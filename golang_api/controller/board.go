package controller

import (
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
	"golang_api/model"
)

func CreateBoard(user *model.User, board *model.Board, db *gorm.DB) error {
	board.CreatedBy = user.ID
	if err := db.Model(user).Association("Boards").Append(board).Error; err != nil {
		return err
	}
	return nil
}

func SelectBoardWhereId(id uint, board *model.Board, db *gorm.DB) error {
	if err := db.Where("id = ?", id).First(board); err != nil {
		return err.Error
	}
	return nil
}

func CreateDisplayOnUser(rejectedBy *model.User, display *model.Display, db *gorm.DB) error {
	display.RejectedBy = rejectedBy.ID
	if err := db.Model(rejectedBy).Association("Displays").Append(display); err != nil {
		return err.Error
	}
	return nil
}

func CreateDisplayOnBoard(board *model.Board, display *model.Display, db *gorm.DB) error {
	if err := db.Model(board).Association("Displays").Append(display); err != nil {
		return err.Error
	}
	return nil
}

func SelectNotDisplayUsers(userID string, notDisplayUsers *[]model.NotDisplayUser, db *gorm.DB) error {
	if err := db.Table("leads").Where("consumer = ?", userID).Find(&notDisplayUsers); err != nil {
		return err.Error
	}
	return nil
}

func SelectNotDisplayBoards(userID string, notDisplayBoards *[]model.NotDisplayBoard, db *gorm.DB) error {
	if err := db.Table("displays").Where("rejected_by = ?", userID).Find(notDisplayBoards); err != nil {
		return err.Error
	}
	return nil
}

func CreateDisplayBoardsQuery(notDisplayUserID []string, notDisplayBoardID []uint) (string, []interface{}, error) {
	query, args, err := sqlx.In("SELECT * FROM boards boards WHERE created_by NOT IN (?) AND id NOT IN (?)", notDisplayUserID, notDisplayBoardID)
	if err != nil {
		return "", nil, err
	}
	return query, args, nil
}

func SelectDisplayBoards(query string, args []interface{}, boards *[]model.Board, db *gorm.DB) error {
	if err := db.Raw(query, args...).Find(boards); err != nil {
		return err.Error
	}
	return nil
}

func SelectBoardsWhereCreatorID(userID string, boards *[]model.Board, db *gorm.DB) error {
	if err := db.Where("created_by = ?", userID).Find(boards); err != nil {
		return err.Error
	}
	return nil
}

func SelectOldBoardWhereStrID(searchID string, board *model.Board, db *gorm.DB) error {
	if err := db.Where("id = ?", searchID).First(board); err != nil {
		return err.Error
	}
	return nil
}

func UpdateBoard(oldBoard *model.Board, newBoard *model.Board, db *gorm.DB) error {
	newBoard.ID = oldBoard.ID
	newBoard.CreatedBy = oldBoard.CreatedBy
	if err := db.Model(oldBoard).Update(newBoard); err != nil {
		return err.Error
	}
	return nil
}

func RemoveBoard(board *model.Board, db *gorm.DB) error {
	if err := db.Delete(board); err != nil {
		return err.Error
	}
	return nil
}
