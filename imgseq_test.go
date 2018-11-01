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
		t.FailNow()
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
		t.FailNow()
	}
}

//func TestFileSequenceEndsWithMissingFile
func TestFileSequenceEndsWithMissingFile(t *testing.T) {
	createTestDir()
	defer removeTestDir()
	files := []string{"./temp/img10.png",
		"./temp/img11.png",
		"./temp/img12.png",
		"./temp/img13.png"}
	for _, v := range files {
		createImageFile(v, 10, 10)
		defer os.Remove(v)
	}

	seq, err := initImgSeqString("./temp/img10.png")
	if err != nil {
		panic("Failed to create image seq")
	}
	if len(seq.images) != 4 {
		t.FailNow()
	}
}

//func TestFileSequenceEndsWithInvalidFile
func TestFileSequenceEndsWithInvalidFile(t *testing.T) {
	createTestDir()
	defer removeTestDir()
	createImageFile("./temp/img13.png", 11, 10) // Different dimentions
	defer os.Remove("./temp/img13.png")
	os.Open("./temp/img11.png")
	seq, err := initImgSeqString("./temp/img10.png")
	if err != nil {
		panic("Failed to create image seq")
	}
	if len(seq.images) != 3 {
		t.FailNow()
	}
}
