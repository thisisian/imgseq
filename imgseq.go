package main

import (
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

func initImgSeqString(filepath string) (ImgSeq, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return ImgSeq{}, err
	}
	defer file.Close()

	imgseq, err := initImgSeqFile(file)
	if err != nil {
		return ImgSeq{}, nil
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
	}
	return imgseq, nil
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
