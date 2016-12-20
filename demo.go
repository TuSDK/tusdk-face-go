package main

import (
	"fmt"
	"log"
	"github.com/TuSDK/tusdk-face-go/api"
)

func main() {
	keys := api.Keys{
		PID: "", // 公有key
		KEY: "", // 私有key
	}

	face := &api.FaceApi{Keys: keys}

	image := map[string]string{"file":"path_to_file"}
	params := map[string]string{}
	data, err := face.Request("detection", image, params)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(data)
	}

}
