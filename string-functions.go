package main

import (
	"fmt"
	s "strings"
	"time"
)

var p = fmt.Println

func main() {
	p("contains", s.Contains("test", "es"))
	p("conut", s.Count("test", "t"))
	p("hasprefix: ", s.HasPrefix("test", "te"))
	p("hassuffix: ", s.HasSuffix("test", "st"))

	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
	p()

	p("Len: ", len("hello"))
	p("Char:", "hello"[1])

	t := time.Now()
	fmt.Printf("%d-%02d-%02d", t.Year(), t.Month(), t.Day())
}
