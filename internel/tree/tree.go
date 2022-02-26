package tree

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type DirNode struct {
	name   string
	rname  string
	cDirs  []*DirNode
	cFiles []*FileNode
}

type FileNode struct {
	name string
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			n.Add(1)
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func treeDir(root *DirNode, n *sync.WaitGroup) {
	defer n.Done()
	//TODO 将自动略过.git替换成可设置项
	/*	if root.name == ".git" {
		return
	}*/
	if strings.HasPrefix(root.name, ".") && root.name != "." {
		return
	}

	for _, entry := range dirents(root.rname) {
		if entry.IsDir() {
			sub := &DirNode{name: entry.Name(), rname: filepath.Join(root.rname, entry.Name())}
			root.cDirs = append(root.cDirs, sub)
			n.Add(1)
			go treeDir(sub, n)
		} else {
			f := &FileNode{entry.Name()}
			root.cFiles = append(root.cFiles, f)
		}
	}
}

//通过sema限制同时运行的dirents
var sema = make(chan struct{}, 20)

// dirents 返回所有文件的入口

func dirents(dir string) []os.FileInfo {
	//这样，同时最多只有20个dirents在运行，sema满了会阻塞
	sema <- struct{}{}
	defer func() { <-sema }()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Printf("dir error: %v", err)
		return nil
	}
	return entries
}

func (root *DirNode) print(depth int) {
	if depth == 0 {
		fmt.Printf("!---%s\n", root.name)
	}

	for i := 0; i < depth; i++ {
		fmt.Printf("   |")
	}
	if depth > 0 {
		fmt.Printf("---%s\n", root.name)
	}

	depth++
	if len(root.cDirs) > 0 {
		for _, c := range root.cDirs {
			c.print(depth)
		}
	}
	if len(root.cFiles) > 0 {
		for _, c := range root.cFiles {
			for i := 0; i < depth; i++ {
				fmt.Printf("   |")
			}
			fmt.Printf("---%v\n", c.name)
		}
	}
}

func PrintTree(dir string) {
	root := &DirNode{
		name:  dir,
		rname: dir,
	}
	var n sync.WaitGroup
	n.Add(1)
	treeDir(root, &n)
	n.Wait()
	root.print(0)
}
