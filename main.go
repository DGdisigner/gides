package main

import (
	"fmt"
	"gides/src/driver"
	"gides/src/tree"
)

func main() {
	ttt()
}

func ttt() {
	c := new(driver.CacheDriver)
	for i, n := 0, 1000; i < n; i++ { // 常见的 for 循环，支持初始化语句。
		go c.Set(fmt.Sprintf("key_%v", i), fmt.Sprintf("val_%v", i), 1000)
	}
	for i, n := 0, 1000; i < n; i++ { // 常见的 for 循环，支持初始化语句。
		go func(i int, c *driver.CacheDriver) {
			val := c.Find(fmt.Sprintf("key_%v", i))
			fmt.Println(val)
		}(i, c)
	}
}

func testTree() {
	words := map[string]string{
		"apple": "苹果",
		"book":  "书",
		"ask":   "问",
		"abort": "关于",
		"dog":   "狗子",
	}
	root := tree.CreateNodeTree(words)
	root.Traversal()
	fmt.Println("")
	var word string
	for {
		fmt.Println("请输入单词：")
		_, err := fmt.Scanln(&word)
		if err != nil {
			fmt.Println("输入错误")
		} else {
			exp, err := root.Search(word)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println(exp)
			}
		}
	}
}
