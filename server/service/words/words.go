package words

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"

	"net/http"

	"gin-eco/server/db"
	"gin-eco/server/enum"
	"gin-eco/server/model"
	"gin-eco/server/service/mylinebot"
)

const (
	GOOD_WORDS    enum.EcoEnum = enum.GOOD_WORDS
	WISDOM_ADAGE  enum.EcoEnum = enum.WISDOM_ADAGE
	PHORISM       enum.EcoEnum = enum.PHORISM
	INSPIRATIONAL enum.EcoEnum = enum.INSPIRATIONAL
	INPUT_WORDS   enum.EcoEnum = enum.INPUT_WORDS
)

const (
	TITLE       = enum.TITLE
	DEFAULT_IMG = enum.DEFAULT_IMG
	WORDS_LIMIT = enum.WORDS_LIMIT
	QUICK_MENU  = enum.QUICK_MENU
	IMG_URL_1   = enum.IMG_URL_ZEN
	IMG_URL_2   = enum.IMG_URL_ZEN
	IMG_URL_3   = enum.IMG_URL_ZEN
	IMG_URL_4   = enum.IMG_URL_ZEN
)

var myBot = mylinebot.Init()

// PostHandler - get words
func PostHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		events, err := myBot.ParseRequest(context.Request)
		fmt.Println(events)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				context.JSON(http.StatusBadRequest, nil)
			} else {
				context.JSON(http.StatusInternalServerError, nil)
			}
			return
		}

		for _, event := range events {
			switch event.Type {
			case linebot.EventTypeMessage:
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					defaultMsg := fmt.Sprintf(
						"%s: %s, %s, %s, %s",
						INPUT_WORDS.String(),
						GOOD_WORDS.String(),
						WISDOM_ADAGE.String(),
						PHORISM.String(),
						INSPIRATIONAL.String(),
					)
					repMsg := defaultMsg
					if message.Text == GOOD_WORDS.String() {
						repMsg = wordsRandWords()
					} else if message.Text == WISDOM_ADAGE.String() {
						repMsg = wisdomAdagesRandWords()
					} else if message.Text == PHORISM.String() {
						repMsg = phorismsRandWords()
					} else if message.Text == INSPIRATIONAL.String() {
						repMsg = inspirationalsRandWords()
					}

					imageURL := DEFAULT_IMG
					template := linebot.NewButtonsTemplate(
						imageURL, TITLE, defaultMsg,
						linebot.NewMessageAction(GOOD_WORDS.String(), GOOD_WORDS.String()),
						linebot.NewMessageAction(WISDOM_ADAGE.String(), WISDOM_ADAGE.String()),
						linebot.NewMessageAction(PHORISM.String(), PHORISM.String()),
						linebot.NewMessageAction(INSPIRATIONAL.String(), INSPIRATIONAL.String()),
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
	db.Db.Raw("SELECT * FROM eco_words where LENGTH(words) < ? ORDER BY RAND() LIMIT 1", WORDS_LIMIT).Scan(&words)

	return words.Words
}

func wisdomAdagesRandWords() string {
	words := model.EcoWisdomAdages{}
	db.Db.Raw("SELECT * FROM eco_wisdom_adages where LENGTH(words) < ? ORDER BY RAND() LIMIT 1", WORDS_LIMIT).Scan(&words)

	return words.Words
}

func inspirationalsRandWords() string {
	words := model.EcoInspirationals{}
	db.Db.Raw("SELECT * FROM eco_inspirationals where LENGTH(words) < ? ORDER BY RAND() LIMIT 1", WORDS_LIMIT).Scan(&words)

	return words.Words
}

func phorismsRandWords() string {
	words := model.EcoPhorisms{}
	db.Db.Raw("SELECT * FROM eco_phorisms where LENGTH(words) < ? ORDER BY RAND() LIMIT 1", WORDS_LIMIT).Scan(&words)

	return words.Words
}

func myQuickReply() linebot.SendingMessage {
	content := QUICK_MENU
	imageURLs := []string{IMG_URL_1, IMG_URL_2, IMG_URL_3, IMG_URL_4}
	labels := []string{
		GOOD_WORDS.String(),
		WISDOM_ADAGE.String(),
		PHORISM.String(),
		INSPIRATIONAL.String(),
	}
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
