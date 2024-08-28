package video

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func Splice(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			if file.Name() == ".DS_Store" {
				continue
			}
			videoPath := filepath.Join(dir, file.Name())
			name := file.Name()
			tokens := strings.Split(name, ".")
			num := tokens[0]
			numInt, _ := strconv.Atoi(num)
			if numInt%2 == 0 && numInt > 0 {
				d, _ := GetVideoDuration(videoPath)
				fmt.Println(num, numInt, d)
				makeAudioFromSmall(dir, videoPath, name)
			}
		}
	}
}

func makeAudioFromSmall(dir, path, name string) {
	output := dir + "/" + name + ".mp3"
	cmd := exec.Command("ffmpeg",
		"-i", path,
		"-y",
		output)
	cmd.CombinedOutput()
}
