package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	r *gin.Engine
)

func Route(db *gorm.DB) {
	r = gin.New()
	gin.SetMode(gin.ReleaseMode)
	UserRoute(db)
	_ = r.Run(":8080")
}
