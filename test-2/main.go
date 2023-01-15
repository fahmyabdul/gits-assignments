package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	for {
		denseRankApp := NewDenseRankApp(bufio.NewScanner(os.Stdin))
		denseRankApp.Start()
	}
}

type DenseRankApp struct {
	inputScanner *bufio.Scanner
	playerTotal  int
	playersScore map[string]int
	gameTotal    int
	gamesScore   []int
}

func NewDenseRankApp(sc *bufio.Scanner) *DenseRankApp {
	return &DenseRankApp{inputScanner: sc}
}

func (p *DenseRankApp) Start() {
	fmt.Println(strings.Repeat("-", 30))
	fmt.Println("Dense Ranking Application")
	fmt.Println(strings.Repeat("-", 30))
	p.choice()

	fmt.Println("Set the number of players:")
	p.playerTotal = p.inputInt()

	fmt.Println("Set players score:")
	p.inputPlayersScore()

	fmt.Println("Set the total of game you want to play:")
	p.gameTotal = p.inputInt()

	fmt.Println("Set your games score:")
	p.inputGamesScore()

	fmt.Println(strings.Repeat("-", 30))
	fmt.Println("Your Rank")
	fmt.Println(strings.Repeat("-", 30))

	output := ""
	for k, v := range p.gamesScore {
		_, myRank := p.getRank(v)
		fmt.Printf("game %d, rank: %d, score: %d\n", k+1, myRank, v)
		output += fmt.Sprintf(" %d", myRank)
	}
	fmt.Println(strings.Repeat("-", 30))
	fmt.Println("Output:", strings.TrimSpace(output))
}

func (p *DenseRankApp) choice() {
	fmt.Println("1) Start")
	fmt.Println("0) Exit")

inputChoice:
	fmt.Printf("-> ")
	p.inputScanner.Scan()
	if choice := strings.TrimSpace(p.inputScanner.Text()); choice == "0" {
		fmt.Println("App terminated")
		os.Exit(0)
	} else if choice != "1" {
		fmt.Println("Unrecognized type")
		goto inputChoice
	}
}

func (p *DenseRankApp) inputInt() int {
inputInt:
	fmt.Printf("-> ")
	p.inputScanner.Scan()
	output, err := strconv.Atoi(strings.TrimSpace(p.inputScanner.Text()))
	if err != nil {
		fmt.Println("input must be an integer")
		goto inputInt
	}

	return output
}

func (p *DenseRankApp) inputPlayersScore() {
	p.playersScore = make(map[string]int, p.playerTotal)

	for i := 1; i <= p.playerTotal; i++ {
	inputPlayerScores:
		nKey := fmt.Sprintf("p_%d", i)
		fmt.Printf("-> %s: ", nKey)
		p.inputScanner.Scan()
		playerScore, err := strconv.Atoi(strings.TrimSpace(p.inputScanner.Text()))
		if err != nil {
			fmt.Println("input must be an integer")
			goto inputPlayerScores
		}
		p.playersScore[nKey] = playerScore
	}
}

func (p *DenseRankApp) inputGamesScore() []int {
	for i := 1; i <= p.gameTotal; i++ {
	inputGamesScore:
		nKey := fmt.Sprintf("game %d", i)
		fmt.Printf("-> %s: ", nKey)
		p.inputScanner.Scan()
		gameScore, err := strconv.Atoi(strings.TrimSpace(p.inputScanner.Text()))
		if err != nil {
			fmt.Println("input must be an integer")
			goto inputGamesScore
		}
		p.gamesScore = append(p.gamesScore, gameScore)
	}

	return p.gamesScore
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p *DenseRankApp) getRank(score int) (PairList, int) {
	var myRank int

	p.playersScore["new"] = score

	pl := make(PairList, len(p.playersScore))
	i := 0
	for k, v := range p.playersScore {
		pl[i] = Pair{Key: k, Value: v}
		i++
	}
	sort.Sort(sort.Reverse(pl))

	lastRank := 0
	lastScore := 0
	for _, v := range pl {
		if v.Value != lastScore {
			lastRank++
		}
		lastScore = v.Value

		if v.Key == "new" {
			myRank = lastRank
		}
	}

	return pl, myRank
}
