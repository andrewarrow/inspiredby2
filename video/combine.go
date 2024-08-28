package video

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func CombineTwoFiles(dir, file1, file2, file3 string) {
	tmpFile1 := dir + "/" + file1
	tmpFile2 := dir + "/" + file2

	normalizeCmd1 := exec.Command("ffmpeg", "-i", file1, "-c:v", "libx264", "-c:a", "aac", "-ar", "48000", "-ac", "2", "-strict", "experimental", "-y", tmpFile1)
	normalizeCmd2 := exec.Command("ffmpeg", "-i", file2, "-c:v", "libx264", "-c:a", "aac", "-ar", "48000", "-ac", "2", "-strict", "experimental", "-y", tmpFile2)

	_, err := normalizeCmd1.CombinedOutput()
	if err != nil {
		fmt.Println("Error during normalization of file1:", err)
		return
	}

	_, err = normalizeCmd2.CombinedOutput()
	if err != nil {
		fmt.Println("Error during normalization of file2:", err)
		return
	}

	line1 := fmt.Sprintf("file '%s'", tmpFile1)
	line2 := fmt.Sprintf("file '%s'", tmpFile2)
	lines := []string{line1, line2}
	data := strings.Join(lines, "\n")
	ioutil.WriteFile("/Users/aa/list.txt", []byte(data), 0644)

	concatCmd := exec.Command("ffmpeg", "-f", "concat", "-safe", "0", "-i", "/Users/aa/list.txt", "-c", "copy", "-y", file3)
	_, err = concatCmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error during concatenation:", err)
	}
}

func CombineTwoFilesOld(file1, file2, file3 string) {
	line1 := fmt.Sprintf("file '/Users/aa/os/inspiredby2/%s'", file1)
	line2 := fmt.Sprintf("file '/Users/aa/os/inspiredby2/%s'", file2)
	lines := []string{line1, line2}
	data := strings.Join(lines, "\n")
	ioutil.WriteFile("/Users/aa/list.txt", []byte(data), 0644)
	cmd := exec.Command("ffmpeg", "-f", "concat", "-safe", "0", "-i", "/Users/aa/list.txt", "-c:v", "libx264", "-c:a", "aac", "-strict", "experimental", "-b:a", "192k", "-y", file3)
	cmd.CombinedOutput()
	//b, err := cmd.CombinedOutput()
	//fmt.Println(string(b), err)
}

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
			name := file.Name()
			if name == ".DS_Store" {
				continue
			}
			if strings.HasSuffix(name, ".mp3") {
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
