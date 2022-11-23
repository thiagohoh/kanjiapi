package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kyary/kanjiapi/controllers"
)



func SetUpRoute() *gin.Engine{
    return gin.Default()
}

func TestRouteGetN(t *testing.T){

    r := SetUpRoute()
    r.GET("kanjis/:level", controllers.GetByN)

}
