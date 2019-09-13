package types

type Payload struct {
	Path string  `json:"path"`
	Data []Quote `json:"quotes"`
}

type ErrorPayload struct {
	Path        string `json:"path"`
	Description string `json:"error"`
	Message     string `json:"details"`
}

type Quote struct {
	Author string `json:"author"`
	Text   string `json:"text"`
}
