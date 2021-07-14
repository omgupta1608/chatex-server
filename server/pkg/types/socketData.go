package types

type (
	Message struct {
		M_ID         string `json:"mid"` // Id of new message | Id of message to delete | "" in case of typing and user_conn
		S_ID         string `json:"sid"`
		R_ID         string `json:"rid"`
		Content      string `json:"content"` // "" is case of typing, user_conn and delete msg
		Timestamp    string `json:"ts"`
		Message_Type string `json:"m_type"` // "text", "img", "video", "file", "doc" etc. | "event" is case of typing, user_conn and delete msg
	}

	EventFormat struct {
		Event_Name string  `json:"name"` // NEW_MESSAGE | DELETE_MESSAGE | TYPING | USER_CONN
		Data       Message `json:"data"`
	}
)
