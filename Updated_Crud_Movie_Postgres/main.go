package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MovieTest struct {
	ID        uint   `gorm:"primaryKey"`
	MovieName string `gorm:"<-:create"`
	Actor     string `gorm:"<-:create"`
	Actress   string `gorm:"<-:create"`
}

var db *gorm.DB
var err error

func init() {
	// dsn := `host=localhost
	// 		user=postgres
	// 		password=admin
	// 		dbname=test
	// 		port=5432
	// 		sslmode=disable
	// 		TimeZone=Asia/Shanghai`

	os.Setenv("DB_HOST", "localhost")
	host := os.Getenv("DB_HOST")
	fmt.Println("DB_HOST : ", host)

	os.Setenv("DB_USER", "postgres")
	user := os.Getenv("DB_USER")
	fmt.Println("DB_USER : ", user)

	os.Setenv("DB_PASS", "admin")
	password := os.Getenv("DB_PASS")
	fmt.Println("DB_PASS : ", password)

	os.Setenv("DB_NAME", "test")
	dbname := os.Getenv("DB_NAME")
	fmt.Println("DB_NAME : ", dbname)

	os.Setenv("DB_PORT", "5432")
	port := os.Getenv("DB_PORT")
	fmt.Println("DB_PORT : ", port)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
	os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	fmt.Println("DSN:", dsn)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	fmt.Println(db, err)
	if err != nil {
		return
	}
	db.AutoMigrate(&MovieTest{})
	fmt.Println("Database Created!")
}

// INSERT INTO "movie_tests" ("moviename","actor","actress") VALUES ('Stree','Rajkumar Rao','Shradhha Kapoor')
func main() {

	//createDbRecord(db)
	// fetchWithConditionRecord(db)
	//updateRecord(db)
	// fetchAllRecord(db)
	deleteRecord(db)
}

func createDbRecord(db *gorm.DB) {
	u := MovieTest{MovieName: "Stree", Actor: "Rajkumar Rao", Actress: "Shradhha Kapoor"}
	resp := db.Create(&u)
	fmt.Println(resp.Error, resp.RowsAffected)
}

func fetchWithConditionRecord(db *gorm.DB) {
	var u MovieTest
	id := "1; drop database name;"
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		return
	}

	tc := db.Where("id=?", idInt).Find(&u)
	fmt.Println(tc.RowsAffected)
	fmt.Println(u)
}

func fetchAllRecord(db *gorm.DB) {
	var u []MovieTest
	db.Find(&u)
	fmt.Println(u)
}

func updateRecord(db *gorm.DB) {
	resp := db.Table("movie_tests").Where("id=?", 10).Updates(map[string]interface{}{"movie_name": "Tiger3", "actor": "Salman Khan", "actress": "Katrina Kaif"})
	fmt.Println(resp, resp.Error, resp.RowsAffected)
}

func deleteRecord(db *gorm.DB) {
	resp := db.Where("id=?", 16).Delete(&MovieTest{})
	fmt.Println(resp, resp.Error, resp.RowsAffected)
}

//Feedback
// Hardcoded Credentials:
// Issue: Database credentials are hardcoded in the dsn.
// Fix: For better security and maintainability, store these values in environment variables or a configuration file.
// Example:
// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"),os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

//Consistent Naming Convention:
// Observation: The struct fields use inconsistent naming conventions (Movie_Name, Price_of_movie, Desc_of_movie).
// Fix: Use Go's camelCase naming conventions for struct fields (e.g., MovieName, PriceOfMovie, DescOfMovie) to improve code readability and follow Go best practices.
// Example:
// type Movie struct {
//     ID          uint   `gorm:"primaryKey"`
//     MovieName   string `gorm:"<-:create"`
//     Actor       string `gorm:"<-:create"`
// 	   Actress     string `gorm:"<-:create"`
// }
