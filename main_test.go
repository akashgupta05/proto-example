package main

import (
	"proto-example/messages"
	"testing"
)

func Benchmark_JSONExample(b *testing.B) {
	person := getPerson()
	for i := 0; i < b.N; i++ {
		JSONExample(&person)
	}
}

func Benchmark_ProtoExample(b *testing.B) {
	person := getPerson()
	for i := 0; i < b.N; i++ {
		ProtoExample(&person)
	}
}

func getPerson() messages.Person {
	return messages.Person{
		Name:  "Akash",
		Id:    "random-id",
		Email: "akashg@getsimpl.com",
		Phones: []*messages.Person_PhoneNumber{
			{
				Number: "1234567890",
				Type:   messages.Person_MOBILE,
			},
		},
	}
}
