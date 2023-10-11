package main

import (
        "fmt"
        "google.golang.org/protobuf/proto"
        pb"github.com/amsem/protobufs/protofiles"
)

func main()  {
    p := &pb.Person{
        Id: 1234,
        Name: "Gol D Roger",
        Email: "luffy@onepiece.com",
        Phones: []*pb.Person_PhoneNumber{
            {Number: "432-423", Type: pb.Person_HOME},
        },
    }
    p1 := &pb.Person{}
    body, _ := proto.Marshal(p)
    _ = proto.Unmarshal(body, p1)
    fmt.Printf("Original struct :%v\n",p)
    fmt.Printf("Marshalled proto %v \n", body)
    fmt.Printf("Unmarshaled struct : %v \n", p1)
}
