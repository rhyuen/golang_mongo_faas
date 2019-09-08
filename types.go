package types

type Payload struct {
	Path string  `json:"path"`
	Data []Quote `json:"quotes"`
}

type Quote struct {
	Author string `json:"author"`
	Text   string `json:"text"`
}
