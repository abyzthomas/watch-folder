package main

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"time"

	"github.com/whiteShtef/clockwork"
)

type fileItem struct {
	Path string
	Info os.FileInfo
}

type files struct {
	Files []fileItem
}

var oldList files

var sched = clockwork.NewScheduler()

var monitorDir string

func main() {

	if len(os.Args) == 1 {
		fmt.Println("\nError: Please provide the path to the folder that you want to watch!\n")
		os.Exit(1)
	}
	monitorDir = os.Args[1]

	if _, err := os.Stat(monitorDir); !os.IsNotExist(err) {

		if err != nil {
			fmt.Println("\nError: Invalid Folder Path Given to Watch!\n")
			os.Exit(1)
		}

		scheduleJobs()

		sched.Run()

	}

}

func compareFileLists(old files, new files) {

	if reflect.DeepEqual(old, new) {
		// No file Changes in the folder
	} else {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "Something Changed in the folder")
		printFileList(new)
	}
}

func scheduleJobs() {
	sched.Schedule().Every(30).Seconds().Do(checkPath)
}

func checkPath() {

	var newList files

	filepath.Walk(monitorDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		var item fileItem
		item.Path = path
		item.Info = info
		newList.Files = append(newList.Files, item)

		return nil
	})

	compareFileLists(oldList, newList)
	oldList = newList

}

func printFileList(listObj files) {
	for _, v := range listObj.Files {

		fmt.Println(v.Info.ModTime(), v.Path, v.Info.Size())

	}
}
