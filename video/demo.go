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
			fmt.Println(name)
		}
	}
}
