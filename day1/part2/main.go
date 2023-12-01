package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// var (
// 	one, two, three, four, five, six, seven, eight, nine
// )

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
		nums := []int{}
		chars := strings.Split(line, "")
		lc := len(chars)
		for idx, char := range chars {
			// char が数値に変換できるとき
			if num, err := strconv.Atoi(char); err == nil {
				nums = append(nums, num)
				continue
			}

			switch char {
				case "o", "t", "f", "s", "e", "n":
				default:
					continue
			}

			buildStr := char
			for i := idx+1; i < lc; i++ {
				c := chars[i]
				if _, err := strconv.Atoi(c); err == nil {
					break
				}
				buildStr += c
				
				if 3 <= len(buildStr) {
					switch buildStr {
						case "one":
							nums = append(nums, 1)
						case "two":
							nums = append(nums, 2)
						case "three":
							nums = append(nums, 3)
						case "four":
							nums = append(nums, 4)
						case "five":
							nums = append(nums, 5)
						case "six":
							nums = append(nums, 6)
						case "seven":
							nums = append(nums, 7)
						case "eight":
							nums = append(nums, 8)
						case "nine":
							nums = append(nums, 9)
						default:
						}
				}
				if 5 <= len(buildStr) {
					break
				}
			}
		}

		

		l := len(nums)
		switch {
			case l == 0:
			case l == 1:
				result += nums[0] * 10 + nums[0]
			case l >= 2:
				result += nums[0] * 10 + nums[l-1]
			}
		// fmt.Printf("%#v -- %#v\n", line, nums)
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
