package db

import (
	"fmt"
	"flag"
	"io/ioutil"
	"log"
	"path/filepath"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	
	
	basePath := filepath.Base("")

	pathToScripts := filepath.Join(basePath, "db", "scripts")

	err = execAllSqlScripts(conn, pathToScripts)
	if err != nil {
		log.Fatalf("Can't exec scripts: %v", err)
	}
}

func execAllSqlScripts(
	conn *sqlx.DB, 
	pathToScripts string,
) error {
	dir, err := ioutil.ReadDir(pathToScripts)
	if err != nil {
		return err
	}

	for _, script := range dir {
		err := execSqlScript(conn, filepath.Join(pathToScripts, script.Name() ) )
		if err != nil {
			return err
		}
	}

	return nil
}

func execSqlScript(conn *sqlx.DB, pathToScript string) error {
	data, err := ioutil.ReadFile(pathToScript)
	if err != nil {
		return err
	}

	if _, err := conn.Exec(string(data)); err != nil {
		return err
	}

	return nil
}

// if change db change this method
func connectToDB() (*sqlx.DB, error) {
	conn, err := sqlx.Connect("mysql", connectionString)
	if err != nil {
		for i := 0; i < 10; i++ {
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
		&gorm.Config{},
	)
}