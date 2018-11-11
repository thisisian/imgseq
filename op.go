package main

import (
	"errors"
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
			filterImg, err = initImgSeqString(vals[i])
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
		return timeshift{}, errors.New("Missing option or options")
	}
	return timeshift{filterImg, filterRange}, nil
}

func (t timeshift) Apply(imgSeq ImgSeq) (ImgSeq, error) {
	for i, v := imgSeq.images {
		f, err := os.Open(v)
		if err != nil {
			return ImgSeq{}, err
		}
	}
}

// createShiftMap takes a single in a timeshift object and creates an associative map
// shift distance -> pixel location
func (t timeshift) createShiftMap(i int) (map[int]int) {

}

// Parses options into parallel slices of options and values
// Options will be in format 'option1=value1:option2=value2' etc
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
