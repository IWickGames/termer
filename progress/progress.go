package progress

import (
	"fmt"
	"math"
	"strings"
)

type ProgressBar struct {
	MaxProgress int64
	Width       int64
}

/*
Create nicly formated bar
*/
func Create(maxwidth int64, maxprogress int64) ProgressBar {
	return ProgressBar{
		MaxProgress: maxprogress,
		Width:       maxwidth,
	}
}

/*
Outputs a created progress bar with given new values of progress
*/
func (p ProgressBar) Make(text string, progress int64) string {
	if progress > p.MaxProgress {
		panic("progress was higher than defined MaxProgress for ProgressBar")
	}

	var filled int = int(math.Round(
		float64(p.Width) / float64(p.MaxProgress) * float64(progress),
	))
	var bar string = strings.Repeat(
		"=",
		filled,
	) + strings.Repeat(
		" ",
		(int(p.Width)-filled),
	)
	var percent int = int(math.Round((100.0 / float64(p.MaxProgress)) * float64(progress)))

	return fmt.Sprintf("%s : %d%% [%s]", text, percent, bar)
}
