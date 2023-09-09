package app

import (
	"xm/app/router"
	"xm/config"
	database "xm/db"
	"xm/internal/handler"
	"xm/internal/middleware"
	"xm/internal/repo"
	"xm/internal/usecase"
	"xm/logger"
)

func Start() {
	config := config.GetConfig()

	db, err := database.PrepareDatabase()
	if err != nil {
		panic(err)
	}

	logger.InitLogger(config)

	middleware := middleware.InitMiddleware(*config)
	dom := repo.InitRepo(db)
	uc := usecase.Init(dom, middleware)
	hndlr := handler.Init(uc)

	router := router.PrepareRouter(&router.Capsule{
		DB:         db,
		Repo:       dom,
		Usecase:    uc,
		Handler:    hndlr,
		Config:     config,
		Middleware: middleware,
	})

	logger.Infof("server running at port %s", config.ServerPort)
	err = router.Run(":" + config.ServerPort)
	if err != nil {
		logger.Fatalf("error running server - %s", err.Error())
	}
}
