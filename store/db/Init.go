package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	migrateSql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/prince-bansal/go-otp/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"path/filepath"
)

var db *gorm.DB

type Store struct {
	Config *config.Config
	Db     *gorm.DB
}

func NewStore(config *config.Config) *Store {
	c := config.Db
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true",
		c.User,
		c.Password,
		c.Host,
		c.Database,
	)

	dbInstance, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("getting error establishing db connection", err.Error())
	} else {
		fmt.Println("connected with db")
	}
	db = dbInstance
	return &Store{
		Config: config,
		Db:     db,
	}
}

func (s *Store) Migrate() {
	db := s.Db

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("getting error getting sql instance", err.Error())
		return
	}

	driver, _ := migrateSql.WithInstance(sqlDB, &migrateSql.Config{})

	cwd, _ := os.Getwd()

	migrationPath := filepath.Join(cwd, "store/migrations")
	fmt.Println(">>>migrationPath ", migrationPath)
	absPath, _ := filepath.Abs(migrationPath)
	fmt.Println("abs path", absPath)
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+absPath,
		"mysql",
		driver,
	)

	if err != nil {
		fmt.Println("not able to get migration path", err.Error())
		return
	}

	err = m.Up()
	if err != nil {
		fmt.Println("getting error applying migrations", err.Error())
		return
	} else {
		fmt.Println("migrations ran successfully!")
	}
}
