package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "time"
)

type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

func getBooking(context *gin.Context) {
	var booking Booking
	if err := context.ShouldBind(&booking); err == nil {
		context.JSON(
            http.StatusOK,
            gin.H {
                "message": "Booking dates are valid!",
            },
        )
	} else {
		context.JSON(
            http.StatusBadRequest,
            gin.H {
                "error": err.Error(),
            },
        )
	}
}

func main() {
	route := gin.Default()

	route.GET("/booking", getBooking)
	route.Run()
}
