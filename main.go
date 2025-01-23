package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/gin/binding"
    "github.com/go-playground/validator/v10"
    "net/http"
    "time"
)

type Booking struct {
    CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
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

var bookableDateValidator validator.Func = func(fieldLevel validator.FieldLevel) bool {
    date, ok := fieldLevel.Field().Interface().(time.Time)
    if ok {
        now := time.Now()
        if now.After(date) {
            return false
        }
    }

    return true
}

func main() {
    route := gin.Default()
    validate, ok := binding.Validator.Engine().(*validator.Validate)
    if ok {
        validate.RegisterValidation("bookabledate", bookableDateValidator)
    }
    route.GET("/booking", getBooking)
    route.Run()
}
