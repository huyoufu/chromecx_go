package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
)

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.LittleEndian, &x)

	return int(x)
}

type Msg struct {
	Text string `json:text`
}

func in() string {
	//读取4个字节是 头部的内容长度
	slice_header := make([]byte, 4)
	os.Stdin.Read(slice_header)
	len_header := BytesToInt(slice_header)

	slice_body := make([]byte, len_header)
	//再次读取剩下的内容
	os.Stdin.Read(slice_body)
	msg := &Msg{}
	//是json格式的所以转换下
	json.Unmarshal(slice_body, msg)
	return msg.Text
}
func main() {
	//死循环 可以一直使用 当然你也可以使用一次性的
	for {
		msg := in()
		//s:="{\"text\":\"响应内容啊实打实"+strconv.Itoa(len_header)+msg.Text+"\"}"
		s := "{\"text\":\"价格下降了!!!:" + msg + "\"}"
		out(s)
	}
}

func out(s string) {
	//小端编码 以32位形式写出数据长度
	len1 := len(s) >> 0 & 0xFF
	len2 := len(s) >> 8 & 0xFF
	len3 := len(s) >> 16 & 0xFF
	len4 := len(s) >> 24 & 0xFF

	fmt.Printf("%c", len1)
	fmt.Printf("%c", len2)
	fmt.Printf("%c", len3)
	fmt.Printf("%c", len4)
	fmt.Printf("%s", s)
}
