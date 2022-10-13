package termer

import (
	"errors"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

type Terminal struct {
	Width  int
	Height int
}

func GetTerminal() (Terminal, error) {
	if runtime.GOOS == "windows" {
		out, err := exec.Command("powershell", "-command", "&{$H=get-host;$H.ui.rawui;}", "|", "findstr", "/b", "WindowSize").Output()
		if err != nil {
			return Terminal{}, err
		}
		raw := strings.Split(
			strings.Split(
				strings.ReplaceAll(string(out), "\r\n", ""),
				": ")[1],
			",",
		)

		if len(raw) != 2 {
			return Terminal{}, errors.New("command returned an invalid responce")
		}

		w, err := strconv.Atoi(raw[0])
		if err != nil {
			return Terminal{}, errors.New("invalid responce returned")
		}

		h, err := strconv.Atoi(raw[1])
		if err != nil {
			return Terminal{}, errors.New("invalid responce returned")
		}

		return Terminal{w, h}, nil
	}

	out, err := exec.Command("echo", "$(tput cols),$(tput lines)").Output()
	if err != nil {
		return Terminal{}, errors.New("invalid responce returned")
	}

	raw := strings.Split(string(out), ",")

	if len(raw) != 2 {
		return Terminal{}, errors.New("command returned an invalid responce")
	}

	w, err := strconv.Atoi(raw[0])
	if err != nil {
		return Terminal{}, errors.New("invalid responce returned")
	}

	h, err := strconv.Atoi(raw[1])
	if err != nil {
		return Terminal{}, errors.New("invalid responce returned")
	}

	return Terminal{w, h}, nil
}
