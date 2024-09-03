package util

import (
	"fmt"
	"os/exec"
	"strings"
)

func RunFF(s, output string) {

	items := strings.Split(s, " ")
	items = append(items, "-y")
	items = append(items, output)
	fmt.Println(strings.Join(items, " "))
	cmd := exec.Command("ffmpeg", items...)
	cmd.CombinedOutput()
	//b, err := cmd.CombinedOutput()
	//fmt.Println(string(b), err)
}
