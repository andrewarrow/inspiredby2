package video

import (
	"fmt"
	"io/ioutil"
)

func Splice(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}

	for i, file := range files {
		if !file.IsDir() {
			if file.Name() == ".DS_Store" {
				continue
			}
			//videoPath := filepath.Join(dir, file.Name())
			name := file.Name()
			fmt.Println(i, name)
		}
	}
}
