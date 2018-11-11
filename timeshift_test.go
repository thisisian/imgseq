package main

import (
	"os"
	"testing"
)

func TestTsOperationParsing(t *testing.T) {
	createTestDir()
	defer removeTestDir()
	createImageFile("./temp/img01.png", 10, 10)
	defer os.Remove("./temp/img01.png")
	createImageFile("./temp/img02.png", 10, 10)
	defer os.Remove("./temp/img02.png")
	createImageFile("./temp/img03.png", 10, 10)
	defer os.Remove("./temp/img03.png")
	createImageFile("./temp/img04.png", 10, 10)
	defer os.Remove("./temp/img04.png")

	ts, err := initTimeshift("filterimg=./temp/img01.png:range=4")
	if err != nil {
		t.Fatal(err)
	}
	if ts.filterRange != 4 {
		t.FailNow()
	}
	if len(ts.filterImg.images) != 4 {
		t.FailNow()
	}
}
func TestTsOpParsingFailsWithMissingFilterFile(t *testing.T) {
	_, err := initTimeshift("filterimg=./temp/img01.png:range=4")
	if err == nil {
		t.FailNow()
	}
}

//TestTsOpParsingFailsWithMissingOption

//TestTsOpParsingFailsWithInvalidOption

//TestTsOpParsingFailsWithMissingValues
