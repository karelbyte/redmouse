package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"sync"
)

type driverProperties struct {
	name string
	strConect string
}

type drivers map[string]*driverProperties

var (
	dialectors drivers

	driver string

	user string

	password string

	host string

	port string

	database string

	strConect string

	dsn string

	err error

	once sync.Once

	instance *gorm.DB
)

func (p *driverProperties) getDialect(host string, user string, password string, database string, port string) gorm.Dialector {
	var dial gorm.Dialector
	dsn := ""
	switch p.name {
	case "mysql":
		dsn = fmt.Sprintf(p.strConect, user, password, host, port, database)
		dial = mysql.Open(dsn)
	case "postgres":
		dsn = fmt.Sprintf(p.strConect, host, user, password, database, port)
		dial = postgres.Open(dsn)
	}
	return dial
}

func driverConection() gorm.Dialector {

	driver = os.Getenv("DBDRIVER")

	user = os.Getenv("DBUSER")

	password = os.Getenv("DBPASSWORD")

	host = os.Getenv("DBHOST")

	port = os.Getenv("DBPORT")

	database = os.Getenv("DBNAME")

	dialectors := make(drivers)

	dialectors["mysql"] = &driverProperties{
		name:      "mysql",
		strConect: "%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
	}

	dialectors["postgres"] = &driverProperties{
		name:      "postgres",
		strConect: "host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
	}

	dialect := dialectors[driver]

	return dialect.getDialect(host, user, password, database, port)
  
}


func Conect() *gorm.DB {

	once.Do(func() {
		instance, err = gorm.Open(driverConection(), &gorm.Config{})
	})

	if err != nil {
		fmt.Print("Error: " + err.Error())
		os.Exit(1)
	}

	return instance
}