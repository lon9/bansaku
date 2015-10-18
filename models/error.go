package models

// Error is used for response error message.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
