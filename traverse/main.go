package main

import (
	"flag"
	"traverse"
)

var rootDir *string = flag.String("d", "./tmp/", "traverse dir")
var ignoreDir *string = flag.String("i", "", "ignore dir when traverse")
var outFile *string = flag.String("o", "", "ignore dir when traverse")

func main() {
	flag.Parse()
	//	fmt.Println("-d 制定遍历文件  -i 制定忽略文件 -o 输出文件")
	traverse.TraverseFile(*rootDir, *ignoreDir, *outFile)
}
