package mydb

import (
	"flag"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Employee struct {
	gorm.Model
	Name         string
	ManagerID    uint
	Manager      string
	DepartmentID uint
}

type Department struct {
	gorm.Model
	Name      string
	Employees []Employee
}

func connectDB() {

	dbName := flag.String("dbname", "postgres", "a string")

	flag.Parse()
	fmt.Println(dbName)
	db, err := gorm.Open("postgres", "user=postgres password=root dbname=postgres sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	db.DropTableIfExists(&Employee{})
	db.CreateTable(&Employee{})
	db.DropTableIfExists(&Department{})
	db.CreateTable(&Department{})

	department := Department{
		Name: "Maths",
		Employees: []Employee{
			{Name: "gst", DepartmentID: 2},
		},
	}
	db.Save(&department)

}
