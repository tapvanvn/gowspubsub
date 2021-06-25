package entity

type Message struct {
	Topic      string            `json:"topic"`
	Attributes map[string]string `json:"attributes,omitempty"`
	Message    string            `json:"message,omitempty"`
	Tier       int               `json:"-"`
}
