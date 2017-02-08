package main

import "testing"

import "os"
import "crypto/sha1"
import "bufio"
import "strings"
import "fmt"
import "strconv"

func TestTraverse(t *testing.T) {
	TraverseFile("./tmp", "igno*", "./traverse_result")
	file, err := os.Open("./traverse_result")
	if err != nil {
		t.Error(err)
	}

	defer file.Close()
	buf := bufio.NewReader(file)
	line, _, _ := buf.ReadLine()
	lineArr := strings.Split(string(line), ",")
	fileLen, _ := strconv.Atoi(lineArr[2])

	h := sha1.New()
	byteData := []byte("test TraveseFile function")
	h.Write(byteData)
	fmt.Println(byteData)
	data := fmt.Sprintf("%x", h.Sum(nil))

	if "tmp/test" != lineArr[0] || data != lineArr[1] || len(byteData) != fileLen {
		t.Errorf("expect %s,%s,%d got %s,%s,%d", lineArr[0], lineArr[1], line[2])
	}

}
