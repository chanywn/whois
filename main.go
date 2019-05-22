package main

import (
	"strings"
	"fmt"
	"bufio"
	"io"
	"os"
	"regexp"
	"sync"
	"time"
	"github.com/likexian/whois-go"
)

var fileSuccFile *os.File
var wgSuccFileFile sync.Mutex
var bf int
var maxBf int = 20

func main() {
	bf = 0;
	fileSuccFile, _ = os.Create("./a")
	
	dict := "./a.txt"
	f, err := os.Open(dict)
	defer f.Close()
	if (err != nil) {
		fmt.Println(err)
	}
	buff := bufio.NewReader(f) //读入缓存
	for {
		bf++;
		line, err := buff.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}
		line = strings.Trim(line, " ")
		line = strings.Trim(line, "\n")
		domain := line + ".com"
		waitingForBf()
		go match(domain);
	}
}

func waitingForBf() {
	for {
		if bf <= maxBf {
			return;
		} else {
			time.Sleep(time.Duration(1))
		}
	}
}

func match(domain string) {
	result, err := whois.Whois(domain)
	if err != nil {
		// fmt.Println(err)
	}
	reg := regexp.MustCompile("No match for")
	re := reg.FindAllString(result, -1)
	if re == nil { 
		fmt.Printf("%s 已被注册\n", domain)
	} else {
		fmt.Printf("%s 恭喜，可以被注册\n", domain)
		wgSuccFileFile.Lock()
		fileSuccFile.WriteString(domain + "\r\n")
		wgSuccFileFile.Unlock()
	}
	bf--;
}