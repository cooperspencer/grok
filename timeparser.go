package grok

import (
	"errors"
	"time"
)

var (
	formats = []string{
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
		"01/02/2006",
		"02/Jan/2006:15:04:05 -0700",
		"02/Jan/2006:15:04:05",
	}
)

func (g *Grok) AddFormat(format string) {
	formats = append(formats, format)
}

func (g *Grok) parseTime(timestring string) (time.Time, error) {
	for _, format := range formats {
		t, err := time.Parse(format, timestring)
		if err == nil {
			return t, nil
		}
	}
	return time.Time{}, errors.New("can't find a format for '" + timestring + "'")
}
