package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/hiteshrepo/grpc-demo/dist"
)

const (
	jsonOutputPath  = "out/data.json"
	protoOutputPath = "out/data.pb"
)

func main() {
	jsonData := dist.JsonLarge
	pbData := dist.ProtocLarge

	// seriallize json
	b, err := json.Marshal(&jsonData)
	check(err)

	err = os.WriteFile(jsonOutputPath, b, 0644)
	check(err)
	fileSize(jsonOutputPath)

	// seriallize proto
	d, err := proto.Marshal(&pbData)
	check(err)

	err = os.WriteFile(protoOutputPath, d, 0644)
	check(err)
	fileSize(protoOutputPath)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fileSize(path string) {
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("unable to fetch file stats", err)
	}
	x := info.Size()
	fmt.Printf("size of file at %s is %d bytes\n", path, x)
}
