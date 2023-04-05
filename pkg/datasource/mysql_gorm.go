package datasource

import (
	"errors"
	"fmt"
	configPkg "github.com/Adesubomi/lightning-node-manager/pkg/config"
	logPkg "github.com/Adesubomi/lightning-node-manager/pkg/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"reflect"
)

const (
	POSTGRES string = "postgres"
	MYSQL    string = "mysql"
)

func ConnectDatabase(dbConfig *configPkg.DatabaseConfig) *gorm.DB {

	connection := dbConfig.Connection
	if connection == POSTGRES {
		postgresConn, err := postgresConnection(dbConfig)
		if err != nil {
			log.Fatal(err)
		}
		logPkg.PrintlnGreen("  ✔ PostGreSQL Connection Established")
		return postgresConn
	} else if connection == MYSQL {
		mysqlConn, err := mysqlConnection(dbConfig)
		if err != nil {
			log.Fatal(err)
		}
		logPkg.PrintlnGreen("  ✔ MySQL Connection Established")
		return mysqlConn
	}

	log.Fatal(
		errors.New("could not connect to any database"),
	)
	return nil
}

func MigrateTables(dbClient *gorm.DB, entities []interface{}) {
	for _, entity := range entities {
		err := dbClient.AutoMigrate(entity)

		if err != nil {
			msg := fmt.Sprintf(
				"    ✗ Migration Failed [%v] - %v",
				reflect.TypeOf(entity).Elem().Name(),
				err.Error(),
			)
			logPkg.PrintlnRed(msg)
		} else {
			message := fmt.Sprintf(
				"    ✔ Migrated [%v]",
				reflect.TypeOf(entity).Elem().Name())
			logPkg.PrintlnGreen(message)
		}
	}
}

func mysqlConnection(dbConfig *configPkg.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DbName,
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf(" ?? Connection to database failed: %v\n", err)
		return database, err
	}

	return database, nil
}

func postgresConnection(dbConfig *configPkg.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Africa/Lagos",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DbName,
		dbConfig.Port,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return database, err
	}

	return database, nil
}
