package utils

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	config "shouyindemo/conf"
	"strconv"
)

type mySqlUtil struct {
	DB *gorm.DB
}

//全局变量, 外部使用utils.MySqlClient来访问
var MySqlClient mySqlUtil

func InitMySqlUtil(config config.DatabaseConfig) error {
	dsn := buildDatabaseDSN(config)
	fmt.Println(dsn)
	dd, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		panic("[SetupDefaultDatabase#newConnection error]: " + err.Error() + " " + dsn)
	}

	sqlDB, err := dd.DB()
	if err != nil {
		panic("[SetupDefaultDatabase#newConnection error]: " + err.Error() + " " + dsn)
	}
	if config.MaxOpenConns > 0 {
		sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	}
	if config.MaxIdleConns > 0 {
		sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	}

	fmt.Printf("\nDefault atabase connection successful: %s\n", dsn)
	//初始化全局redis结构体
	MySqlClient = mySqlUtil{DB: dd}
	return nil
}

func buildDatabaseDSN(config config.DatabaseConfig) string {
	switch config.Driver {
	case "mysql":
		return fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?%s",
			config.User,
			config.Password,
			config.Host,
			strconv.Itoa(config.Port),
			config.DbName,
			config.Options,
		)
	case "postgres":
		return fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s options='%s'",
			config.Host,
			strconv.Itoa(config.Port),
			config.User,
			config.DbName,
			config.Password,
			config.Options,
		)
	case "sqlite3":
		return config.DbName
	case "mssql":
		return fmt.Sprintf(
			"sqlserver://%s:%s@%s:%s?database=%s&%s",
			config.User,
			config.Password,
			config.Host,
			strconv.Itoa(config.Port),
			config.DbName,
			config.Options,
		)
	}

	panic("DB Connection not supported:" + config.Driver)
	return ""
}
