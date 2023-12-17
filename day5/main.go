package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var input = "input"

type LandMap struct {
	Start int
	SourceStart int
	Width int
}

type LandMapSlice []LandMap

func (a *LandMap) GetDest(index int) int {
	if a.SourceStart <= index && index <= a.SourceStart + a.Width {
		fmt.Printf("%#v %#v %#v %#v\n", index, a.Start, a.SourceStart, a.Width)
		v := a.Start - a.SourceStart
		return index + v
	}
	return index
}

func (a *LandMapSlice) GetDest(index int) int {
	re := map[int]struct{}{}
	for _, item := range *a {
		dest := item.GetDest(index)
		fmt.Printf("*** %#v --> %#v\n", index, dest)
		re[dest] = struct{}{}
	}
	if 1 < len(re) {
		for key := range re {
			if key != index {
				return key
			}
		}
	}
	return index
}

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

	// var result int

	nums := []int{}
	seedToSoils := LandMapSlice{}
	soilToFertilizers := LandMapSlice{}
	fertilizerToWater := LandMapSlice{}
	waterToLights := LandMapSlice{}
	lightToTemperatures := LandMapSlice{}
	temperatureToHumidities := LandMapSlice{}
	humidityToLocations := LandMapSlice{}

	// 行ごとにスキャン
	for scanner.Scan() {
		line := scanner.Text()
		switch {
			case strings.HasPrefix(line, "seeds"):
				_nums, _ := extractNumbersPrefix(line)
				nums = _nums
			case strings.HasPrefix(line, "seed-to-soil"):
				seedToSoils = readLandMaps(scanner)
			case strings.HasPrefix(line, "soil-to-fertilizer"):
				soilToFertilizers = readLandMaps(scanner)
			case strings.HasPrefix(line, "fertilizer-to-water"):
				fertilizerToWater = readLandMaps(scanner)
			case strings.HasPrefix(line, "water-to-light"):
				waterToLights = readLandMaps(scanner)
			case strings.HasPrefix(line, "light-to-temperature"):
				lightToTemperatures = readLandMaps(scanner)
			case strings.HasPrefix(line, "temperature-to-humidity"):
				temperatureToHumidities = readLandMaps(scanner)
			case strings.HasPrefix(line, "humidity-to-location"):
				humidityToLocations = readLandMaps(scanner)
		}
	}

	lands := []int{}
	start := "\nstart\n"
	// nums = []int{14}
	for _, item := range nums {
		fmt.Printf("item %#v\n\n", item)
		fmt.Printf(start)
		a := seedToSoils.GetDest(item)
		fmt.Printf("%#v %#v\n", "seedToSoils", a)
		fmt.Printf(start)
		a = soilToFertilizers.GetDest(a)
		fmt.Printf("%#v %#v\n", "soilToFertilizers", a)
		fmt.Printf(start)
		a = fertilizerToWater.GetDest(a)
		fmt.Printf("%#v %#v\n", "fertilizerToWater", a)
		fmt.Printf(start)
		a = waterToLights.GetDest(a)
		fmt.Printf("%#v %#v\n", "waterToLights", a)
		fmt.Printf(start)
		a = lightToTemperatures.GetDest(a)
		fmt.Printf("%#v %#v\n", "lightToTemperatures", a)
		fmt.Printf(start)
		a = temperatureToHumidities.GetDest(a)
		fmt.Printf("%#v %#v\n", "temperatureToHumidities", a)
		fmt.Printf(start)
		a = humidityToLocations.GetDest(a)
		lands = append(lands, a)
	}

	// fmt.Printf("%#v\n", seedToSoils)
	// fmt.Printf("%#v\n", soilToFertilizers)
	// fmt.Printf("%#v\n", waterToLights)
	// fmt.Printf("%#v\n", lightToTemperatures)
	// fmt.Printf("%#v\n", temperatureToHumidities)
	// fmt.Printf("%#v\n", humidityToLocations)
	sort.Sort(sort.IntSlice(lands))
	fmt.Printf("result ------> %#v\n", lands[0])
}

func extractNumbersPrefix(input string) ([]int, error) {
	// **: ** 部分を無視する
	_parts := strings.Split(input, ":")
	return extractNumbers(_parts[1])
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

func readLandMaps(scanner *bufio.Scanner) (re LandMapSlice) {
	for {
		scanner.Scan()
		line := scanner.Text()
		if line == "" {
			break
		}
		nums := strings.Split(line, " ")
		start, _ := strconv.Atoi(nums[0])
		SourceStart, _ := strconv.Atoi(nums[1])
		width, _ := strconv.Atoi(nums[2])
		re = append(re, LandMap{start, SourceStart, width})
	}
	return 
}