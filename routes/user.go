package routes

import (
	"DemoApiPostgre/controllers"
	"gorm.io/gorm"
)

func UserRoute(db *gorm.DB) {
	router := r.Group("/user")
	{
		router.GET("", controllers.GetAllUser(db))
		router.POST("", controllers.CreateUser(db))
	}
}
