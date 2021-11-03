package lark

// 文本消息
type Text struct {
	MsgType string  `json:"msg_type"`
	Content Content `json:"content"`
}

type Content struct {
	T string `json:"text"`
}

// 富文本消息
type RichText struct {
	MsgType string      `json:"msg_type"`
	Content RichContent `json:"content"`
}

type RichContent struct {
	Post Post `json:"post"`
}

type Post struct {
	ZhCn ZhCn `json:"zh_cn"`
}

type ZhCn struct {
	Title   string   `json:"title"`
	Content [][]Body `json:"content"`
}

type Body struct {
	Tag    string `json:"tag"`
	Text   string `json:"text"`
	Href   string `json:"href"`
	UserId string `json:"user_id"`
}

// 群名片
type BussinessCard struct {
	MsgType string `json:"msg_type"`
	Content Sc     `json:"content"`
}

type Sc struct {
	ShareChatId string `json:"share_chat_id"`
}

// 图片
type Image struct {
	MsgType string `json:"msg_type"`
	Content Ik     `json:"content"`
}

type Ik struct {
	ImageKey string `json:"image_key"`
}

// 消息卡片
type MessageCard struct {
	MsgType string `json:"msg_type"`
	Card    Card   `json:"card"`
}

type Card struct {
	Config   Config    `json:"config"`
	Elements []Element `json:"elements"`
	Header   Header    `json:"header"`
}

type Config struct {
	EnableForward  bool `json:"enable_forward"`
	WideScreenMode bool `json:"wide_screen_mode"`
}

type Element struct {
	Actions []Action `json:"actions"`
	Tag     string   `json:"tag"`
	Text    CardText `json:"text"`
}

type Header struct {
	Title CardTitle `json:"title"`
}

type CardTitle struct {
	Content string `json:"content"`
	Tag     string `json:"tag"`
}

type Action struct {
	Tag   string   `json:"tag"`
	Text  CardText `json:"text"`
	Type  string   `json:"type"`
	URL   string   `json:"url"`
	Value Value    `json:"value"`
}

type CardText struct {
	Content string `json:"content"`
	Tag     string `json:"tag"`
}

type Value struct {
}
