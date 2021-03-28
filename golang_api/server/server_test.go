package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRouter(t *testing.T) {
	var db *gorm.DB
	var checkEngine *gin.Engine
	r := Router(db, "test", "test", "test", "test")
	assert.NotEmpty(t, r)
	assert.IsType(t, r, checkEngine)
}
