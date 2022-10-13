package termer

import (
	"fmt"
	"strings"

	"github.com/iwickgames/termer/array"
)

/*
Clamps a string to a certain size and appends "..." if it goes over
*/
func Clamp(maxsize int, text string) string {
	if maxsize >= len(text) {
		return text
	}

	return text[:maxsize-3] + "..."
}

/*
Offsets text by a given amount
*/
func Offset(amount int, text string) {
	var multiline bool = false
	if strings.ContainsRune(text, '\n') {
		multiline = true
	}

	if !multiline {
		fmt.Println(strings.Repeat(" ", amount) + text)
		return
	}

	for _, v := range strings.Split(text, "\n") {
		fmt.Println(strings.Repeat(" ", amount) + v)
	}
}

/*
Centeres text with a given width
*/
func CenterText(width int, text string) string {
	var length int = len([]rune(text))
	var screenhalf float32 = float32(width) / 2
	var spaces int = int(screenhalf - (float32(length) / 2))
	return strings.Repeat(" ", spaces) + text
}

/*
Handles multiline CenterText outputs and prints to STDOUT
*/
func (t Terminal) PrintCenter(text string) {
	var multiline bool = false
	if strings.ContainsRune(text, '\n') {
		multiline = true
	}

	if !multiline {
		fmt.Println(CenterText(t.Width, text))
		return
	}

	for _, v := range strings.Split(text, "\n") {
		fmt.Println(CenterText(t.Width, v))
	}
}

/*
Creates an outlined box
*/
func (t Terminal) CreateBox(width int, lines []string) string {
	var boxstring string
	var maxwidth = width + 2

	boxstring += ("╔" + strings.Repeat("═", maxwidth-2) + "╗") + "\n"

	for _, v := range lines {
		if (len([]rune(v)) + 2) > maxwidth {
			panic("TextOverMaxWidth: A line supplied to CreateBox was longer than its defined width")
		}

		boxstring += ("║" + v + strings.Repeat(" ", maxwidth-(len([]rune(v))+2)) + "║") + "\n"
	}

	boxstring += ("╚" + strings.Repeat("═", maxwidth-2) + "╝") + "\n"

	return boxstring
}

/*
Create an array element
*/
func (t Terminal) CreateArray(colums []array.ArrayColumn) string {
	var maxwidth int = 0
	var maxheight int = 0
	for _, arr := range colums {
		maxwidth += arr.Width
		if len(arr.Values) > maxheight {
			maxheight = len(arr.Values)
		}
	}
	if maxwidth > t.Width {
		panic("ArrayTotalWidthOverMax: The total array width was greater than the terminal width")
	}

	var arraystring string

	arraystring += "╔"
	for num, arr := range colums {
		arraystring += strings.Repeat("═", arr.Width)

		if len(colums)-1 != num {
			arraystring += "╦"
		}
	}
	arraystring += "╗\n"

	var slot = 0
	for {
		arraystring += "║"
		for num, arr := range colums {
			var ending string = "║"
			if num == len(colums)-1 {
				ending = "║\n"
			}

			if slot > len(arr.Values)-1 {
				arraystring += strings.Repeat(" ", arr.Width) + ending
				continue
			}

			if len(arr.Values[slot]) > arr.Width {
				panic("ArrayColumnValueOverWidth: An array value was over the defined width")
			}

			arraystring += arr.Values[slot] + strings.Repeat(" ", arr.Width-len(arr.Values[slot])) + ending
		}

		slot++

		if slot >= maxheight {
			break
		}
	}

	arraystring += "╚"
	for num, arr := range colums {
		arraystring += strings.Repeat("═", arr.Width)

		if len(colums)-1 != num {
			arraystring += "╩"
		}
	}
	arraystring += "╝\n"

	return arraystring
}
