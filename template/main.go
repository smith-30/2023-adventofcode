package main

import (
	"bufio"
	"fmt"
	"os"
)

var input = "example_input"

func main() {
	// ファイルを開く
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// bufio.Scannerを作成
	scanner := bufio.NewScanner(file)

	var result int

	// 行ごとにスキャン
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("%#v\n", line)
	}

	fmt.Printf("result ------> %#v\n", result)
}