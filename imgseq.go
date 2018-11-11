package main

import (
	"fmt"
	"image"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type ImgSeq struct {
	images []*os.File
	config image.Config
}

// initImgSeqString returns an image sequence whose first element
// is the file located in input string.
func initImgSeqString(filepath string) (ImgSeq, error) {
	filepath = path.Clean(filepath)
	file, err := os.Open(filepath)
	if err != nil {
		return ImgSeq{}, err
	}
	defer file.Close()

	imgseq, err := initImgSeqFile(file)
	if err != nil {
		return ImgSeq{}, err
	}
	return imgseq, nil
}

func initImgSeqFile(f *os.File) (ImgSeq, error) {
	config, _, err := image.DecodeConfig(f)
	if err != nil {
		return ImgSeq{}, err
	}

	images := []*os.File{f}
	imgseq := ImgSeq{images, config}

	filepath := f.Name()
	nextFilePath := filepath
	for err == nil {
		nextFilePath, err = nextFile(nextFilePath)
		if err != nil {
			return imgseq, nil
		}
		err = imgseq.append(nextFilePath)
	}
	return imgseq, nil
}

// Returns the next file in a series from the filename
// Assumes clean filename
func nextFile(file string) (string, error) {
	ext := path.Ext(file)
	noExt := strings.TrimSuffix(path.Base(file), ext)
	dir := path.Dir(file)

	re := regexp.MustCompile(`^(\d*)(\D*)(\d*)$`)
	match := re.FindStringSubmatch(noExt)
	if len(match) == 0 ||
		len(match[1]) > 0 && len(match[3]) > 0 ||
		len(match[1]) == 0 && len(match[3]) == 0 {
		// If no match, or numerals at beginning and end of filename
		// or no numerals in filename
		return "", errors.Errorf("invalid filename: %s", file)
	}

	var nums string
	if len(match[1]) > 0 {
		nums = match[1]
	} else if len(match[3]) > 0 {
		nums = match[3]
	}

	width := len(nums)
	i, err := strconv.Atoi(nums)
	if err != nil {
		panic(fmt.Sprintf("Error parsing filename: %s\nFailed to convert integer: '%s'", file, nums))
	}
	i++

	numformat := fmt.Sprintf("%%0%dd", width)
	num := fmt.Sprintf(numformat, i)
	if len(match[1]) > 0 {
		newBase := fmt.Sprintf("%s%s%s", num, match[2], ext)
		return path.Join(dir, newBase), nil
	} else if len(match[3]) > 0 {
		newBase := fmt.Sprintf("%s%s%s", match[2], num, ext)
		return path.Join(dir, newBase), nil
	} else {
		panic("Error parsing filename")
	}
}

func (imgseq *ImgSeq) append(filepath string) error {
	// Check if file exists and is a valid image
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	config, _, err := image.DecodeConfig(file)
	if err != nil {
		return err
	}
	if config != imgseq.config {
		// Image does not have same dimentions or
		return fmt.Errorf("non-conforming image: %s", filepath)
	}
	imgseq.images = append(imgseq.images, file)
	return nil
}
