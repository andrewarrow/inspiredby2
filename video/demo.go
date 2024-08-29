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

	items := []string{}
	more := []string{}
	for _, file := range files {
		if !file.IsDir() {

			name := file.Name()
			if strings.HasSuffix(name, ".mp4") == false {
				continue
			}
			tokens := strings.Split(name, "_")
			id := tokens[0]
			fmt.Println(id)
			items = append(items, id)
			file2 := dir + "/" + name
			file1 := dir + "/foo/three_fixed_" + id + "_poster.jpg.mp4"
			more = append(more, file2)
			more = append(more, file1)
			//file3 := dir + "/foo3/" + id + ".mp4"
		}
	}

	for i, file := range more {
		file2 := fmt.Sprintf(dir+"/foo3/%06d.mp4", i)
		copyFile12(file, file2)
	}
}

// three_fixed_31a5d07b-b9d8-46cb-a0fe-2f7058ee03aa_poster.jpg.mp4
