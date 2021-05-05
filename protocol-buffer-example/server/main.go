package main

import (
	"fmt"
	grpc "github.com/Hello-gRPC/protocol-buffer-example"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {

	elliot := &grpc.Person{
		Name: "Elliot",
		Age:  24,
	}

	data, err := proto.Marshal(elliot)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	// printing out our raw protobuf object
	fmt.Println(data)

	// let's go the other way and unmarshal
	// our byte array into an object we can modify
	// and use
	newElliot := &grpc.Person{}
	err = proto.Unmarshal(data, newElliot)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	// print out our `newElliot` object
	// for good measure
	fmt.Println(newElliot.GetAge())
	fmt.Println(newElliot.GetName())

}
