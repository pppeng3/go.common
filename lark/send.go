package lark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func NewText(text string) *Text {
	return &Text{
		MsgType: "text",
		Content: Content{
			T: text,
		},
	}
}

func NewRichText(title string, text string, href string, userId string) *RichText {
	richText := &RichText{
		MsgType: "post",
		Content: RichContent{
			Post: Post{
				ZhCn: ZhCn{
					Title: title,
					Content: [][]Body{
						{},
					},
				},
			},
		},
	}
	if text != "" {
		richText.Content.Post.ZhCn.Content = append(richText.Content.Post.ZhCn.Content, []Body{{Tag: "text", Text: text + "\n"}})
	}
	if href != "" {
		richText.Content.Post.ZhCn.Content = append(richText.Content.Post.ZhCn.Content, []Body{{Tag: "a", Text: "请查看\n", Href: href}})
	}
	if userId != "" {
		richText.Content.Post.ZhCn.Content = append(richText.Content.Post.ZhCn.Content, []Body{{Tag: "at", UserId: userId}})
	}
	return richText
}

func NewBussinessCard(shareChatId string) *BussinessCard {
	return &BussinessCard{
		MsgType: "share_chat",
		Content: Sc{
			ShareChatId: shareChatId,
		},
	}
}

func NewImage(imageKey string) *Image {
	return &Image{
		MsgType: "image",
		Content: Ik{
			ImageKey: imageKey,
		},
	}
}

func NewMessageCard(enableForward bool, wideScreenMode bool) *MessageCard {
	return &MessageCard{
		MsgType: "interactive",
		Card: Card{
			Config: Config{
				EnableForward:  enableForward,
				WideScreenMode: wideScreenMode,
			},
			Elements: []Element{
				{
					Actions: []Action{
						{
							Tag: "button",
							Text: CardText{
								Content: "西湖",
								Tag:     "lark_md",
							},
							Type:  "default",
							URL:   "http://www.baidu.com",
							Value: Value{},
						},
					},
					Tag: "div",
					Text: CardText{
						Content: "东湖",
						Tag:     "lark_md",
					},
				},
			},
			Header: Header{
				Title: CardTitle{
					Content: "今日旅游推荐",
					Tag:     "plain_text",
				},
			},
		},
	}
}

func SendLark(larkWebHook string, data interface{}, contentType string) error {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	jsonStr, err := json.Marshal(data)
	if err != nil {
		// log
		return err
	}
	resp, err := client.Post(larkWebHook, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		// log
		return err
	}
	fmt.Println(bytes.NewBuffer(jsonStr))
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(result))
	return nil
}
