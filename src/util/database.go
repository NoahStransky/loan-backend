package util

import (
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

var engine *xorm.Engine

const retryTimes = 10

func CreateConnection(driveName, host, port, user, password, database string) *xorm.Engine {
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, database)

	engine, err := xorm.NewEngine(driveName, dbinfo)
	if err != nil {
		panic(err)
	}
	if err := engine.Ping(); err != nil {
		for i := 0; i < retryTimes; i++ {
			fmt.Println("reconnecting... for ", i+1, " times")
			engine, err = xorm.NewEngine(driveName, dbinfo)
			if err != nil {
				if i == 10 {
					panic(err)
				}
				time.Sleep(1 * time.Second)
				continue
			}
			break
		}
	}
	engine.ShowSQL(false)
	engine.SetConnMaxLifetime(0)
	engine.SetMaxIdleConns(10)
	engine.SetMaxOpenConns(10)
	return engine
}
