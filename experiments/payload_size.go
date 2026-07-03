package main

import (
	"encoding/json"
	"fmt"
	"strings"

	pb "github.com/shreyasrajiv327/APIBench/proto"
	"google.golang.org/protobuf/proto"
)

type RESTMessage struct {
	ID        string `json:"id"`
	Payload   string `json:"payload"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdAt"`
}

func main() {

	payload := strings.Repeat("A", 100)

	// ---------------- REST ----------------

	rest := RESTMessage{
		ID:        "msg-1",
		Payload:   payload,
		Status:    "QUEUED",
		CreatedAt: "2026-07-01T10:00:00Z",
	}

	restBytes, _ := json.Marshal(rest)

	// ---------------- GraphQL ----------------

	graphqlResponse := map[string]interface{}{
		"data": map[string]interface{}{
			"publish": rest,
		},
	}

	graphqlBytes, _ := json.Marshal(graphqlResponse)

	// ---------------- gRPC ----------------

grpcMessage := &pb.Message{
    Id:        "msg-1",
    Payload:   []byte(payload),
    Status:    "QUEUED",
    CreatedAt: "2026-07-01T10:00:00Z",
}

	protoBytes, _ := proto.Marshal(grpcMessage)

	fmt.Println("========== Payload Size Comparison ==========")
	fmt.Printf("REST JSON      : %d bytes\n", len(restBytes))
	fmt.Printf("GraphQL JSON   : %d bytes\n", len(graphqlBytes))
	fmt.Printf("ProtocolBuffer : %d bytes\n", len(protoBytes))
}