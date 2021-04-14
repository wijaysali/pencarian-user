package app

import (
	"pencarian_user/server/controller"
	"pencarian_user/server/middleware"
)

func route() {
	router.Use(middleware.CORSMiddleware()) //to enable api request between client and server

	router.POST("/username", controller.Username)
}
