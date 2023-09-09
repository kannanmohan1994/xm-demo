package app

import (
	"fmt"
	"os"
	"xm/app/router"
	"xm/config"
	"xm/consts"
	database "xm/db"
	"xm/internal/handler"
	"xm/internal/middleware"
	kafkaNotifier "xm/internal/notifiers/kafka"
	"xm/internal/repo"
	"xm/internal/usecase"
	"xm/logger"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Start() {
	config.Setup()
	config := config.GetConfig()

	db, err := database.PrepareDatabase()
	if err != nil {
		panic(err)
	}

	logger := logger.NewLogger(config)

	publisher, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%s", config.KafkaHost, config.KafkaPort),
	})
	if err != nil {
		logger.Errorf("failed to create producer: %s\n", err)
		panic(err)
	}

	defer publisher.Close()

	go consumer(config.KafkaHost, config.KafkaPort)

	notifier := kafkaNotifier.InitKafka(publisher)

	middleware := middleware.InitMiddleware(*config, logger)
	dom := repo.InitRepo(db, logger)
	uc := usecase.Init(dom, middleware, logger, notifier)
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

func consumer(host, port string) {
	fmt.Println("Consumer is starting...")

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%s", host, port),
		"group.id":          consts.KafkaGroupId,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create consumer: %v\n", err)
		os.Exit(1)
	}

	defer consumer.Close()

	err = consumer.SubscribeTopics([]string{consts.KafkaEventNotifierTopic}, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error subscribing to topic: %v\n", err)
		os.Exit(1)
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			fmt.Printf("Consumer error: %v\n", err)
			continue
		}
		fmt.Printf("Received message: %s\n", string(msg.Value))
	}
}
