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

	var result int
	ss := map[int]Schematic{}
	sidx := 0

	// 行ごとにスキャン
	for scanner.Scan() {
		line := scanner.Text()
		_line := strings.Split(line, "")
		sc := Schematic{}
		idxs := make(map[int]struct{})
		nums := []string{}

		_lineLen := len(_line)
		for idx, item := range _line {
			// 数字に変換できるとき
			if _, err := strconv.Atoi(item); err == nil {
				idxs[idx] = struct{}{}
				nums = append(nums, item)
				// 最後が数字だったときのケア
				if _lineLen-1 == idx {
					if 0 < len(nums) {
						num, _ := strconv.Atoi(strings.Join(nums, ""))
						sc.PartNumbers = append(sc.PartNumbers, PartNumber{
							ID: fmt.Sprintf("%v_%v_%v", sidx+1, idx, strings.Join(nums, "")),
							Index: idxs,
							Value: num,
						})
					}
				}
				continue
			}
			switch item {
			case ".":
			case "*":
				sc.Symbols = append(sc.Symbols, Symbol{
					Index: idx,
					Value: item,
				})
			default:
			}
			if 0 < len(nums) {
				num, _ := strconv.Atoi(strings.Join(nums, ""))
				sc.PartNumbers = append(sc.PartNumbers, PartNumber{
					ID: fmt.Sprintf("%v_%v_%v", sidx+1, idx, strings.Join(nums, "")),
					Index: idxs,
					Value: num,
				})
			}
			nums = []string{}
			idxs = make(map[int]struct{})
		}
		ss[sidx] = sc
		sidx++
	}

	for i := 0; i < len(ss); i++ {
		// Symbol の範囲にある PartNumber を取得
		for _, symbol := range ss[i].Symbols {
			if symbol.Value != "*" {
				continue
			}
			nums := []int{}
			cs := getAdjacentCoordinates(i, symbol.Index)
			for _, xy := range cs {
				for x, y := range xy {
					if _, ok := ss[x]; !ok {
						continue
					}
					for idx, n := range ss[x].PartNumbers {
						if _, ok := n.Index[y]; ok {
							nums = append(nums, n.Value)
							firstPart := ss[x].PartNumbers[:idx]
							secondPart := ss[x].PartNumbers[idx+1:]


							// 前半と後半を結合して新しいスライスを作成
							result := append(firstPart, secondPart...)

							sc := ss[x]
							sc.PartNumbers = result
							ss[x] = sc
						}
					}
				}
			}
			if 1 < len(nums) {
				result += nums[0] * nums[len(nums)-1]
			}
		}
	}

	fmt.Printf("result    ------> %#v\n", result)
}

type Schematic struct {
	PartNumbers []PartNumber
	Symbols []Symbol
}

type PartNumber struct {
	ID string
	Index map[int]struct{}
	Value int
}

type Symbol struct {
	Index int
	Value string
}

// 隣接する座標を取得
func getAdjacentCoordinates(idx, v int) (re []map[int]int) {
	re = []map[int]int{}
	for i := idx-1; i < idx+2; i++ {
		for j := v-1; j < v+2; j++ {
			if i == idx && j == v {
				continue
			}
			re = append(re, map[int]int{i: j})
		}
	}
	return
}

func extractNumbers(input string) []int {
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

	return numbers
}