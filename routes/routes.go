package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kyary/kanjiapi/controllers"
)

func HandleRequests() {

	r := gin.Default()

	r.GET("/", controllers.Greetings)
	r.GET("/kanjis", controllers.GetAllKanjis)
    r.POST("/kanji", controllers.CreateKanji)
    r.DELETE("/kanji/:id", controllers.DeleteKanji)
    r.PUT("/kanji/:id", controllers.UpdateKanji)
    r.GET("/kanji-search/:meaning", controllers.SearchKanjiByMeaning)

	r.GET("/n/:level", controllers.GetByN)

    r.POST("/jlpt", controllers.CreateJlptLvl)
    r.PUT("/jlpt/:id", controllers.UpdateJlptLvl)
	r.Run()
}
