package main

import (
	"errors"
	"strings"
)

type Operation interface {
	Apply(ImgSeq, opStr string)
}

type timeshift struct {
	filterimg ImgSeq
}

func CreateTimeshift(opStr string) (timeshift, error) {
	opts, vals, err := parseOptions(opStr)
	if err != nil {
		return timeshift{}, err
	}
	var filterimg ImgSeq
	for i, o := range opts {
		switch o {
		case "filterimg":
			filterimg, err = FromString(vals[i])
			if err != nil {
				return timeshift{}, err
			}
		default:
			return timeshift{}, errors.New("Invalid option")
		}
	}
	return timeshift{filterimg}, nil
}

func (t timeshift) Apply(imgSeq ImgSeq) (ImgSeq, error) {
	return ImgSeq, nil // TODO
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
		opts := append(opts, opSplit[0])
		vals := append(vals, opSplit[1])
	}
	return opts, vals, nil
}
