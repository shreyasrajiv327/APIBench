package models

import "time"

type MessageStatus string

const(
	StatusQueue      MessageStatus = "QUEUED"
	StatusProcessing MessageStatus = "PROCESSING"
	StatusAcked      MessageStatus = "ACKED"
	StautsNacked     MessageStatus = "NACKED"
)

type Message struct{
	ID       string
	Payload  []byte
	Status   MessageStatus
	Created  time.Time
}