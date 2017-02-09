package main

import (
	"flag"
	"fmt"
	"syscall"
)

var rootDir *string = flag.String("d", "./tmp/", "traverse dir")
var ignoreDir *string = flag.String("i", "", "ignore dir when traverse")
var outFile *string = flag.String("o", "", "ignore dir when traverse")

func main() {
	flag.Parse()
	var rlimit syscall.Rlimit
	rlimit.Cur = uint64(100)
	rlimit.Max = uint64(1024)
	err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rlimit)
	if err != nil {
		fmt.Println(err)
		return
	}
	//C.mem_limit()
	//	fmt.Println("-d 制定遍历文件  -i 制定忽略文件 -o 输出文件")
	TraverseFile(*rootDir, *ignoreDir, *outFile)
}
