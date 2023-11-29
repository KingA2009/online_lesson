package main

import (
	"fmt"
	"github.com/KingA2009/online_lesson/lesson_9"
	"github.com/golang/protobuf/proto"
	"log"
)

func main() {
	fmt.Println("Hello world")

	elliot := &lesson_9.Person{
		Name: "Elliot",
		Age:  24,
		SocialFollowers: &lesson_9.SocialFollowers{
			Youtube: 2500,
			Twitter: 1400,
		},
	}

	data, err := proto.Marshal(elliot)
	if err != nil {
		log.Fatal("Marshaling error: ", err)
	}
	fmt.Println(data)

	newElliot := &lesson_9.Person{}
	err = proto.Unmarshal(data, newElliot)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	fmt.Println(newElliot.GetAge())
	fmt.Println(newElliot.GetName())
	fmt.Println(newElliot.SocialFollowers.GetYoutube())
	fmt.Println(newElliot.SocialFollowers.GetTwitter())
}
