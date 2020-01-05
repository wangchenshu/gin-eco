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

const (
	// WordStr -
	WordStr = "好語"
	// WisdomAdageStr -
	WisdomAdageStr = "自在語"
	// PhorismStr -
	PhorismStr = "靜思語"
	// InspirationalStr -
	InspirationalStr = "勵志語"
)

// TemplateTitle - template title
const TemplateTitle = "禪念 Bot Go 2.0"

// DefaultImg - default img
const DefaultImg = "https://firebasestorage.googleapis.com/v0/b/walter-bot-a2142.appspot.com/o/line-bot%2Fimage%2Fother%2Fgirl_img%2F27367-5nYPUB.jpg?alt=media&token=9ec89929-5b2d-478c-b8da-c37f61f338a0"

// WordsLimit -
const WordsLimit = 200

var myBot = mylinebot.Init()

// GetWords - get words
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
					defaultMsg := fmt.Sprintf("請輸入%s，%s 或是 %s", WordStr, WisdomAdageStr, PhorismStr)
					repMsg := defaultMsg
					if message.Text == WordStr {
						repMsg = wordsRandWords()
					} else if message.Text == WisdomAdageStr {
						repMsg = wisdomAdagesRandWords()
					} else if message.Text == PhorismStr {
						repMsg = phorismsRandWords()
					} else if message.Text == InspirationalStr {
						repMsg = inspirationalsRandWords()
					}

					imageURL := DefaultImg
					template := linebot.NewButtonsTemplate(
						imageURL, TemplateTitle, defaultMsg,
						linebot.NewMessageAction(WordStr, WordStr),
						linebot.NewMessageAction(WisdomAdageStr, WisdomAdageStr),
						linebot.NewMessageAction(PhorismStr, PhorismStr),
						linebot.NewMessageAction(InspirationalStr, InspirationalStr),
					)
					if _, err := myBot.ReplyMessage(
						event.ReplyToken,
						linebot.NewTemplateMessage(repMsg, template),
						linebot.NewTextMessage(repMsg),
						myQuickReply(),
					).Do(); err != nil {
						log.Print(err)
					}
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
	db.Db.Raw("SELECT * FROM eco_words where LENGTH(words) < ? ORDER BY RAND() LIMIT 1", WordsLimit).Scan(&words)

	return words.Words
}

func wisdomAdagesRandWords() string {
	words := model.EcoWisdomAdages{}
	db.Db.Raw("SELECT * FROM eco_wisdom_adages where LENGTH(words) < ? ORDER BY RAND() LIMIT 1", WordsLimit).Scan(&words)

	return words.Words
}

func inspirationalsRandWords() string {
	words := model.EcoInspirationals{}
	db.Db.Raw("SELECT * FROM eco_inspirationals where LENGTH(words) < ? ORDER BY RAND() LIMIT 1", WordsLimit).Scan(&words)

	return words.Words
}

func phorismsRandWords() string {
	words := model.EcoPhorisms{}
	db.Db.Raw("SELECT * FROM eco_phorisms where LENGTH(words) < ? ORDER BY RAND() LIMIT 1", WordsLimit).Scan(&words)

	return words.Words
}

func myQuickReply() linebot.SendingMessage {
	content := fmt.Sprintf("快速選單")
	imageURLs := []string{
		"https://firebasestorage.googleapis.com/v0/b/walter-bot-a2142.appspot.com/o/line-bot%2Fimage%2Fother%2Fzen%2F3.png?alt=media&token=44770a07-a661-40dc-960e-45da1699e4f2",
		"https://firebasestorage.googleapis.com/v0/b/walter-bot-a2142.appspot.com/o/line-bot%2Fimage%2Fother%2Fzen%2F3.png?alt=media&token=44770a07-a661-40dc-960e-45da1699e4f2",
		"https://firebasestorage.googleapis.com/v0/b/walter-bot-a2142.appspot.com/o/line-bot%2Fimage%2Fother%2Fzen%2F3.png?alt=media&token=44770a07-a661-40dc-960e-45da1699e4f2",
		"https://firebasestorage.googleapis.com/v0/b/walter-bot-a2142.appspot.com/o/line-bot%2Fimage%2Fother%2Fzen%2F3.png?alt=media&token=44770a07-a661-40dc-960e-45da1699e4f2",
	}
	labels := []string{WordStr, WisdomAdageStr, PhorismStr, InspirationalStr}
	quickReplyButtons := []*linebot.QuickReplyButton{}

	for k, v := range labels {
		quickReplyButtons = append(quickReplyButtons, linebot.NewQuickReplyButton(
			imageURLs[k], linebot.NewMessageAction(v, v),
		))
	}
	quickReply := linebot.NewTextMessage(content).
		WithQuickReplies(linebot.NewQuickReplyItems(quickReplyButtons...))

	return quickReply
}

// GetEcoWords -
func GetEcoWords() gin.HandlerFunc {
	return func(c *gin.Context) {
		words := []model.EcoWords{}
		db.Db.Find(&words)

		c.JSON(http.StatusOK, words)
	}
}

// GetEcoWordsRand -
func GetEcoWordsRand() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, wordsRandWords())
	}
}

// GetEcoWisdomAdages -
func GetEcoWisdomAdages() gin.HandlerFunc {
	return func(c *gin.Context) {
		words := []model.EcoWisdomAdages{}
		db.Db.Find(&words)

		c.JSON(http.StatusOK, words)
	}
}

// GetEcoWisdomAdagesRand -
func GetEcoWisdomAdagesRand() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, wisdomAdagesRandWords())
	}
}

// GetEcoInspirationals -
func GetEcoInspirationals() gin.HandlerFunc {
	return func(c *gin.Context) {
		words := []model.EcoInspirationals{}
		db.Db.Find(&words)

		c.JSON(http.StatusOK, words)
	}
}

// GetEcoInspirationalsRand -
func GetEcoInspirationalsRand() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, inspirationalsRandWords())
	}
}

// GetEcoPhorisms -
func GetEcoPhorisms() gin.HandlerFunc {
	return func(c *gin.Context) {
		words := []model.EcoPhorisms{}
		db.Db.Find(&words)

		c.JSON(http.StatusOK, words)
	}
}

// GetEcoPhorismsRand -
func GetEcoPhorismsRand() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, phorismsRandWords())
	}
}
