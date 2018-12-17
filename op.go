package main

import (
	"errors"
	"fmt"
	"image"
	"image/png"
	"math"
	"os"
	"strconv"
	"strings"
)

type Operation interface {
	Apply(ImgSeq, opStr string)
}

type timeshift struct {
	filterImg   ImgSeq
	filterRange int
}

func initTimeshift(opStr string) (timeshift, error) {
	opts, vals, err := parseOptions(opStr)
	if err != nil {
		return timeshift{}, err
	}
	var filterImg ImgSeq
	var filterRange int
	var filterImgFlag, filterRangeFlag = false, false
	for i, o := range opts {
		switch o {
		case "filterimg":
			filterImg, err = initImgSeq(vals[i])
			if err != nil {
				return timeshift{}, err
			}
			filterImgFlag = true
		case "range":
			filterRange, err = strconv.Atoi(vals[i])
			if err != nil {
				return timeshift{}, err
			}
			filterRangeFlag = true
		default:
			return timeshift{}, errors.New("Invalid option")
		}
	}
	if !(filterRangeFlag || filterImgFlag) {
		return timeshift{}, errors.New("Missing options")
	}
	return timeshift{filterImg, filterRange}, nil
}

// createShiftMap takes a an index to a single frame in  a timeshift object and creates an associative map
// shift distance -> pixel index
func (t timeshift) createShiftMap(i int) (map[uint][]image.Point, error) {
	f, err := os.Open(t.filterImg.images[i])
	if err != nil {
		return nil, err
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	shifts := make(map[uint][]image.Point)
	bounds := img.Bounds()
	maxV := float64(math.MaxUint16)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			// Convert to grayscale and calculate shift distance
			dist := uint((0.2989*float64(r)/maxV +
				0.5870*float64(g)/maxV +
				0.1140*float64(b)/maxV) * float64(t.filterRange))
			pts, ok := shifts[dist]
			if !ok {
				shifts[dist] = []image.Point{image.Point{x, y}}
			} else {
				shifts[dist] = append(pts, image.Point{x, y})
			}
		}
	}
	return shifts, nil
}

// parseOptions parses opStr into parallel slices of options and values
// opStr expected format is 'option1=value1:option2=value2'
func parseOptions(opStr string) ([]string, []string, error) {
	var opts []string
	var vals []string
	split := strings.Split(opStr, ":")
	for _, op := range split {
		opSplit := strings.Split(op, "=")
		if len(opSplit) != 2 {
			return opts, vals, errors.New("Invalid option")
		}
		opts = append(opts, opSplit[0])
		vals = append(vals, opSplit[1])
	}
	return opts, vals, nil
}

func (t timeshift) Apply(imgSeq ImgSeq) (ImgSeq, error) {
	ret := ImgSeq{
		make([]string, 0, 32),
		imgSeq.config,
	}

	for i := range imgSeq.images {
		curr := image.NewRGBA(image.Rect(0, 0, imgSeq.config.Width, imgSeq.config.Height))
		// Create shift map from corresponding image from filter sequence
		shifts, err := t.createShiftMap(i % len(t.filterImg.images))
		if err != nil {
			return ret, err
		}
		// Copy pixels from img corresponding to shift distance
		for dist, pts := range shifts {
			srcF, err := os.Open(imgSeq.images[(dist+uint(i))%(uint(len(imgSeq.images)))])
			if err != nil {
				return ret, err
			}
			src, _, err := image.Decode(srcF)
			if err != nil {
				return ret, err
			}
			for _, pt := range pts {
				x := pt.X
				y := pt.Y
				color := src.At(x, y)
				curr.Set(x, y, color)
			}
			srcF.Close()
		}
		// Save the image
		numformat := fmt.Sprintf("%%0%dd", base10Width(uint(len(imgSeq.images))))
		num := fmt.Sprintf(numformat, i)
		outPath := fmt.Sprintf("out%s.png", num)
		outF, err := os.Create(outPath)
		if err != nil {
			return ret, err
		}
		err = png.Encode(outF, curr)
		if err != nil {
			return ret, err
		}
		outF.Close()
	}
	return ImgSeq{}, nil
}
