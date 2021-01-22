package db

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBManger *Manager

func init() {
	DBManger = NewManager(
		"root",
		"root",
		"db:3306",
		20 * time.Second,
	)
}

// Manager help us  to make  connection to database
type Manager struct {
	// look like "user:password@tcp(adresss:port)/"
	connectionStr 	string
	
	// wait n*seconds if can't  connect return error
	waitTimeConn	time.Duration

	// dbNameList		[]string
}

/* 
NewManager return a new manager that can provide connection to current db
dbListJSONsrcFile should look like:
	{
		"names": 
		[
			"name_1",
			"name_2",
		]
	}
*/
func NewManager(
	dbUser 				string,
	dbPassword  		string,
	network				string,
	waitTimeConn		time.Duration,
) (*Manager) {

	// type names struct {
	// 	List []string `json:"names"`
	// }
	// n := &names{}

	manager := &Manager{
		connectionStr: fmt.Sprintf("%s:%s@tcp(%s)/", dbUser, dbPassword, network),
		waitTimeConn: waitTimeConn,
	}

	// block of code for dbNameList
	// data, err := ioutil.ReadAll(dbListJSONsrcFile)
	// if err != nil {
	// 	return nil, err
	// }

	// err = json.Unmarshal(data, n)
	// if err != nil {
	// 	return nil, err
	// }

	// manager.databaseNames = n.List
	// n = nil

	// if len(manager.databaseNames) == 0 {
	// 	return nil, errors.New("databaseNames is  empty")
	// }

	return manager
}

// OpenDataBase return a pointer to GormDataBase
func (m *Manager) OpenDataBase(dbName string) (*gorm.DB, error) {
	sqlDB, err := connectToDBOrWait("mysql", m.connectionStr+dbName, m.waitTimeConn)
	if err != nil {
		return nil, err
	}

	return gorm.Open(
		mysql.New(
			mysql.Config{Conn: sqlDB},
		),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		},
	)
}

func connectToDBOrWait(
	driverName, 
	connectionString string, 
	waitTime time.Duration,
) (*sqlx.DB, error) {
	conn, err := sqlx.Connect(driverName, connectionString)
	if err != nil {
		for i := 0; i < int(waitTime.Seconds()); i++ {
			t := time.NewTimer(time.Second)
			<-t.C
			conn, err := sqlx.Connect(driverName, connectionString)
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