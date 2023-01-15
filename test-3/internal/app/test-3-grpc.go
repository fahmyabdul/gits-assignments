package app

import (
	"strings"
	"time"

	"github.com/fahmyabdul/gits-assignments/test-3/internal/usecase/repo"

	"github.com/fahmyabdul/gits-assignments/test-3/pkg/redispkg"

	"github.com/fahmyabdul/gits-assignments/test-3/config"
	"github.com/fahmyabdul/gits-assignments/test-3/internal/controllers/grpc"
	"github.com/fahmyabdul/gits-assignments/test-3/internal/usecase/impl"
	"github.com/fahmyabdul/gits-assignments/test-3/pkg/logger"
	"github.com/fahmyabdul/gits-assignments/test-3/pkg/pggorm"
)

func RunTest3Grpc(appVersion string, config *config.ConfigStruct, logPath, serviceName string) {
	// Set Logger
	log := logger.New(logger.Config{
		Level:                 config.Logger.Level,
		ConsoleLoggingEnabled: config.Logger.ConsoleLog,
		FileLoggingEnabled:    config.Logger.FileLog,
		EncodeLogsAsJson:      config.Logger.LogAsJson,
		Directory:             logPath,
		Filename:              strings.ToLower(serviceName) + ".log",
		MaxSize:               config.Logger.MaxSize,
		MaxBackups:            config.Logger.MaxBackups,
		MaxAge:                config.Logger.MaxAge,
	})

ConnectPostgres:
	pgConn, err := pggorm.NewPostgreGORM(
		config.Databases.Postgre.Host,
		config.Databases.Postgre.Port,
		config.Databases.Postgre.User,
		config.Databases.Postgre.Pass,
		config.Databases.Postgre.DB,
		config.Databases.Postgre.Schema,
	)
	if err != nil {
		log.Error().Err(err).Msg("Unable to connect to PostgreDB, retry in 5 seconds")
		time.Sleep(5 * time.Second)
		goto ConnectPostgres
	}

ConnectRedis:
	redisConn, err := redispkg.NewRedisPkg(
		config.Databases.Redis.Host,
		config.Databases.Redis.Auth,
		config.Databases.Redis.DB,
		config.Databases.Redis.MaxIdle,
		config.Databases.Redis.MaxActive,
	)
	if err != nil {
		log.Error().Err(err).Msg("Unable to connect to Redis, retry in 5 seconds")
		time.Sleep(5 * time.Second)
		goto ConnectRedis
	}

	repoAuthor := repo.NewRepoAuthorImpl(pgConn, redisConn)
	repoPublisher := repo.NewRepoPublisherImpl(pgConn, redisConn)
	repoBook := repo.NewRepoBookImpl(pgConn, redisConn)

StartGrpc:
	err = grpc.NewGrpcCtrl(
		log,
		&config.Controllers.Grpc,
		grpc.SetUsecase("author", impl.NewUsecaseAuthorImpl(repoAuthor, log)),
		grpc.SetUsecase("publisher", impl.NewUsecasePublisherImpl(repoPublisher, log)),
		grpc.SetUsecase("book", impl.NewUsecaseBookImpl(repoBook, log)),
	)
	if err != nil {
		log.Error().Err(err).Msg("Unable to start Grpc Server, retry in 5 seconds")
		time.Sleep(5 * time.Second)
		goto StartGrpc
	}
}
