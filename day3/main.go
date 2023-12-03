package main

import (
	"bufio"
	"fmt"
	"os"
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
		for idx, item := range _line {
			// 数字に変換できるとき
			if _, err := strconv.Atoi(item); err == nil {
				idxs[idx] = struct{}{}
				nums = append(nums, item)
				continue
			}
			switch item {
			case ".":
			default:
				sc.Symbols = append(sc.Symbols, Symbol{Index: idx})
			}
			if 0< len(nums) {
				num, _ := strconv.Atoi(strings.Join(nums, ""))
				sc.PartNumbers = append(sc.PartNumbers, PartNumber{
					ID: fmt.Sprintf("%v_%v_%v", sidx, idx, strings.Join(nums, "")),
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

	hitCache := map[string]struct{}{}
	for i := 0; i < len(ss); i++ {
		// Symbol の範囲にある PartNumber を取得
		for _, symbol := range ss[i].Symbols {
			cs := getAdjacentCoordinates(i, symbol.Index)
			for _, xy := range cs {
				for x, y := range xy {
					if _, ok := ss[x]; !ok {
						continue
					}
					for _, n := range ss[x].PartNumbers {
						if _, ok := n.Index[y]; ok {
							if _, ok := hitCache[n.ID]; !ok {
								result += n.Value
								hitCache[n.ID] = struct{}{}
							}
						}
					}
				}
			}
		}
	}

	// for i := 0; i < len(ss); i++ {
	// 	sc := ss[i]
	// 	for _, item := range sc.PartNumbers {
	// 		if _, ok := hitCache[item.ID]; !ok {
	// 			fmt.Printf("%#v -- %#v\n", item.ID, item.Value)
	// 		}
	// 	}
	// }

	fmt.Printf("result ------> %#v\n", result)
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