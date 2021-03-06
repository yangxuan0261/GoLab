package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

func Test_001(t *testing.T) {
	log.Println("aaa:", "123", true) // 2019/10/03 15:47:38 aaa: 123 true
	log.Printf("bbb：%s\n", "wolegequ")
}

func Test_setflag(t *testing.T) {
	log.SetPrefix("【UserCenter】") // 设置日志前缀
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	Test_001(t) // 【UserCenter】2019/10/03 07:49:39 test_flag.go:6: aaa: 123 true
}

func Test_logFile(t *testing.T) {
	// 打开日志文件
	// 第二个参数为打开文件的模式，可选如下：
	/*
	   O_RDONLY // 只读模式打开文件
	       O_WRONLY // 只写模式打开文件
	       O_RDWR   // 读写模式打开文件
	       O_APPEND // 写操作时将数据附加到文件尾部
	       O_CREATE // 如果不存在将创建一个新文件
	       O_EXCL   // 和O_CREATE配合使用，文件必须不存在
	       O_SYNC   // 打开文件用于同步I/O
	       O_TRUNC  // 如果可能，打开时清空文件
	*/
	// 第三个参数为文件权限，请参考linux文件权限，664在这里为八进制，代表：rw-rw-r--
	logFile, err := os.OpenFile("e:/go.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// 第一个参数为输出io，可以是文件也可以是实现了该接口的对象，此处为日志文件；第二个参数为自定义前缀；第三个参数为输出日志的格式选项，可多选组合
	// 第三个参数可选如下：
	/*
	   Ldate         = 1             // 日期：2009/01/23
	       Ltime         = 2             // 时间：01:23:23
	       Lmicroseconds = 4             // 微秒分辨率：01:23:23.123123（用于增强Ltime位）
	       Llongfile     = 8             // 文件全路径名+行号： /a/b/c/d.go:23
	       Lshortfile    = 16            // 文件无路径名+行号：d.go:23（会覆盖掉Llongfile）
	       LstdFlags     = Ldate | Ltime // 标准logger的初始值
	*/
	debugLog := log.New(logFile, "[debug]", log.Ldate|log.Ltime|log.Llongfile)

	// debugLog.SetOutput(os.Stdout) // 设置之后就不会输出文件中

	// 日志输出
	debugLog.Print("日志测试Print输出，处理同fmt.Print")
	debugLog.Println("日志测试Println输出，处理同fmt.Println")
	debugLog.Printf("日志测试%s输出，处理同fmt.Printf", "Printf")

	// 日志输出，同时直接终止程序，后续的操作都不会执行
	// debugLog.Fatal("日志测试Fatal输出，处理等价于：debugLog.Print()后，再执行os.Exit(1)")
	// debugLog.Fatalln("日志测试Fatalln输出，处理等价于：debugLog.Println()后，再执行os.Exit(1)")
	// debugLog.Fatalf("日志测试%s输出，处理等价于：debugLog.Print()后，再执行os.Exit(1)", "Fatalf")

	// 日志输出，同时抛出异常，可用recover捕捉
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("===========", r)
		}
	}()
	// debugLog.Panic("日志测试Panic输出，处理等价于：debugLog.Print()后，再执行Panic()")
	// debugLog.Panicln("日志测试Panicln输出，处理等价于：debugLog.Println()后，再执行Panic()")
	// debugLog.Panicf("日志测试%s输出，处理等价于：debugLog.Printf()后，再执行Panic()", "Panicf")

	// 设置前缀
	debugLog.SetPrefix("[info]")
	// 设置输出选项
	debugLog.SetFlags(log.LstdFlags)
	debugLog.Print("222日志测试Print输出，处理同fmt.Print")
}

func Test_multiIO(t *testing.T) {
	logFile, err := os.OpenFile("e:/go.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	writers := []io.Writer{
		logFile,
		os.Stdout,
	}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	logger := log.New(fileAndStdoutWriter, "", log.Ldate|log.Ltime|log.Lshortfile) // 既写入到文件, 又输出到控制台
	// 使用新的log对象，写入日志内容
	logger.Println("--> logger :  check to make sure it works")
}

func main() {
	// test_001()
	// test_setflag()
	// test_logFile()
	// test_multiIO()
}
