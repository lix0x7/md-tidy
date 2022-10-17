package main

import (
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var (
	help    = flag.String("help", "", "展示帮助信息")
	isCheck = flag.Bool("check", false, "仅检查待清除的图片，不实际进行删除操作")
)

func main() {
	flag.Parse()
	removeUselessImages(*isCheck)
}

func removeUselessImages(isCheck bool) {
	docs := listDocFiles()
	images := listImgFiles()
	imageMap := make(map[string]bool, 128) // 图像资源的文件名
	for _, image := range images {
		imageMap[image] = false
	}

	docContents := make([]string, 100) // 存储文档的内容，用于之后检查图像是否存在
	for _, doc := range docs {
		content, _ := ioutil.ReadFile(doc)
		docContents = append(docContents, string(content))
	}

	// iter
	fmt.Println("start to purge useless images...")
	for image, _ := range imageMap {
		for _, content := range docContents {
			if strings.Contains(content, path.Base(image)) {
				imageMap[image] = true
				break
			}
		}
	}

	// find useless images
	count := 0
	for image, useful := range imageMap {
		if !useful {
			err := purge(image, isCheck)
			if err != nil {
				continue
			}
			count++
		}
	}
	fmt.Println("purged", count, "useless images.")
}

func listImgFiles() (files []string) {
	return listFiles([]string{".png", ".jpg", ".jpeg"})
}

func listDocFiles() (files []string) {
	return listFiles([]string{".md", ".txt"})
}

// 列出当前目录下制定后缀结尾的文件路径
func listFiles(suffixes []string) (files []string) {
	m := make(map[string]bool)
	for _, suffix := range suffixes {
		m[suffix] = true
	}

	err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() && m[filepath.Ext(info.Name())] {
			files = append(files, path)
			fmt.Println(path)
		}
		return nil
	})
	if err != nil {
		return nil
	}

	return
}

// 标记删除目标路径的文件
func markRemove(path string) (err error) {
	return
}

// 删除目标路径的文件
func purge(path string, isCheck bool) (err error) {
	if !isCheck {
		err = os.Remove(path)
		fmt.Println("removed " + path)
	} else {
		fmt.Println("[check mode]should remove " + path)
	}
	if err != nil {
		fmt.Println("error when remove file, err: ", err, "path: ", path)
	}
	return
}
