package models

type Message struct {
	Type string
	Data []byte
}

type MessageLst struct {
	Msgs         []Message
	ShiftCounter int
}
