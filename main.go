package main

import (
	"food-delivery/module/transport/ginrestaurant"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Restaurant struct {
	Id      int    `json:"id" gorm:"column:id"`
	Name    string `json:"name" gorm:"column:name"`
	Address string `json:"address" gorm:"column:address"`
}

func main() {
	dsn := "food-delivery:12345@tcp(localhost:3306)/food-delivery?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := os.Getenv("MYSQL_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	log.Println(db)

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		restaurants := v1.Group("/restaurants")
		{
			restaurants.POST("", restaurantgin.CreateRestaurantHandler(db))
			//restaurants.GET("/:restaurant-id", restaurantgin.GetRestaurantHandler(db))
			//restaurants.GET("", restaurantgin.ListRestaurant(db))
			//restaurants.PUT("/:restaurant-id", restaurantgin.UpdateRestaurantHandler(db))
			restaurants.DELETE("/:id", restaurantgin.DeleteRestaurantByIdHandler(db))
		}
	}

	router.Run(":3003") //default 8080

}
