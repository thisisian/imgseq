package main

import (
	"errors"
	"fmt"
	"image"
	"os"
	"regexp"
	"strconv"
)

type ImgSeq struct {
	images []*os.File
	config image.Config
}

func FromString(filepath string) (ImgSeq, error) {
	var s string

	file, err := os.Open(filepath)
	if err != nil {
		s = fmt.Sprintf("Error opening file %s", filepath)
		return ImgSeq{}, errors.New(s)
	}
	defer file.Close()

	imgseq, err := fromFile(file)
	if err != nil {
		return ImgSeq{}, nil
	}
	return imgseq, nil
}

func fromFile(f *os.File) (ImgSeq, error) {
	config, _, err := image.DecodeConfig(f)
	if err != nil {
		return ImgSeq{}, err
	}

	images := []*os.File{f}
	imgseq := ImgSeq{images, config}

	filepath := f.Name()
	re := regexp.MustCompile(`(\d+)(\..*)?$`)
	match := re.FindStringSubmatch(filepath)
	if len(match) > 0 {
		// Load image sequence
		i, err := strconv.Atoi(match[1])
		if err != nil {
			panic("Error processing filename")
		}
		width := len(match[1])
		extention := match[2]
		format := fmt.Sprintf("%%0%dd%s", width, extention)
		for err == nil {
			i++
			nums := fmt.Sprintf(format, i)
			nextFilePath := re.ReplaceAllString(filepath, nums)
			err = imgseq.append(nextFilePath)
		}
		fmt.Println(err)
	}
	return imgseq, nil
}

func (imgseq *ImgSeq) append(filepath string) error {
	var s string
	// Check if file exists and is a valid image
	file, err := os.Open(filepath)
	if err != nil {
		s = fmt.Sprintf("Error opening file %s", filepath)
		return errors.New(s)
	}
	defer file.Close()
	config, _, err := image.DecodeConfig(file)
	if err != nil {
		return err
	}
	if config != imgseq.config {
		// Image does not have same dimentions or
		s = fmt.Sprintf("Image %s does not match sequence", filepath)
		return errors.New(s)
	}
	imgseq.images = append(imgseq.images, file)
	return nil
}
