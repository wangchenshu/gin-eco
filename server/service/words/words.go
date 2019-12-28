package words

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"

	"net/http"

	"gin-eco/server/db"
	"gin-eco/server/model"
	"gin-eco/server/service/mylinebot"
)

var myBot = mylinebot.Init()

func GetWords() gin.HandlerFunc {
	return func(context *gin.Context) {
		events, err := myBot.ParseRequest(context.Request)
		fmt.Println(events)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				context.JSON(http.StatusBadRequest, nil)
			} else {
				context.JSON(500, nil)
			}
			return
		}

		for _, event := range events {
			switch event.Type {
			case linebot.EventTypeMessage:
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					repMsg := "請輸入好語，自在語 或是 靜思語"

					if message.Text == "好語" {
						repMsg = wordsRandWords()
					} else if message.Text == "自在語" {
						repMsg = wisdomAdagesRandWords()
					} else if message.Text == "靜思語" {
						repMsg = phorismsRandWords()
					} else if message.Text == "勵志語" {
						repMsg = inspirationalsRandWords()
					}
					imageURL := "https://firebasestorage.googleapis.com/v0/b/walter-bot-a2142.appspot.com/o/line-bot%2Fimage%2Fother%2Fgirl_img%2F27367-5nYPUB.jpg?alt=media&token=9ec89929-5b2d-478c-b8da-c37f61f338a0"
					template := linebot.NewButtonsTemplate(
						imageURL, "禪念 Bot Go 1.0", repMsg,
						linebot.NewMessageAction("好語", "好語"),
						linebot.NewMessageAction("自在語", "自在語"),
						linebot.NewMessageAction("靜思語", "靜思語"),
						linebot.NewMessageAction("勵志語", "勵志語"),
					)
					if _, err := myBot.ReplyMessage(
						event.ReplyToken,
						linebot.NewTemplateMessage(repMsg, template),
						linebot.NewTextMessage(repMsg),
					).Do(); err != nil {
						log.Print(err)
					}

					// if _, err = myBot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(repMsg)).Do(); err != nil {
					// 	log.Print(err)
					// }
				case *linebot.ImageMessage:
					log.Print(message)
				case *linebot.VideoMessage:
					log.Print(message)
				case *linebot.AudioMessage:
					log.Print(message)
				case *linebot.FileMessage:
					log.Print(message)
				case *linebot.LocationMessage:
					log.Print(message)
				case *linebot.StickerMessage:
					log.Print(message)
				default:
					log.Printf("Unknown message: %v", message)
				}
			default:
				log.Printf("Unknown event: %v", event)
			}
		}
		context.JSON(http.StatusOK, gin.H{
			"success": events,
		})
	}
}

func wordsRandWords() string {
	words := model.EcoWords{}
	db.Db.Raw("SELECT * FROM eco_words where LENGTH(words) < ? ORDER BY RAND() LIMIT 1", 60).Scan(&words)

	return words.Words
}

func wisdomAdagesRandWords() string {
	words := model.EcoWisdomAdages{}
	db.Db.Raw("SELECT * FROM eco_wisdom_adages where LENGTH(words) < ? ORDER BY RAND() LIMIT 1", 60).Scan(&words)

	return words.Words
}

func inspirationalsRandWords() string {
	words := model.EcoInspirationals{}
	db.Db.Raw("SELECT * FROM eco_inspirationals where LENGTH(words) < ? ORDER BY RAND() LIMIT 1", 60).Scan(&words)

	return words.Words
}

func phorismsRandWords() string {
	words := model.EcoPhorisms{}
	db.Db.Raw("SELECT * FROM eco_phorisms where LENGTH(words) < ? ORDER BY RAND() LIMIT 1", 60).Scan(&words)

	return words.Words
}

func GetEcoWords() gin.HandlerFunc {
	return func(c *gin.Context) {
		words := []model.EcoWords{}
		db.Db.Find(&words)

		c.JSON(http.StatusOK, words)
	}
}

func GetEcoWordsRand() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, wordsRandWords())
	}
}

func GetEcoWisdomAdages() gin.HandlerFunc {
	return func(c *gin.Context) {
		words := []model.EcoWisdomAdages{}
		db.Db.Find(&words)

		c.JSON(http.StatusOK, words)
	}
}

func GetEcoWisdomAdagesRand() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, wisdomAdagesRandWords())
	}
}

func GetEcoInspirationals() gin.HandlerFunc {
	return func(c *gin.Context) {
		words := []model.EcoInspirationals{}
		db.Db.Find(&words)

		c.JSON(http.StatusOK, words)
	}
}

func GetEcoInspirationalsRand() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, inspirationalsRandWords())
	}
}

func GetEcoPhorisms() gin.HandlerFunc {
	return func(c *gin.Context) {
		words := []model.EcoPhorisms{}
		db.Db.Find(&words)

		c.JSON(http.StatusOK, words)
	}
}

func GetEcoPhorismsRand() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, phorismsRandWords())
	}
}
