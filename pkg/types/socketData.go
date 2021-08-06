package types

type (
	Message struct {
		CLIENT_CHAT_ID  string      `json:"ccid"` // Id of new message | Id of message to delete | "" in case of typing and user_conn
		SERVER_CHAT_ID  string      `json:"scid"` // Id of new message | Id of message to delete | "" in case of typing and user_conn
		S_ID            string      `json:"sid"`
		R_ID            string      `json:"rid"`
		Content         string      `json:"content"` // "" is case of typing, user_conn and delete msg
		ServerTimestamp int64       `json:"sts"`
		Data            interface{} `json:"data"`
		Message_Type    string      `json:"m_type"` // "text", "img", "video", "file", "doc" etc. | "event" is case of typing, user_conn and delete msg
	}

	EventFormat struct {
		Event_Name string  `json:"name"` // NEW_MESSAGE | DELETE_MESSAGE | TYPING | USER_CONN
		Data       Message `json:"data"`
	}
)
