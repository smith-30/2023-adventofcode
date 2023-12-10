package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
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

	result := 0

	type Card struct {
		LeftNumbers map[int]struct{}
		RightNumbers []int
	}

	cards := []Card{}
	hitCache := map[int]int{}
	cardsRe := map[int]int{}

	idx := 1
	// 行ごとにスキャン
	for scanner.Scan() {
		line := scanner.Text()
		// "|" を区切りにして左右の数字を取得
		leftNumbers, rightNumbers, err := extractNumbersFromCard(line)
		if err != nil {
			return
		}
		cards = append(cards, Card{
			LeftNumbers: leftNumbers,
			RightNumbers: rightNumbers,
		})
		hitCache[idx] = 0
		for _, item := range rightNumbers {
			if _, ok := leftNumbers[item]; ok {
				hitCache[idx]++
			}
		}
		cardsRe[idx]++
		idx++
	}

	for i := 1; i < idx; i++ {
		cardNum := cardsRe[i]
		hitNum, ok := hitCache[i]
		if !ok {
			continue
		}
		for {
			if cardNum == 0 {
				break
			}
			for j := i+1; j < i+1+hitNum; j++  {
				cardsRe[j]++
			}
			cardNum--
		}
		
	}

	for _, item := range cardsRe {
		result += item
	}

	
	fmt.Printf("result ------> %#v\n", result)
}

func extractNumbersFromCard(input string) (map[int]struct{}, []int, error) {
	// Card: ** 部分を無視する
	_parts := strings.Split(input, ":")
	// "|" を区切りに文字列を分割
	parts := strings.Split(_parts[1], "|")
	if len(parts) != 2 {
		return nil, nil, fmt.Errorf("invalid input format")
	}

	re := map[int]struct{}{}
	// 左辺と右辺の数字を抽出
	leftNumbers, err := extractNumbers(parts[0])
	if err != nil {
		return nil, nil, err
	}
	for _, item := range leftNumbers {
		re[item] = struct{}{}
	}

	rightNumbers, err := extractNumbers(parts[1])
	if err != nil {
		return nil, nil, err
	}

	return re, rightNumbers, nil
}

func extractNumbers(input string) ([]int, error) {
	// 正規表現パターンで数値を抽出
	re := regexp.MustCompile("\\d+")
	matches := re.FindAllString(input, -1)

	// 文字列を整数に変換
	var numbers []int
	for _, match := range matches {
		number, err := strconv.Atoi(match)
		if err == nil {
			numbers = append(numbers, number)
		}
	}

	return numbers, nil
}