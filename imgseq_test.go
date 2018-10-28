package main

import (
	"image"
	"image/png"
	"os"
	"testing"
)

func createImageFile(name string, width int, height int) {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}

	err = png.Encode(f, img)
	if err != nil {
		f.Close()
		panic(err)
	}

	if err := f.Close(); err != nil {
		panic(err)
	}

}

func TestSingleFileNoNums(t *testing.T) {
	createImageFile("./files/img.png", 10, 10)
	seq, err := FromString("./files/img.png")
	if err != nil {
		panic("Failed to create image seq")
	}
	os.Remove("./files/img.png")
	if len(seq.images) != 1 {
		t.Fail()
	}
}

func TestSingleFileWithNums(t *testing.T) {
	createImageFile("./files/img09.png", 10, 10)
	seq, err := FromString("./files/img09.png")
	if err != nil {
		panic("Failed to create image seq")
	}
	os.Remove("./files/img09.png")
	if len(seq.images) != 1 {
		t.Fail()
	}
}

//func TestFileSequenceEndsWithMissingFile
func TestFileSequenceEndsWithMissingFile(t *testing.T) {
	createImageFile("./files/img10.png", 10, 10)
	createImageFile("./files/img11.png", 10, 10)
	createImageFile("./files/img12.png", 10, 10)
	createImageFile("./files/img13.png", 10, 10)
	os.Open("./files/img11.png")

	seq, err := FromString("./files/img10.png")
	if err != nil {
		panic("Failed to create image seq")
	}
	os.Remove("./files/img10.png")
	os.Remove("./files/img11.png")
	os.Remove("./files/img12.png")
	os.Remove("./files/img13.png")
	if len(seq.images) != 4 {
		t.Fail()
	}
}

//func TestFileSequenceEndsWithInvalidFile
func TestFileSequenceEndsWithInvalidFile(t *testing.T) {
	createImageFile("./files/img10.png", 10, 10)
	createImageFile("./files/img11.png", 10, 10)
	createImageFile("./files/img12.png", 10, 10)
	createImageFile("./files/img13.png", 11, 10)
	os.Open("./files/img11.png")

	seq, err := FromString("./files/img10.png")
	if err != nil {
		panic("Failed to create image seq")
	}
	os.Remove("./files/img10.png")
	os.Remove("./files/img11.png")
	os.Remove("./files/img12.png")
	os.Remove("./files/img13.png")
	if len(seq.images) != 3 {
		t.Fail()
	}
}

//func TestFileSequenceEndsWith
