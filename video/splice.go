package video

import (
	"fmt"
	"io/ioutil"
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
			//videoPath := filepath.Join(dir, file.Name())
			name := file.Name()
			tokens := strings.Split(name, ".")
			num := tokens[0]
			numInt, _ := strconv.Atoi(num)
			if numInt%2 == 0 && numInt > 0 {
				fmt.Println(num)
			}
		}
	}
}
