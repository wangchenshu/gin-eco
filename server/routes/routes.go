package routes

import (
	"gin-eco/server/service/words"

	"github.com/gin-gonic/gin"
)

func Engine() *gin.Engine {
	r := gin.Default()

	// words
	r.GET("/api/eco/words", words.GetEcoWords())
	r.GET("/api/eco/words/rand", words.GetEcoWordsRand())

	// wisdom-adages
	r.GET("/api/eco/wisdom-adages", words.GetEcoWisdomAdages())
	r.GET("/api/eco/wisdom-adages/rand", words.GetEcoWisdomAdagesRand())

	// inspirationals
	r.GET("/api/eco/inspirationals", words.GetEcoInspirationals())
	r.GET("/api/eco/inspirationals/rand", words.GetEcoInspirationalsRand())

	// phorisms
	r.GET("/api/eco/phorisms", words.GetEcoPhorisms())
	r.GET("/api/eco/phorisms/rand", words.GetEcoPhorismsRand())

	// Line bot callback
	r.POST("/callback", words.PostHandler())

	return r
}
