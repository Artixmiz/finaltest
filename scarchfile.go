package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type MyFile struct {
	Path    string
	Size    int64
	Name    string
	ModTime time.Time
}

func getDrives(inputdrive string) (r []string) {
	for _, drive := range inputdrive {
		f, err := os.Open(string(inputdrive) + ":\\")
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
func FindFileFromExtension(extension map[string]string, dir string, files *[]MyFile, wg *sync.WaitGroup) {
	defer wg.Done()
	fs, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}
	for _, f := range fs {
		var subWg sync.WaitGroup
		subWg.Add(1)
		if f.IsDir() {
			path := dir + "/" + f.Name()
			go FindFileFromExtension(extension, path, files, &subWg)

		} else {
			ProcessingExtension(dir, f, extension, files, &subWg)
		}
		subWg.Wait()
	}
}

func txtCreate(Path, Size []string) {
	file, err := os.Create("./output.txt")
	defer file.Close()
	if err != nil {
		return
	}
	for i := range Path {
		file.WriteString(Size[i] + "\t\t" + Path[i] + "\n")
	}

}

func main() {
	var wg sync.WaitGroup
	t1 := time.Now()
	myfiles := []MyFile{}
	var inputextension string
	var inputdrive string
	extensionMap := make(map[string]string)
	fmt.Print("extension file scarch \n input :")
	fmt.Scan(&inputextension)
	if inputextension = "txt"{
		extensionMap["txt"] = ".txt"
	}else if inputextension = "jpg" {
		extensionMap["jpg"] = ".jpg"
	}
	extensionMap["jpg"] 
	extensionMap["txt"] = inputextension
	fmt.Print("select drive for scarch :")
	fmt.Scan(&inputdrive)
	drives := getDrives(inputdrive)
	wg.Add(len(drives))
	for _, drive := range drives {
		go FindFileFromExtension(extensionMap, drive, &myfiles, &wg)

	}
	wg.Wait()
	var pathFiles, sizeFiles []string
	for _, pathToFiles := range myfiles {
		pathFiles = append(pathFiles, pathToFiles.Path)
		sizeFiles = append(sizeFiles, strconv.Itoa(int(pathToFiles.Size)))
	}
	t2 := time.Now()
	diftime := t2.Sub(t1)
	txtCreate(pathFiles, sizeFiles)
	fmt.Println("total files = ", len(myfiles))
	fmt.Println("total time = ", diftime)
}
