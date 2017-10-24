package api

type updateResponse struct {
	Result []UpdateSchema
	OK     bool
}

// UserSchema according to Telegram
type UserSchema struct {
	ID int
}

// MessageSchema as per Telegram
type MessageSchema struct {
	From UserSchema
	Text string
}

// UpdateSchema for the update array that will be passed back
type UpdateSchema struct {
	Message  MessageSchema
	UpdateID int `json:"update_id"`
}
