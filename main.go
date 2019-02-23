package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/secanis/docker-image-checker/pkg/dic"
)

func main() {
	// load given params
	paramBaseImage := flag.String("base", "", "the base image, format example: hub.docker.com/library/alpine:latest")
	paramImageTest := flag.String("image", "", "the image to test, format example: hub.docker.com/library/alpine:xyz")
	flag.Parse()

	// if we do not have the required parameters, print it and quit
	if len(*paramBaseImage) == 0 || len(*paramImageTest) == 0 {
		log.Fatal("Please provide params, for help execute add the param -h")
	}

	// prepare the paramas to structured values
	paramBaseImageArr := strings.Split(*paramBaseImage, "/")
	paramBaseImageArrTag := strings.Split(paramBaseImageArr[2], ":")
	paramImageTestArr := strings.Split(*paramImageTest, "/")
	paramImageTestArrTag := strings.Split(paramImageTestArr[2], ":")

	// input pattern for i.e.: dic.GetTagObject("hub.docker.com", "library", "alpine", "latest")
	firstImage, err := dic.GetTagObject(paramBaseImageArr[0], paramBaseImageArr[1], paramBaseImageArrTag[0], paramBaseImageArrTag[1])
	secondImage, err := dic.GetTagObject(paramImageTestArr[0], paramImageTestArr[1], paramImageTestArrTag[0], paramImageTestArrTag[1])

	// when we got an error we will exit the tool here and log it
	if err != nil {
		log.Fatal(err)
	}

	// format the time we got from the Docker API to a comparable timestamp
	baseImageLastUpdated, _ := time.Parse(time.RFC3339, firstImage.LastUpdated)
	imageToCheckLastUpdated, _ := time.Parse(time.RFC3339, secondImage.LastUpdated)

	fmt.Printf("%s | updated: %s\n", *paramBaseImage, baseImageLastUpdated)
	fmt.Printf("%s | updated: %s\n", *paramImageTest, imageToCheckLastUpdated)

	// print the result, if the given image is up to date, if not exit with status 1
	result := dic.OutdatedImage(baseImageLastUpdated, imageToCheckLastUpdated)
	if result {
		log.Fatal("Your given docker image is not up to date :(")
	} else {
		fmt.Println("Your given Image is up to date :)")
	}
}
