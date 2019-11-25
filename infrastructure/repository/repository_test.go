package repository

import (
	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/tockn/emukone/infrastructure/repository/mysql"
	"log"
	"os"
	"time"
)

var (
	dbconf = "../../dbconfig.yml"
	env    = "test"
)

func InitConn() *gorm.DB {
	db := getConn()
	if _, err := db.DB().Exec(`SET FOREIGN_KEY_CHECKS = 0`); err != nil {
		panic(err)
	}

	//var tables []string
	//dbx := sqlx.NewDb(db.DB(), db.Dialect().GetName())
	//if err := dbx.Select(&tables, `SHOW TABLES`); err != nil {
	//	panic(err)
	//}
	//for _, t := range tables {
	//	if _, err := db.DB().Exec(fmt.Sprintf("TRUNCATE TABLE %s", t)); err != nil {
	//		panic(err)
	//	}
	//}

	if _, err := db.DB().Exec(`SET FOREIGN_KEY_CHECKS = 1`); err != nil {
		panic(err)
	}

	return db
}

func getConn() *gorm.DB {
	conf, err := GetDBConf()
	if err != nil {
		panic(err)
	}
	conn, err := gorm.Open("mysql", conf.Datasource)
	if err != nil {
		panic(err)
	}
	file, _ := os.OpenFile("../../mysql.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	conn.LogMode(true)
	conn.SetLogger(log.New(file, "[test]", 0))
	return conn
}

func getConnAndDSN() (*gorm.DB, string, error) {
	conf, err := GetDBConf()
	if err != nil {
		panic(err)
	}
	conn, err := gorm.Open("mysql", conf.Datasource)
	if err != nil {
		panic(err)
	}
	file, _ := os.OpenFile("../../mysql.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	conn.LogMode(true)
	conn.SetLogger(log.New(file, "[test]", 0))

	return getConn(), conf.Datasource, nil
}

func getConnAndMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	conn, err := gorm.Open("mysql", db)
	if err != nil {
		return nil, nil, err
	}
	file, _ := os.OpenFile("../../mysql.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	conn.LogMode(true)
	conn.SetLogger(log.New(file, "[test]", 0))

	return conn, mock, nil
}

func GetDBConf() (*mysql.DBConfig, error) {
	confs, err := mysql.NewDBConfigsFromFile(dbconf)
	if err != nil {
		return nil, err
	}

	if os.Getenv("CI") == "true" {
		env = "ci/test"
	}
	return confs[env], nil
}

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}
