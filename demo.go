package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Second)
		fmt.Println(s)
	}
}

type Engine struct {
	Start int
	Stop int
}

type Car struct {
	Engine        //包含 Engine 类型的匿名字段
}

func (c *Car) work()  {
	fmt.Println(c.Start + c.Stop)
}

/*
{ # 对象容器，下面全是这个对象中的属性。注意key全都是字符串
    "id":1,   # 文章ID号，元素，value类型为number
    "content":"hello world",  # 文章内容
    "author":{   # 子对象，文章作者
        "id":2,   # 作者ID
        "name":"userA"    # 作者名称，注意子容器结束，没有逗号
    },
    "published":true,   # 文章是否发布，布尔类型
    "label":[],         # 文章标签，没有给标签，所以空数组
    "nextPost":null,    # 下一篇文章，是对象，因为没有，所以为null
    "comments":[     # 文章评论，因为可能有多条评论，每条评论都是一个对象结构
        {     # 对象容器，表示评论对象
            "id":3,        # 评论的ID号
            "content":"good post1",    # 评论的内容
            "author":"userB"         # 评论者
        },
        {
            "id":4,
            "content":"good post2",
            "author":"userC"
        }
    ]
}
*/

type Author struct {
	Id 		uint32 		`json:"id"`
	Name 	string 		`json:"name"`
}

type Comment struct {
	Id uint32 			`json:"id"`
	Content string 		`json:"content"`
	Author string 		`json:"author"`
}

type A struct {
	Id uint32 				`json:"id"`
	Content string 			`json:"content"`
	Author Author 			`json:"author"`
	Published bool 			`json:"published"`
	Label []string			`json:"label"`
	NextPost interface{}	`json:"nextPost"`
	Comments []Comment     	`json:"comments"`
}

func main() {
	go say("world")
	say("hello")

	var a string = "aaa"

	fmt.Println("aaa" + a + "bbb")

	//json
	author := Author{
		Id:14,
		Name:"xx",
	}
	comment := Comment{
		Id:      0,
		Content: "ss",
		Author:  "ss",
	}
	comments := []Comment{comment}

	t := []string{}

	one := A {
		Id:  1,
		Content: "hello world",
		Author: author,
		Published: true,
		Label: t,
		NextPost: nil,
		Comments:comments,
	}

	ss, err := json.MarshalIndent(one, "", "\t")

	if err != nil {

	}

	fmt.Println(string(ss))
}
