package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var input = "input"

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
		numbers := extractNumbers(line)
		l := len(numbers)
		switch {
			case l == 0:
			case l == 1:
				result += numbers[0] * 10 + numbers[0]
			case l >= 2:
				result += numbers[0] * 10 + numbers[l-1]
			}
	}

	fmt.Printf("result ------> %#v\n", result)

	// エラーの確認
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	// 最初の桁

	// 最後の桁

	// 最後の桁がなかった場合、最初の桁と結合した数値になる. ex. 7 -> 77


}

var (
	// 数字だけを取り出す正規表現
	re = regexp.MustCompile(`\d`)
)


func extractNumbers(input string) []int {
	// 正規表現にマッチする部分をすべて取得
	matches := re.FindAllString(input, -1)

	// 文字列を整数に変換
	var numbers []int
	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err == nil {
			numbers = append(numbers, num)
		}
	}

	return numbers
}