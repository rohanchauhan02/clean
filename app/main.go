package main

import (
	"fmt"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	mysqlLib "github.com/rohanchauhan02/clean/common/database/mysql"
	datadogLib "github.com/rohanchauhan02/clean/common/datadog"
	transporterLib "github.com/rohanchauhan02/clean/common/transporter"
	"github.com/rohanchauhan02/clean/intenal/config"
	redislib "github.com/rohanchauhan02/common/database/redis"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	// datadog libs
	echoDatadog "gopkg.in/DataDog/dd-trace-go.v1/contrib/labstack/echo"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {
	e := echo.New()
	cfg := config.NewImmutableConfig()

	redis := redislib.NewRedis(redislib.RedisConfig{
		Host:     cfg.GetRedis().Host,
		Password: cfg.GetRedis().Password,
	})
	err := redis.InitClient()
	if err != nil {
		e.Logger.Errorf("Failed to open redis connection: %s", err.Error())
	}

	transporterClient, err := transporterLib.NewClient(&transporterLib.ClientOptions{
		Provider:        "aws",
		AccessKeyID:     cfg.GetAWS().AccessKey,
		AccessKeySecret: cfg.GetAWS().SecretKey,
		Region:          cfg.GetAWS().Region,
	})
	if err != nil {
		msgError := fmt.Sprintf("Failed to create transporter client: %s", err.Error())
		e.Logger.Errorf(msgError)
		panic(msgError)
	}
	mysql := mysqlLib.NewMysql(mysqlLib.MysqlConfig{
		DatabaseHost:     cfg.GetDB().Host,
		DatabaseUser:     cfg.GetDB().User,
		DatabasePassword: cfg.GetDB().Password,
		DatabasePort:     cfg.GetDB().Port,
		DatabaseName:     cfg.GetDB().Name,
		MaxIddleConns:    cfg.GetDB().MaxIdleConns,
		MaxOpenConns:     cfg.GetDB().MaxOpenConns,
		ConnMaxLifetime:  time.Duration(cfg.GetDB().MaxLifetimeConns * int(time.Minute)),
		DatabaseDebug:    true,
	})
	err = mysql.InitClient()
	if err != nil {
		msgError := fmt.Sprintf("Failed to open mysql connection: %s", err.Error())
		e.Logger.Errorf(msgError)
		panic(msgError)
	}

	datadog, err := datadogLib.NewDatadogClient(datadogLib.Config{
		Namespace: cfg.GetDatadog().Namespace,
		Env:       cfg.GetDatadog().ServiceEnv,
		Unit:      cfg.GetDatadog().Unit,
		Host:      cfg.GetDatadog().Host,
	})
	if err != nil {
		e.Logger.Errorf("Failed to create datadog client: %s", err.Error())
	}
	tracer.Start(
		tracer.WithServiceName(cfg.GetDatadog().ServiceName),
		tracer.WithEnv(cfg.GetDatadog().ServiceEnv),
	)
	defer tracer.Stop()

	// Adds the datadog middleware that is run after the router is done
	// IMPORTANT: add this before the application context middleware code
	e.Use(echoDatadog.Middleware(echoDatadog.WithServiceName(cfg.GetDatadog().ServiceName)))
	// register middlewares
	e.Use(middleware.Logger())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Gzip())
	e.Use(middleware.CORS())
	e.Use(middlewareLib.MiddlewareRequestID())
	e.HTTPErrorHandler = errorLib.ErrorHandler
}
