package mysql

import (
	"fmt"
	"time"

	mysqlLib "github.com/go-sql-driver/mysql"
	log "github.com/rohanchauhan02/common/logs"

	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql" // register mysql with dd-trace
	gormtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorm.io/gorm.v1"
)

var (
	logger = log.NewCommonLog()
)

type Mysql interface {
	InitClient() error
	GetClient() *gorm.DB
}

type MysqlConfig struct {
	DatabaseHost       string
	DatabaseUser       string
	DatabasePassword   string
	DatabasePort       string
	DatabaseName       string
	DatabaseDebug      bool
	MaxIddleConns      int
	MaxOpenConns       int
	ConnMaxLifetime    time.Duration
	DatadogServiceName string
}

type mysql struct {
	databaseHost       string
	databaseUser       string
	databasePassword   string
	databasePort       string
	databaseName       string
	databaseDebug      bool
	maxIddleConns      int
	maxOpenConns       int
	connMaxLifetime    time.Duration
	datadogServiceName string
	client             *gorm.DB
}

// NewMysql is a factory that implement of mysql database configuration
func NewMysql(config MysqlConfig) Mysql {
	return &mysql{
		databaseHost:       config.DatabaseHost,
		databaseUser:       config.DatabaseUser,
		databasePassword:   config.DatabasePassword,
		databasePort:       config.DatabasePort,
		databaseName:       config.DatabaseName,
		databaseDebug:      config.DatabaseDebug,
		maxIddleConns:      config.MaxIddleConns,
		maxOpenConns:       config.MaxOpenConns,
		connMaxLifetime:    config.ConnMaxLifetime,
		datadogServiceName: config.DatadogServiceName,
	}
}

func (m *mysql) InitClient() error {

	logger.Info("Start open mysql connection...")

	// Register augments the provided driver with tracing, enabling it to be loaded by gormtrace.Open.
	sqltrace.Register("mysql", &mysqlLib.MySQLDriver{}, sqltrace.WithServiceName(m.datadogServiceName))
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		m.databaseUser,
		m.databasePassword,
		m.databaseHost,
		m.databasePort,
		m.databaseName,
	)

	gormConfig := &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	}
	if m.databaseDebug {
		gormConfig.Logger = gormLogger.Default.LogMode(gormLogger.Info)
	}

	gormDB, err := gormtrace.Open(gormMysql.Open(connString), gormConfig)
	if err != nil {
		return err
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return err
	}

	if m.maxIddleConns != 0 {
		sqlDB.SetMaxIdleConns(m.maxIddleConns)
	} else {
		sqlDB.SetMaxIdleConns(8)
	}

	if m.maxOpenConns != 0 {
		sqlDB.SetMaxOpenConns(m.maxOpenConns)
	} else {
		sqlDB.SetMaxOpenConns(16)
	}

	if m.connMaxLifetime != 0 {
		sqlDB.SetConnMaxLifetime(m.connMaxLifetime)
	} else {
		sqlDB.SetConnMaxLifetime(1 * time.Hour)
	}

	m.client = gormDB

	return nil
}

func (m *mysql) GetClient() *gorm.DB {
	return m.client
}
