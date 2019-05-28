package initial

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var conf Config
var db *gorm.DB

// Config ...
type Config struct {
	Database database
}

func init() {

	if _, err := toml.DecodeFile("../configs/config.toml", &conf); err != nil {
		panic(err)
	}

	dbConn, err := gorm.Open(conf.Database.Provider,
		fmt.Sprintf("user=%v password=%v dbname=%v host=%v port=%v sslmode=disable",
			conf.Database.User,
			conf.Database.Password,
			conf.Database.Name,
			conf.Database.Server,
			conf.Database.Port,
		),
	)

	if err != nil {
		panic(err)
	}

	dbConn.DB().SetMaxIdleConns(20)
	dbConn.DB().SetMaxOpenConns(50)

	db = dbConn

	fmt.Println("database is running ....")
}

// GetDB ...
func GetDB() *gorm.DB {
	// db = db.LogMode(true)
	return db
}
