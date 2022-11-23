package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kyary/kanjiapi/database"
	"github.com/kyary/kanjiapi/models"
)

func Greetings(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"Status": "The server is on",
	})
}

func GetAllKanjis(c *gin.Context) {

	var kanjis []models.Kanji

	database.DB.Find(&kanjis)

	c.JSON(http.StatusOK, kanjis)
}

func CreateJlptLvl(c *gin.Context){

    var jl models.Jlptlvl

    if err := c.ShouldBindJSON(&jl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
}

func UpdateJlptLvl(c *gin.Context){
    var jl models.Jlptlvl

    id := c.Params.ByName("id")

    database.DB.Find(&jl, id)

    if err := c.ShouldBindJSON(&jl); err != nil{
        c.JSON(http.StatusBadRequest, gin.H{
            "Error": "error" + err.Error(),
        })
        return
    }

    if err := models.JlptValidation(&jl); err != nil{
        c.JSON(http.StatusBadRequest, gin.H{
            "Error": err.Error(),
        })
        return
    }

    database.DB.Save(&jl)
    c.JSON(http.StatusOK, jl)
}
func GetByN(c *gin.Context) {

	n := c.Params.ByName("level")

	var kanjis []models.Jlptlvl

	err := database.DB.Model(&models.Jlptlvl{}).Preload("Kanjis").Where("level = ?", n).Find(&kanjis).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, kanjis)
}

func SearchKanjiByMeaning(c *gin.Context){

    var k []models.Kanji

    m := c.Params.ByName("meaning")

    database.DB.Where("meaning LIKE  ?","%"+m+"%").Find(&k)

    c.JSON(http.StatusOK, &k)
}

func CreateKanji(c *gin.Context) {
	var k models.Kanji

	if err := c.ShouldBindJSON(&k); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Bad Request" +err.Error(),
		})
		return
	}

	if err := models.KanjiValidation(&k); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	database.DB.Create(&k)
	c.JSON(http.StatusOK, k)
}


func UpdateKanji(c *gin.Context){
    var k models.Kanji

    id := c.Params.ByName("id")
    
    database.DB.Find(&k, id)

    if err := c.ShouldBindJSON(&k); err != nil{
        c.JSON(http.StatusBadRequest, gin.H{
            "Error": err.Error(),
        })
        return
    }

    if err := models.KanjiValidation(&k); err != nil{
        c.JSON(http.StatusBadRequest, gin.H{
            "Error": err.Error(),
        })
        return
    }

    database.DB.Save(&k)
    c.JSON(http.StatusOK, k)
}

func DeleteKanji(c *gin.Context){
    var k models.Kanji

    id := c.Params.ByName("id")

    database.DB.Delete(&k,id)

    c.JSON(http.StatusOK, gin.H{
        "Data": "Kanji deleted :" + k.Kanji,
    })
}
