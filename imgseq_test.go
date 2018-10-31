package main

import (
	"image"
	"image/png"
	"os"
	"testing"
)

func createTestDir() {
	err := os.Mkdir("./temp", 0777)
	if err != nil {
		panic(err)
	}
}

func removeTestDir() {
	err := os.Remove("./temp")
	if err != nil {
		panic(err)
	}
}

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
	createTestDir()
	defer removeTestDir()
	createImageFile("./temp/img.png", 10, 10)
	seq, err := initImgSeqString("./temp/img.png")
	if err != nil {
		panic("Failed to create image seq")
	}
	os.Remove("./temp/img.png")
	if len(seq.images) != 1 {
		t.Fail()
	}
}

func TestSingleFileWithNums(t *testing.T) {
	createTestDir()
	defer removeTestDir()
	createImageFile("./temp/img09.png", 10, 10)
	seq, err := initImgSeqString("./temp/img09.png")
	if err != nil {
		panic("Failed to create image seq")
	}
	os.Remove("./temp/img09.png")
	if len(seq.images) != 1 {
		t.Fail()
	}
}

//func TestFileSequenceEndsWithMissingFile
func TestFileSequenceEndsWithMissingFile(t *testing.T) {
	createTestDir()
	defer removeTestDir()
	createImageFile("./temp/img10.png", 10, 10)
	createImageFile("./temp/img11.png", 10, 10)
	createImageFile("./temp/img12.png", 10, 10)
	createImageFile("./temp/img13.png", 10, 10)
	os.Open("./temp/img11.png")

	seq, err := initImgSeqString("./temp/img10.png")
	if err != nil {
		panic("Failed to create image seq")
	}
	os.Remove("./temp/img10.png")
	os.Remove("./temp/img11.png")
	os.Remove("./temp/img12.png")
	os.Remove("./temp/img13.png")
	if len(seq.images) != 4 {
		t.Fail()
	}
}

//func TestFileSequenceEndsWithInvalidFile
func TestFileSequenceEndsWithInvalidFile(t *testing.T) {
	createTestDir()
	defer removeTestDir()
	createImageFile("./temp/img10.png", 10, 10)
	createImageFile("./temp/img11.png", 10, 10)
	createImageFile("./temp/img12.png", 10, 10)
	createImageFile("./temp/img13.png", 11, 10)
	os.Open("./temp/img11.png")

	seq, err := initImgSeqString("./temp/img10.png")
	if err != nil {
		panic("Failed to create image seq")
	}
	os.Remove("./temp/img10.png")
	os.Remove("./temp/img11.png")
	os.Remove("./temp/img12.png")
	os.Remove("./temp/img13.png")
	if len(seq.images) != 3 {
		t.Fail()
	}
}

//func TestFileSequenceEndsWith
