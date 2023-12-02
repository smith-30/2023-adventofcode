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

	// 行ごとにスキャン
	for scanner.Scan() {
		line := scanner.Text()
		game := parseGameData(line)
		limits := map[string]int{
			"red": 1,
			"green": 1,
			"blue": 1,
		}
		for _, item := range game.Plays {
			for _, play := range item {
				limit := limits[play.Color]
				if play.Count > limit {
					limits[play.Color] = play.Count
				}
			}
		}
		result += (limits["red"] * limits["green"] * limits["blue"])
	}

	fmt.Printf("result ------> %#v\n", result)
}

type Game struct {
	ID int
	Plays [][]Play
}

type Play struct {
	Color string // red, green, blue
	Count int
}

func parseGameData(data string) Game {
	// 正規表現でGame IDと各プレイのデータを抽出
	re := regexp.MustCompile(`Game (\d+): (.+)`)
	matches := re.FindStringSubmatch(data)

	// Game IDを抽出
	gameID := matches[1]
	// 文字列から数値に変換
	gameIDInt, _ := strconv.Atoi(gameID)


	// 各プレイのデータを抽出
	playData := matches[2]
	playsData := strings.Split(playData, ";")
	ps := [][]Play{}
	for _, item := range playsData {
		_ps := []Play{}
		// 空白を除去
		item = strings.TrimSpace(item)

		// 空行を無視
		if item == "" {
			continue
		}

		cplays := strings.Split(item, ",")
		for _, citem := range cplays {
			p := parsePlayData(citem)
			_ps = append(_ps, p)
		}
		ps = append(ps, _ps)
	}
	// Game構造体にデータを格納
	game := Game{
		ID:    gameIDInt,
		Plays: ps,
	}

	return game
}

func parsePlaysData(data string) []Play {
	// セミコロンでデータを分割
	playData := strings.Split(data, ";")

	// 空のPlayスライスを作成
	var plays []Play

	// 各プレイのデータを処理
	for _, playStr := range playData {
		// 空白を除去
		playStr = strings.TrimSpace(playStr)

		// 空行を無視
		if playStr == "" {
			continue
		}

		// Play構造体にデータをマッピング
		play := parsePlayData(playStr)

		// マッピングしたPlayをスライスに追加
		plays = append(plays, play)
	}

	return plays
}

func parsePlayData(data string) Play {
	// カンマでデータを分割
	playParts := strings.Split(data, ",")

	// 空のPlay構造体を作成
	var play Play

	// 各プレイのデータを処理
	for _, part := range playParts {
		// 空白を除去
		part = strings.TrimSpace(part)

		// 数字と色を抽出
		re := regexp.MustCompile(`(\d+) (\w+)`)
		matches := re.FindStringSubmatch(part)
		if len(matches) == 3 {
			count, _ := strconv.Atoi(matches[1])
			color := matches[2]

			// Play構造体にデータを格納
			play = Play{
				Color: color,
				Count: count,
			}
		}
	}

	return play
}