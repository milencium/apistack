package database
import (
	"fmt"
	
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// NewDatabase - returns a pointer to a database object
func NewDatabase()(*gorm.DB, error){
	fmt.Println("Setting up new database connection")
	
	dbUsername := "postgres"
	dbPassword := "postgres"
	dbHost := "localhost"
	dbTable := "postgres"
	dbPort := "5432"

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbTable, dbPassword) 
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		return db, err
	}
	if err := db.DB().Ping(); err != nil {
		return db, err
	}
	return db, nil
}