package main

//go:generate protoc --go_out=. --go_opt=paths=source_relative messages/example.proto

import (
	"encoding/json"
	"fmt"
	"proto-example/messages"

	"google.golang.org/protobuf/proto"
)

func main() {
	person := messages.Person{
		Name:  "Akash",
		Id:    "random-id",
		Email: "akashg@getsimpl.com",
		Phones: []*messages.PhoneNumber{
			{
				Number: "1234567890",
				Type:   messages.PhoneType_MOBILE,
			},
		},
	}

	JSONExample(&person)
	ProtoExample(&person)
}

func JSONExample(person interface{}) {
	messageBytes, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("json message length: %d, message: %s\n", len(messageBytes), string(messageBytes))
	newPerson := messages.Person{}
	err = json.Unmarshal(messageBytes, &newPerson)
	if err != nil {
		fmt.Println(err)
	}
}

func ProtoExample(person proto.Message) {
	messageBytes, err := proto.Marshal(person)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("proto message length: %d, message: %s\n", len(messageBytes), string(messageBytes))
	newPerson := messages.Person{}
	err = proto.Unmarshal(messageBytes, &newPerson)
	if err != nil {
		fmt.Println(err)
	}
}
