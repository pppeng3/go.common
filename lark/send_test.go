package lark

import (
	"testing"
)

func TestSendText(t *testing.T) { // pass
	tt := NewText("text1")
	SendLark("https://open.larksuite.com/open-apis/bot/v2/hook/dcff66c2-7239-4041-9be4-8907233b8dfd", tt, "application/json")
}

func TestSendRichText(t *testing.T) { // pass
	rt := NewRichText("title", "text", "www.baidu.com", "12312312")
	SendLark("https://open.larksuite.com/open-apis/bot/v2/hook/dcff66c2-7239-4041-9be4-8907233b8dfd", rt, "application/json")
}

func TestSendBussinessCard(t *testing.T) {
	bc := NewBussinessCard("https://open.larksuite.com/open-apis/bot/v2/hook/dcff66c2-7239-4041-9be4-8907233b8dfd")
	SendLark("https://open.larksuite.com/open-apis/bot/v2/hook/dcff66c2-7239-4041-9be4-8907233b8dfd", bc, "application/json")
}

func TestSendImage(t *testing.T) { // pass
	im := NewImage("img_ecffc3b9-8f14-400f-a014-05eca1a4310g")
	SendLark("https://open.larksuite.com/open-apis/bot/v2/hook/dcff66c2-7239-4041-9be4-8907233b8dfd", im, "application/json")
}

func TestSendMessageCard(t *testing.T) { // pass
	mc := NewMessageCard(true, true)
	SendLark("https://open.larksuite.com/open-apis/bot/v2/hook/dcff66c2-7239-4041-9be4-8907233b8dfd", mc, "application/json")
}
