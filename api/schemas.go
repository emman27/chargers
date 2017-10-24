package api

type updateResponse struct {
	Result []UpdateSchema
	OK     bool
}

// UserSchema according to Telegram
type UserSchema struct {
	ID        int
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// ChatSchema represents the Chat a messsage came from. Used to reply messages
type ChatSchema struct {
	ID    int
	Type  string
	Title string
}

// MessageSchema as per Telegram
type MessageSchema struct {
	From UserSchema
	Text string
	Chat ChatSchema
}

const baseURL = "https://api.telegram.org/bot359390703:AAHbvNwIrh4M97IEvbhZb1ZvBDygNs50I20/"

// UpdateSchema for the update array that will be passed back
type UpdateSchema struct {
	Message           MessageSchema
	EditedMessage     MessageSchema `json:"edited_message"`
	ChannelPost       MessageSchema `json:"channel_post"`
	EditedChannelPost MessageSchema `json:"edited_channel_post"`
	UpdateID          int           `json:"update_id"`
}
