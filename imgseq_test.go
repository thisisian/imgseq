package main

import (
	"image"
	"image/png"
	"os"
	"testing"
)

// createPNGSequence creates an image sequence in dir of size length
// for testing purposes
// Images are in form `<dir>/img<index>.png`
func createPNGSequence(len int, dir string) {

}

func createTestDir() {
	err := os.Mkdir("./temp", 0777)
	if err != nil {
		panic(err)
	}
}

func removeTestDir() {
	err := os.RemoveAll("./temp")
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
	seq, err := initImgSeq("./temp/img.png")
	if err != nil {
		t.Fatal(err.Error())
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
	seq, err := initImgSeq("./temp/img09.png")
	if err != nil {
		t.Fatal(err.Error())
	}
	os.Remove("./temp/img09.png")
	if len(seq.images) != 1 {
		t.FailNow()
	}
}

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

	seq, err := initImgSeq("./temp/img10.png")
	if err != nil {
		t.Fatal(err.Error())
	}
	if len(seq.images) != 4 {
		t.Fatalf("len of sequence = %d, expected 4", len(seq.images))
	}
}

func TestFileSequenceEndsWithInvalidFile(t *testing.T) {
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
	createImageFile("./temp/img14.png", 11, 10) // Different dimentions
	defer os.Remove("./temp/img14.png")

	seq, err := initImgSeq("./temp/img10.png")
	if err != nil {
		t.Fatal(err.Error())
	}
	if len(seq.images) != 4 {
		t.FailNow()
	}
}
