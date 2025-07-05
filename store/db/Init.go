package db

import (
	"errors"
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
		fmt.Println("failed to establish connection with db", err.Error())
	} else {
		fmt.Println("hola: connected with db")
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
		fmt.Println("failed to initialise db", err.Error())
		return
	}

	driver, _ := migrateSql.WithInstance(sqlDB, &migrateSql.Config{})
	cwd, _ := os.Getwd()
	migrationPath := filepath.Join(cwd, "store/migrations")
	absPath, _ := filepath.Abs(migrationPath)
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+absPath,
		"mysql",
		driver,
	)

	if err != nil {
		fmt.Println("failed to get migration path", err.Error())
		return
	}

	err = m.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("hola: no new migration found! Good to go")
		} else {
			fmt.Println("failed to apply migrations", err.Error())
		}
	} else {
		fmt.Println("hola: migrations run successfully!")
	}
}
