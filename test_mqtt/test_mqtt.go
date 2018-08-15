package main

import (
	"GoLab/test_mqtt/work"
	"flag"
	"fmt"
	"strings"

	MQTT "github.com/eclipse/paho.mqtt.golang"

	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

func main() {
	doConn(nil, "-1", false)
	// doMultiConn()
	// doConnWithInput()
}

func doMultiConn() {
	var num int
	flag.IntVar(&num, "num", -1, "help msg for name")
	flag.Parse()

	var wg sync.WaitGroup
	cnt := 3

	if num > 0 {
		cnt = num
	}

	for index := 1; index <= cnt; index++ {
		wg.Add(1)
		go doConn(&wg, strconv.Itoa(index), true)
	}

	wg.Wait()
	fmt.Println("程序结束 666")
}

func doConn(wg *sync.WaitGroup, flag string, isPing bool) {
	c := make(chan os.Signal, 1)

	this := new(work.MqttWork)
	opts := this.GetDefaultOptions("tcp://127.0.0.1:3563")
	opts.SetConnectionLostHandler(func(client MQTT.Client, err error) {
		fmt.Println("连接断开", err.Error())
		c <- os.Kill
	})
	opts.SetOnConnectHandler(func(client MQTT.Client) {
		fmt.Println("连接成功")
	})
	err := this.Connect(opts)
	if err != nil {
		fmt.Println("连接错误:", err.Error())
		c <- os.Kill
	}

	cnt := 1
	sendFn := func() {
		//访问HelloWorld001模块的HD_Say函数
		hiStr := fmt.Sprintf(`{"say":"我是wilker%s, cnt:%d"}`, flag, cnt)
		msg, err := this.Request("HelloWorld@HelloWorld001/HD_Say", []byte(hiStr))
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(fmt.Sprintf("topic :%s  body :%s", msg.Topic(), string(msg.Payload())))
	}
	sendFn()

	if isPing {
		for {
			cnt++
			sendFn()
			time.Sleep(time.Second * 3)
		}
	}

	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	fmt.Printf("mqant closing down (signal: %v)\n", sig)
	fmt.Println("--- flag done:", flag)
	if wg != nil {
		defer wg.Done()
	}
}

func doConnWithInput() {
	rInputFn := func(c chan string) {
		for {
			is := ""
			// fmt.Println("Please enter some input: ")
			fmt.Scan(&is)
			c <- is
			// inputReader := bufio.NewReader(os.Stdin)
			// input, err := inputReader.ReadString('\n')
			// if err == nil {
			// 	c <- input
			// }
		}
	}

	dealFn := func(c chan string) {
		this := new(work.MqttWork)
		opts := this.GetDefaultOptions("tcp://127.0.0.1:3563")
		opts.SetConnectionLostHandler(func(client MQTT.Client, err error) {
			fmt.Println("连接断开", err.Error())
			// c <- os.Kill
		})
		opts.SetOnConnectHandler(func(client MQTT.Client) {
			fmt.Println("连接成功")
		})
		err := this.Connect(opts)
		if err != nil {
			fmt.Println("连接错误:", err.Error())
			// c <- os.Kill
		}

		parseFn := func(src *string) (string, string) {
			sArr := strings.Split(*src, ";")
			if len(sArr) >= 2 {
				return sArr[0], sArr[1]
			}
			return "", ""
		}

		sendFn := func(topic string, body string) {
			if topic == "" {
				topic = "HelloWorld@HelloWorld001/HD_Say"
			}
			if body == "" {
				body = "Sorry msg"
			}
			hiStr := fmt.Sprintf(`{"say":"msg:%s"}`, body)
			msg, err := this.Request("HelloWorld@HelloWorld001/HD_Say", []byte(hiStr))
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(fmt.Sprintf("topic:%s, body:%s\n", msg.Topic(), string(msg.Payload())))
		}

		for {
			select {
			case msg := <-c:
				// fmt.Println("recv msg:", msg)
				sendFn(parseFn(&msg))
			}
		}
	}

	rc := make(chan string)
	go rInputFn(rc)
	go dealFn(rc)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	fmt.Printf("mqant closing down (signal: %v)\n", sig)

}
