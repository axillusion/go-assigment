package database

import (
	"os"
	"path/filepath"

	"github.com/axillusion/go-assigment/commons"
	"github.com/axillusion/go-assigment/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func loadEnv() {
	envPath := ".env"
	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		envPath = filepath.Join("..", ".env")
	}
	err := godotenv.Load(envPath)
	if err != nil {
		commons.Log.Fatal("Could not load env file from path %s with error: %v", envPath, err)
	} else {
		commons.Log.Info("Loaded env file from path %s", envPath)
	}
}

type Database interface {
	Connect() error
	GetDB() Database
	Create(value interface{}) error
	Table(name string) Database
	Select(query interface{}, args ...interface{}) Database
	Where(query interface{}, args ...interface{}) Database
	Order(value interface{}) Database
	Modify(value interface{}) error
	Update(column string, value interface{}) error
	Delete(value interface{}, conds ...interface{}) error
	Limit(limit int) Database
	Offset(offset int) Database
	Find(dest interface{}, conds ...interface{}) error
	Close() error
}

// GormDatabase is a concrete implementation of the Database interface using GORM
type GormDatabase struct {
	db *gorm.DB
}

// Limit specifies the maximum number of records to retrieve
func (g *GormDatabase) Limit(limit int) *gorm.DB {
	return g.db.Limit(limit)
}

// Offset specifies the number of records to skip
func (g *GormDatabase) Offset(offset int) *gorm.DB {
	return g.db.Offset(offset)
}

// Create inserts a new record into the database
func (g *GormDatabase) Create(value interface{}) error {
	return g.db.Create(value).Error
}

// Table specifies the table to perform operations on
func (g *GormDatabase) Table(name string) *gorm.DB {
	return g.db.Table(name)
}

// Select specifies the fields to retrieve
func (g *GormDatabase) Select(query interface{}, args ...interface{}) *gorm.DB {
	return g.db.Select(query, args...)
}

// Where adds conditions to the query
func (g *GormDatabase) Where(query interface{}, args ...interface{}) *gorm.DB {
	return g.db.Where(query, args...)
}

// Order specifies the order of the results
func (g *GormDatabase) Order(value interface{}) *gorm.DB {
	return g.db.Order(value)
}

func (g *GormDatabase) Update(column string, value interface{}) error {
	return g.db.Model(&models.DialogRow{}).Update(column, value).Error
}
func (g *GormDatabase) Modify(value interface{}) error {
	return g.db.Save(value).Error
}

func (g *GormDatabase) Delete(value interface{}, conds ...interface{}) error {
	return g.db.Delete(value, conds...).Error
}

func (g *GormDatabase) Find(dest interface{}, conds ...interface{}) error {
	return g.db.Find(dest, conds...).Error
}

// Connect establishes a connection to the database
func (g *GormDatabase) Connect() error {
	var err error

	loadEnv()

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbString := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	g.db, err = gorm.Open(mysql.Open(dbString), &gorm.Config{})
	if err != nil {
		return err
	}

	err = g.db.AutoMigrate(&models.DialogRow{})
	if err != nil {
		return err
	}

	commons.Log.Info("Successfully established database connection")
	return nil
}

// Close closes the database connection
func (g *GormDatabase) Close() error {
	sqlDB, err := g.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// GetDB returns the underlying *gorm.DB instance
func (g *GormDatabase) GetDB() *gorm.DB {
	return g.db
}
