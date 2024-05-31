package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var (
	DB *gorm.DB
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func initMySQL() (err error) {
	dsn := "root:Outer233@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func main() {
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	defer DB.Close()
	DB.AutoMigrate(&Todo{})

	router := gin.Default()
	router.Static("/static", "bubble_frontend/bubble_frontend-master/static")
	router.LoadHTMLGlob("bubble_frontend/bubble_frontend-master/templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1Group := router.Group("v1")
	{
		v1Group.POST("/todo", func(c *gin.Context) {
			var todo Todo
			c.ShouldBindJSON(&todo)
			// 3. 反回响应
			if err = DB.Create(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		v1Group.GET("/todo", func(c *gin.Context) {
			var todoList []Todo
			if err = DB.Find(&todoList).Error; err != nil {
				c.JSON(200, gin.H{"msg": "error"})
			} else {
				c.JSON(200, todoList)
			}
		})
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			id, _ := c.Params.Get("id")
			var todo Todo
			if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
				c.JSON(200, err.Error())
				return
			}
			c.ShouldBindJSON(&todo)
			DB.Save(&todo)
		})
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效id"})
			}
			if err = DB.Where("id=?", id).Delete(Todo{}).Error; err != nil {
				c.JSON(200, gin.H{"error": err.Error()})
			} else {
				c.JSON(200, gin.H{id: "deleted"})
			}
		})
	}
	_ = router.Run(":8080")
}
