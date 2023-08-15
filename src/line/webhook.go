package line

import (
	"fmt"
	"log"
	"mahasanbkk-webscraper/pkg/config"
	"net/http"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func Webhook(writer http.ResponseWriter, req *http.Request) {
    // url := linebot.APIEndpointPushMessage
    lineSession := config.LineSession
    events, err := lineSession.ParseRequest(req)
		if err != nil {
			writer.WriteHeader(500)
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if _, err = lineSession.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						log.Print(err)
					}
				case *linebot.StickerMessage:
					replyMessage := fmt.Sprintf(
						"sticker id is %s, stickerResourceType is %s", message.StickerID, message.StickerResourceType)
					if _, err = lineSession.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
}
