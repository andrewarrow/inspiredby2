package video

import (
	"fmt"
	"io/ioutil"
	"math/rand"
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

			name := file.Name()
			if name == ".DS_Store" {
				continue
			}
			if strings.HasSuffix(name, ".mp3") {
				continue
			}
			if strings.Contains(name, "_") {
				continue
			}
			videoPath := filepath.Join(dir, file.Name())
			tokens := strings.Split(name, ".")
			num := tokens[0]
			numInt, _ := strconv.Atoi(num)
			if numInt%2 == 0 && numInt > 0 {
				d1, _ := GetVideoDuration(videoPath)
				rightPika := findRightPika()
				d2, _ := GetVideoDuration("data4/" + rightPika)

				fmt.Println(num, numInt, d1, d2, rightPika)
				makeAudioFromSmall(dir, videoPath, num)
				removeFirstThreeSeconds(dir, videoPath, num)
				mergePikeFileAnd3SecAudio(dir, num, rightPika)
				combineToMakeGoodFile(dir, num)
			}
		}
	}
}

func mergePikeFileAnd3SecAudio(dir, name, rightPika string) {
	//ffmpeg -i video.mp4 -i audio.mp3 -c:v copy -c:a aac -map 0:v -map 1:a output.mp4
	//ffmpeg -i video.mp4 -i audio.mp3 -filter_complex "[0:a][1:a]amix=inputs=2:duration=first:dropout_transition=3[aout]" -map 0:v -map "[aout]" -c:v copy -c:a aac output.mp4

	output := dir + "/foo/" + name + "_opening3.mp4"
	cmd := exec.Command("ffmpeg",
		"-i", "/Users/aa/os/inspiredby2/data4/"+rightPika,
		"-i", dir+"/"+name+".m4a",
		"-filter_complex",
		"[0:a][1:a]amix=inputs=2:duration=first:dropout_transition=3[aout]",
		"-map", "0:v", "-map", "[aout]",
		"-c:v",
		"copy", "-c:a", "aac",
		"-y",
		output)
	//cmd.CombinedOutput()
	//fmt.Println(string(b), err)
	b, err := cmd.CombinedOutput()
	fmt.Println(string(b), err)

}

func combineToMakeGoodFile(dir, name string) {
	file2 := dir + "/foo/" + name + "_opening3.mp4"
	file1 := dir + "/" + name + "_without_first3.mp4"
	file3 := dir + "/foo2/" + name + "_ready.mp4"
	CombineTwoFiles(dir, file1, file2, file3)
}

// ffmpeg -i 000001.mp4 -t 3 -vn -acodec copy output.m4a
// ffmpeg -i 000001.mp4 -t 3 -vn -acodec aac output.m4a

func makeAudioFromSmall(dir, path, name string) {
	output := dir + "/" + name + ".m4a"
	cmd := exec.Command("ffmpeg",
		"-sseof", "-3", // Start 3 seconds before the end of the file
		"-i", path,
		"-vn", "-acodec", "aac",
		"-y",
		output)
	cmd.CombinedOutput()
}

// ffmpeg -i input.mp3 -t 3 -c copy output.mp3

// ffmpeg -i input_video.mp4 -ss 0 -t 3 -q:a 0 -map a output_audio.mp3
// ffmpeg -i input_video.mp4 -t 3 -q:a 0 -map a output_audio.mp3

// ffmpeg -ss 00:00:03 -i input.mp4 -c copy output.mp4
// ffmpeg -ss 00:00:03 -i input.mp4 -c:v libx264 -c:a aac output.mp4

func removeFirstThreeSeconds(dir, path, name string) {
	output := dir + "/" + name + "_without_first3.mp4"
	cmd := exec.Command("ffmpeg",
		"-sseof", "-3", // Start 3 seconds before the end of the file
		//"-ss", "00:00:03",
		"-i", path,
		"-c:v", "libx264",
		"-c:a", "aac",
		"-y",
		output)
	//b, err := cmd.CombinedOutput()
	cmd.CombinedOutput()
}

func findRightPika() string {
	dir := "data4"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}

	items := []string{}
	for _, file := range files {
		if !file.IsDir() {
			if file.Name() == ".DS_Store" {
				continue
			}
			name := file.Name()
			items = append(items, name)
		}
	}

	return items[rand.Intn(len(items))]
}
