package main

import (
	"flag"
	"fmt"
	"image"
	"os"

	"github.com/disintegration/imaging"
	gosseract "github.com/otiai10/gosseract/v2"
)

// ocr --lang=ara --img=xyz.png
func main() {
	lang := flag.String("lang", "eng", " Language of the written text. eng or ara as a language.")
	img := flag.String("img", "", " Image.")
	flag.Parse()

	if *img == "" {
		fmt.Println("Usage : ocr --lang=eng --img=~/Downloads/xyz.png")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *lang != "eng" && *lang != "ara" {
		fmt.Println("--lang must be ara for Arabic or eng for English.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Println(*lang, *img)
	extracted, err := ocr(*img, *lang, false)
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}

	if extracted == "" {
		fmt.Println("no text extracted! something went wrong")
	}

	fmt.Println(extracted)
}

func ocr(imgpath, lang string, isBlackBg bool) (string, error) {
	client := gosseract.NewClient()
	defer client.Close()

	// client.Languages = []string{"eng", "ara"}
	client.SetLanguage(lang)

	if isBlackBg == true {
		imgIo, _ := os.Open(imgpath)
		imgDec, _, _ := image.Decode(imgIo)
		inverted := imaging.Invert(imgDec)
		imaging.Save(inverted, "./temp.jpg")
		defer os.Remove("./temp.jpg")
		client.SetImage("./temp.jpg")
	} else {
		client.SetImage(imgpath)
	}

	//boundingBox, _ := client.GetBoundingBoxes(PageIteratorLevel.RIL_SYMBOL)
	// boundingBox, err := client.GetBoundingBoxes(gosseract.RIL_SYMBOL)
	// if err != nil {
	// 	return "", err
	// }

	text, err := client.Text()
	// text, err := client.HOCRText()
	if err != nil {
		return "", err
	}
	return text, nil
}
