package main

import (
	"bufio"
	"crypto/sha1"
	"math"

	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

// 对于大文件的处理凡是大于5M 都按块读取。
const fileLimit = 1024 * 5

func TraverseFile(rootDir, ignoreDir, outFile string) {

	file, err := os.Create(outFile)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	w := bufio.NewWriter(file)
	err = filepath.Walk(rootDir, func(path string, f os.FileInfo, err error) error {

		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		if ignoreDir != "" {
			match, _ := regexp.MatchString(ignoreDir, path)
			if match {

				return nil
			}
		}
		fileInfo, err := os.Stat(path)
		if err != nil {
			panic(err)
		}
		fileSize := fileInfo.Size()

		h := sha1.New()
		if fileSize < fileLimit {

			dat, _ := ioutil.ReadFile(path)
			h.Write(dat)
		} else { // 超过5M 分块读取
			blockNum := uint64(math.Ceil(float64(fileSize) / float64(fileLimit)))
			readFile, err := os.Open(path)
			if err != nil {
				panic(err)
			}
			for i := uint64(0); i < blockNum; i++ {
				blockSize := int(math.Min(fileLimit, float64(fileSize-int64(i*fileLimit))))
				dat := make([]byte, blockSize)
				readFile.Read(dat)
				io.WriteString(h, string(dat))
			}
			readFile.Close()
		}

		bs := h.Sum(nil)

		_, _ = fmt.Fprintf(w, "%s,%x,%d\n", path, bs, fileInfo.Size())

		w.Flush()
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}
