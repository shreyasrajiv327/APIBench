package models

import "time"

type MessageStatus string

const(
	StatusQueued      MessageStatus = "QUEUED"
	StatusProcessing MessageStatus = "PROCESSING"
	StatusAcked      MessageStatus = "ACKED"
	StatusNacked     MessageStatus = "NACKED"
)

type Message struct{
	ID       string
	Payload  []byte
	Status   MessageStatus
	Created  time.Time
}