package video

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Demo() {
	dir := "data5"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		if !file.IsDir() {

			name := file.Name()
			if strings.HasSuffix(name, ".mp4") == false {
				continue
			}
			tokens := strings.Split(name, "_")
			id := tokens[0]
			fmt.Println(id)
			file1 := dir + "/" + name
			file2 := dir + "/foo/three_fixed_" + id + "_poster.jpg.mp4"
			file3 := dir + "/foo3/" + id + ".mp4"
			CombineTwoFilesOld(file1, file2, file3)
		}
	}
}

// three_fixed_31a5d07b-b9d8-46cb-a0fe-2f7058ee03aa_poster.jpg.mp4
