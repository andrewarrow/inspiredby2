package video

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func Combine(dir string) {

	listFile, err := os.Create("/Users/aa/list.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer listFile.Close()

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
			fmt.Fprintf(listFile, "file '/Users/aa/os/inspiredby2/%s'\n", videoPath)
		}
	}

	if err := listFile.Close(); err != nil {
		fmt.Println(err)
	}

	cmd := exec.Command("ffmpeg", "-f", "concat", "-safe", "0", "-i", "/Users/aa/list.txt", "-c", "copy", "-y", "/Users/aa/Downloads/output.mp4")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Video concatenation completed successfully.")
}
