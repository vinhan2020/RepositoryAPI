package controllers

import (
	"DemoApiPostgre/business"
	"DemoApiPostgre/models"
	"DemoApiPostgre/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetAllUser(db *gorm.DB) gin.HandlerFunc {
	dbPostgre := repository.NewPostgreRepo(db)
	userService := business.NewUserService(dbPostgre)
	return func(c *gin.Context) {
		c.Header("content-type", "application/json")
		users, err := userService.GetAllUser()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	}
}

func CreateUser(db *gorm.DB) gin.HandlerFunc {
	dbPostgre := repository.NewPostgreRepo(db)
	userService := business.NewUserService(dbPostgre)
	return func(c *gin.Context) {
		var (
			user models.User
		)
		_ = c.BindJSON(&user)
		c.Header("content-type", "application/json")
		userReps, err := userService.CreateUser(user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"user": userReps,
		})
	}
}
