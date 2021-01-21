package db

import (
	"flag"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)
var (
	dbUser 		= flag.String("dbUser", "docker", "need to connect db")
	dbPassword	= flag.String("dbPassword", "docker", "need to connect to db")
	network		= flag.String("network", "service.auth.db:3306", "need to connect db")
	dbName		= flag.String("dbName", "auth", "database name")
)
// string to connect to DataBase
var connectionString = "docker:docker@tcp(service.auth.db:3306)/auth"

// Init db package and provide connection
func Init() {
	flag.Parse()
	connectionString = fmt.Sprintf("%s:%s@tcp(%s)/%s", *dbUser, *dbPassword, *network, *dbName)
	log.Println(connectionString)

	conn, err := connectToDB()
	if err != nil {
		log.Fatalf("Can't connect to db: %v", err)
	}
	defer conn.Close()
}

// if change db change this method
// wait 30 seconds and crash
func connectToDB() (*sqlx.DB, error) {
	conn, err := sqlx.Connect("mysql", connectionString)
	if err != nil {
		for i := 0; i < 30; i++ {
			t := time.NewTimer(time.Second)
			<-t.C
			conn, err := sqlx.Connect("mysql", connectionString)
			if err != nil {
				continue
			} else {
				return conn, nil
			}
		}
		return nil, err
	}
	return conn, nil
}


// ConnectToDB return a pointer to database
func ConnectToDB() (*sqlx.DB, error) {
	return connectToDB()
}

// Open .....
func Open() (*sqlx.DB, error) {
	return sqlx.Open("mysql", connectionString)
}

// GormOpen open DB by a gorm
func GormOpen() (*gorm.DB, error){
	sqlDB, err := Open()
	if err != nil {
		return nil, err
	}
	return gorm.Open( 
		mysql.New(
			mysql.Config{Conn: sqlDB},
			), 
		&gorm.Config{ Logger: logger.Default.LogMode(logger.Silent)},
	)
}

// TODO refactor package