package types

import "google.golang.org/protobuf/types/known/timestamppb"

type (
	Message struct {
		S_ID         string                 `json:"sid"`
		R_ID         string                 `json:"rid"`
		Content      string                 `json:"content"`
		Timestamp    *timestamppb.Timestamp `json:"ts"`
		Message_Type string                 `json:"m_type"` // "text", "img", "video", "file", "doc" etc.
	}

	EventFormat struct {
		Event_Name string
		Data       Message
	}
)
