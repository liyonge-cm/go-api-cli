package mysql

import (
	"context"
	"fmt"
	"time"

	"go-api-cli-prj/config"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

type MySQL struct {
	ctx    context.Context
	config *config.MySQL
	Logger *zap.Logger
	DB     *gorm.DB
}

func NewMySQL(ctx context.Context, config *config.MySQL) *MySQL {
	return &MySQL{
		ctx:    ctx,
		config: config,
	}
}

func (s *MySQL) WithLogger(log *zap.Logger) {
	s.Logger = log
}

func (s *MySQL) ConnDB() {
	DSN := &config.MySQL{
		Endpoint: s.config.Endpoint,
		Username: s.config.Username,
		Password: s.config.Password,
		Database: s.config.Database,
	}
	s.DB = s.connect(s.dsn(DSN))
	DB = s.DB
}
func (s *MySQL) connect(dsn string) *gorm.DB {
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //使用单数表名
		},
	}
	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if nil != err {
		s.Logger.Fatal("failed to connect database", zap.String("dsn", dsn), zap.Error(err))
	}

	sqlDB, err := db.DB()
	if nil != err {
		s.Logger.Fatal("failed to connect database", zap.String("dsn", dsn), zap.Error(err))
	}

	sqlDB.SetMaxIdleConns(s.config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(s.config.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db.Debug()
}

func (s *MySQL) dsn(dsn *config.MySQL) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dsn.Username,
		dsn.Password,
		dsn.Endpoint,
		dsn.Database,
	)
}
