package traverse

import (
	"bufio"
	"crypto/sha1"
	"flag"

	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

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
		dat, _ := ioutil.ReadFile(path)
		h := sha1.New()
		h.Write(dat)
		bs := h.Sum(nil)

		_, _ = fmt.Fprintf(w, "%s,%x,%d\n", path, bs, fileInfo.Size())

		w.Flush()
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

var rootDir *string = flag.String("d", "./tmp/", "traverse dir")
var ignoreDir *string = flag.String("i", "", "ignore dir when traverse")
var outFile *string = flag.String("o", "", "ignore dir when traverse")

func main() {
	flag.Parse()
	//	fmt.Println("-d 制定遍历文件  -i 制定忽略文件 -o 输出文件")
	TraverseFile(*rootDir, *ignoreDir, *outFile)
}
