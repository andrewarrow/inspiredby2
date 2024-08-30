package video

import (
	"fmt"
	"inspiredby2/pika"
	"inspiredby2/util"
	"io/ioutil"
	"sort"
	"strings"
	"time"

	"github.com/andrewarrow/feedback/router"
)

type DemoThing struct {
	Tag string
	Key string
}

func Demo(c *router.Context) {
	buffer := []DemoThing{}
	for i := 0; i < 55; i++ {
		for j := 0; j < 6; j++ {
			key := fmt.Sprintf("%d_%d", i, j)
			tag := pika.Lookup[key]
			tokens := strings.Split(tag, " ")
			words := []string{}
			for _, word := range tokens {
				if strings.Contains(word, "'") {
					continue
				}
				if len(word) < 5 {
					continue
				}
				words = append(words, word)
			}
			dt := DemoThing{strings.Join(words, " "), key}
			buffer = append(buffer, dt)
		}
	}

	for {
		peel := buffer[0:9]
		for i, item := range peel {
			id := pika.Generate("", item.Tag)
			c.FreeFormUpdate("update link_sections set id_pika=$1, prompt_text=$2 where section=$3", id, item.Tag, "1_"+item.Key)
			fmt.Println(i, item)
		}

		buffer = buffer[9:]
		if len(buffer) < 9 {
			break
		}
		time.Sleep(time.Second * 90)
	}
}

func Demo9(i int, prompts []string) {
	mapItems := map[string]int{}
	mapItemVideos := map[string]string{}
	mapItemIds := map[string]string{}
	fmt.Println("***", i, prompts)
	for _, k := range prompts {
		tag := fmt.Sprintf("Moody " + k)
		mapItems[tag] = 1
		pika.Generate("", tag)
		time.Sleep(time.Second * 1)
	}
	count := 0
	done := false
	for {
		time.Sleep(time.Second * 9)
		items, _ := pika.List("")
		for _, item := range items {
			fmt.Println(i, item.Id, item.Status)
			if item.Status == "finished" && item.Duration == 3 && mapItems[item.PromptText] == 1 {
				mapItems[item.PromptText] = 2
				mapItemVideos[item.PromptText] = item.Video
				mapItemIds[item.PromptText] = item.Id
			} else if item.Status == "finished" && item.Duration == 7 && mapItems[item.PromptText] == 3 {
				util.Download(item.Id, item.Video)
				pika.Delete(item.Id)
				pika.Delete(mapItemIds[item.PromptText])
				count++
				if count == 9 {
					done = true
					break
				}
			}
		}
		if done {
			break
		}
		for k, v := range mapItems {
			if v == 2 {
				mapItems[k] = 3
				pika.Generate(mapItemVideos[k], k)
			}
		}

	}
}

func Demo3() {
	prompts := pika.FindPrompts()
	prompts = prompts[100:]
	sort.Strings(prompts)
	i := 0
	for {
		if len(prompts) < 9 {
			break
		}
		Demo9(i, prompts[0:9])
		prompts = prompts[9:]
		time.Sleep(time.Second)
		i++
	}
}

func DemoGetOne(i int, k string) {

	done := false
	var pi pika.PikaInfo
	fmt.Println(i, k)
	tag := fmt.Sprintf("Moody " + k)
	pika.Generate("", tag)
	for {
		if done {
			break
		}
		time.Sleep(time.Second * 9)
		items, _ := pika.List("")
		for _, item := range items {
			fmt.Println(i, item.Id, item.Status)
			if item.Status == "finished" && item.PromptText == tag {
				done = true
				pi = item
				break
			}
		}
	}
	fmt.Println(i, pi.Id, pi.Video, pi.PromptText)
	pika.Generate(pi.Video, pi.PromptText)
	WaitFor7SecondVideo(pi.Id, tag)
}

func WaitFor7SecondVideo(id, tag string) {
	done := false
	var pi pika.PikaInfo
	for {
		if done {
			break
		}
		time.Sleep(time.Second * 9)
		items, _ := pika.List("")
		for _, item := range items {
			fmt.Println("1", item.Id, item.Status, item.Duration)
			if item.Status == "finished" && item.Duration == 7 && item.PromptText == tag {
				done = true
				pi = item
				break
			}
		}
	}
	//fmt.Println("2", pi.Video)
	util.Download(pi.Id, pi.Video)
	pika.Delete(id)
	pika.Delete(pi.Id)
}

func DeleteAll() {
	items, _ := pika.List("")
	fmt.Println(items)
	for _, item := range items {
		pika.Delete(item.Id)
		time.Sleep(time.Second * 1)
	}
}

func Demo2() {
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

	for i := len(more) - 1; i > -1; i-- {
		file := more[i]
		file2 := fmt.Sprintf(dir+"/foo3/%06d.mp4", i)
		copyFile12(file, file2)
	}

	//for i, file := range more {
	//}
}

// three_fixed_31a5d07b-b9d8-46cb-a0fe-2f7058ee03aa_poster.jpg.mp4
