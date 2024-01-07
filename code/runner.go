package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main() {
	// go run code-user/main.go
	var bm runtime.MemStats
	runtime.ReadMemStats(&bm)
	// Alloc 已申请，且仍在使用的字节
	fmt.Printf("KB: %v\n", bm.Alloc/1024)
	now := time.Now().UnixMilli()
	println("当前时间(毫秒) ==> ", now)

	cmd := exec.Command("go", "run", "code-user/main.go")
	go func() {
		for {
			if cmd.Process == nil {
				continue
			}
			GetS(strconv.Itoa(cmd.Process.Pid))
		}
	}()
	var out, stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &out
	stdinPipe, err := cmd.StdinPipe()
	if err != nil {
		log.Fatalln(err)
	}
	io.WriteString(stdinPipe, "23 11\n")
	if err := cmd.Run(); err != nil {
		log.Fatalln(err, stderr.String())
	}
	println("Err:", string(stderr.Bytes()))
	fmt.Println(out.String())

	println(out.String() == "34\n")
	var em runtime.MemStats
	runtime.ReadMemStats(&em)
	fmt.Printf("KB: %v\n", em.Alloc/1024)
	end := time.Now().UnixMilli()
	println("当前时间 ==> ", end)
	println("耗时 ==> ", end-now)
}
func GetS(pid string) int {
	// 执行ps命令获取进程内存使用情况
	output, err := exec.Command("ps", "-o", "rss=", "-p", pid).Output()
	if err != nil {
		fmt.Printf("执行ps命令出错：%s\n", err)
		return 0
	}
	// 解析输出结果，并转换为整数（单位为KB）
	memoryUsage := strings.TrimSpace(string(output))
	if usage, err := strconv.Atoi(memoryUsage); err == nil {
		fmt.Printf("进程%s的内存使用量：%d KB\n", pid, usage)
		return usage
	} else {
		fmt.Printf("解析内存使用量出错：%s\n", err)
	}
	return 0
}
