package main

import (
	"fmt"
	"image"
	"log"
	"os"

	"github.com/disintegration/imaging"
)

func main() {

	var dstImage *image.NRGBA

	// Open an image
	srcImage, err := imaging.Open(os.Args[2])
	if err != nil {
		log.Fatalf("ERR: %v", err)
	}

	switch os.Args[1] {
	case "help":
		//To Do
		fmt.Println("TO DO")
	case "blur":
		dstImage = imaging.Blur(srcImage, 3.5)

	case "rotate":
		fmt.Println("Rotating image")
		dstImage = imaging.Rotate90(srcImage)
	}

	//Save
	errSave := imaging.Save(dstImage, "./done.jpg")
	if errSave != nil {
		log.Fatalf("ERR: %v", errSave)
	}
}
