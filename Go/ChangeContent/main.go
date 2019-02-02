// ChangeContent project main.go
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type ReplaceHelper struct {
	Root    string //根目录
	OldText string //需要替换的文本
	NewText string //新的文本
}

func (h *ReplaceHelper) DoWrok() error {

	return filepath.Walk(h.Root, h.walkCallback)

}

func (h ReplaceHelper) walkCallback(path string, f os.FileInfo, err error) error {

	if err != nil {
		return err
	}
	if f == nil {
		return nil
	}
	if f.IsDir() {
		//fmt.Pringln("DIR:",path)
		return nil
	}

	//文件类型需要进行过滤

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		//err
		return err
	}
	content := string(buf)

	fmt.Println("content:", content)
	//替换
	newContent := strings.Replace(content, h.OldText, h.NewText, -1)

	//重新写入
	ioutil.WriteFile(path, []byte(newContent), 0)

	return err
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func getCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	checkErr(err)
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	fmt.Println(path)
	return path
}
func main() {

	helper := ReplaceHelper{
		Root:    getCurrentPath(),
		OldText: "http://kong2.baijiahulian.com",
		NewText: "http://kong.baijiahulian.com",
	}
	err := helper.DoWrok()
	if err == nil {
		fmt.Println("done!")
	} else {
		fmt.Println("error:", err.Error())
	}
}
