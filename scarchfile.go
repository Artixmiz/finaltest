package main

import (
	"os"
	"strings"
	"sync"
)

func getDrives() (r []string) {
	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		f, err := os.Open(string(drive) + ":\\")
		if err == nil {
			d := string(drive) + ":/"
			r = append(r, d)
			f.Close()
		}
	}
	return
}
func ProcessingExtension(dir string, f os.FileInfo, extension map[string]string, files *[]MyFile, wg *sync.WaitGroup) {
	defer wg.Done()
	filename := f.Name()
	index := strings.LastIndex(filename, ".")
	if index < 0 {
		return
	}
	index = index + 1
	size := len(filename)
	ext := filename[index:size]
	_, ok := extension[ext]
	if ok {
		var mf MyFile
		mf.Path = dir + "/" + f.Name()
		mf.Size = f.Size()
		mf.Name = f.Name()
		mf.ModTime = f.ModTime()
		*files = append(*files, mf)
	}
}
