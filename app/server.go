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
	config.Setup()
	config := config.GetConfig()

	db, err := database.PrepareDatabase()
	if err != nil {
		panic(err)
	}

	logger := logger.NewLogger(config)

	middleware := middleware.InitMiddleware(*config, logger)
	dom := repo.InitRepo(db, logger)
	uc := usecase.Init(dom, middleware, logger)
	hndlr := handler.Init(uc, logger)

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
		logger.Infof("error running server - %s", err.Error())
	}
}
