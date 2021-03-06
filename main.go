package main

// only need mysql OR sqlite
// both are included here for reference
import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

var db *gorm.DB
var err error

type Person struct {
	ID        uint   `json:"id"  form:"id"`
	FirstName string `json:"firstname"  form:"firstname"`
	LastName  string `json:"lastname"  form:"lastname"`
	City      string `json:"city"  form:"city"`
}

func main() {

	// NOTE: See we're using = to assign the global var
	//         	instead of := which would assign it only in this function
	// db, err = gorm.Open("sqlite3", "./gorm.db")
	db, _ = gorm.Open("mysql", "root:123456789@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.AutoMigrate(&Person{})
	e := echo.New()
	// r := gin.Default()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
	// r.GET("/people/", GetPeople)
	// r.GET("/people/:id", GetPerson)
	// r.POST("/people", CreatePerson)
	// r.PUT("/people/:id", UpdatePerson)
	// r.DELETE("/people/:id", DeletePerson)

	// r.Run(":8080")
}

func DeletePerson(c *gin.Context) {
	id := c.Params.ByName("id")
	var person Person
	d := db.Where("id = ?", id).Delete(&person)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

func UpdatePerson(c *gin.Context) {

	var person Person
	id := c.Params.ByName("id")

	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&person)

	db.Save(&person)
	c.JSON(200, person)

}

func CreatePerson(c *gin.Context) {
	var person Person
	c.Bind(&person)
	db.Create(&person)
	c.JSON(200, person)
}

func GetPerson(c *gin.Context) {
	id := c.Params.ByName("id")
	var person Person
	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, person)
	}
}
func GetPeople(c *gin.Context) {
	var people []Person
	if err := db.Find(&people).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, people)
	}

}
