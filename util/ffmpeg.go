package util

import (
	"os/exec"
	"strings"
)

func RunFF(s, output string) {

	items := strings.Split(s, " ")
	items = append(items, "-y")
	items = append(items, ouput)
	cmd := exec.Command("ffmpeg", items...)
	cmd.CombinedOutput()
	//b, err := cmd.CombinedOutput()
	//fmt.Println(string(b), err)
}
