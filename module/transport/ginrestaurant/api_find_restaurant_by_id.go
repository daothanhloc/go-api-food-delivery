package restaurantgin

import (
	restaurantbusiness "food-delivery/module/restaurant/business"
	restaurantstorage "food-delivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetRestaurantByIdHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(err)
		}

		storage := restaurantstorage.NewSQLStore(db)
		biz := restaurantbusiness.NewGetRestaurantByIdBusiness(storage)

		restaurant, err := biz.GetRestaurantById(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": restaurant})
	}
}
