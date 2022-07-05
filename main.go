package main

import (
	"fmt"
	"math/rand"
	"time"
)

// カード構造体
type Card struct{
	Suit string		// スート（ハート、ダイヤ、クラブ、スペード）
	Rank string		// A,2,3,4,5,6,7,8,9,10,J,Q,K
	Value int
}

// 初期カード52枚
var Cards = [52]Card{
	{Suit:"ハート", Rank: "A", Value: 1},
	{Suit:"ハート", Rank: "2", Value: 2},
	{Suit:"ハート", Rank: "3", Value: 3},
	{Suit:"ハート", Rank: "4", Value: 4},
	{Suit:"ハート", Rank: "5", Value: 5},
	{Suit:"ハート", Rank: "6", Value: 6},
	{Suit:"ハート", Rank: "7", Value: 7},
	{Suit:"ハート", Rank: "8", Value: 8},
	{Suit:"ハート", Rank: "9", Value: 9},
	{Suit:"ハート", Rank: "10", Value: 10},
	{Suit:"ハート", Rank: "J", Value: 10},
	{Suit:"ハート", Rank: "Q", Value: 10},
	{Suit:"ハート", Rank: "K", Value: 10},

	{Suit:"ダイヤ", Rank: "A", Value: 1},
	{Suit:"ダイヤ", Rank: "2", Value: 2},
	{Suit:"ダイヤ", Rank: "3", Value: 3},
	{Suit:"ダイヤ", Rank: "4", Value: 4},
	{Suit:"ダイヤ", Rank: "5", Value: 5},
	{Suit:"ダイヤ", Rank: "6", Value: 6},
	{Suit:"ダイヤ", Rank: "7", Value: 7},
	{Suit:"ダイヤ", Rank: "8", Value: 8},
	{Suit:"ダイヤ", Rank: "9", Value: 9},
	{Suit:"ダイヤ", Rank: "10", Value: 10},
	{Suit:"ダイヤ", Rank: "J", Value: 10},
	{Suit:"ダイヤ", Rank: "Q", Value: 10},
	{Suit:"ダイヤ", Rank: "K", Value: 10},

	{Suit:"クラブ", Rank: "A", Value: 1},
	{Suit:"クラブ", Rank: "2", Value: 2},
	{Suit:"クラブ", Rank: "3", Value: 3},
	{Suit:"クラブ", Rank: "4", Value: 4},
	{Suit:"クラブ", Rank: "5", Value: 5},
	{Suit:"クラブ", Rank: "6", Value: 6},
	{Suit:"クラブ", Rank: "7", Value: 7},
	{Suit:"クラブ", Rank: "8", Value: 8},
	{Suit:"クラブ", Rank: "9", Value: 9},
	{Suit:"クラブ", Rank: "10", Value: 10},
	{Suit:"クラブ", Rank: "J", Value: 10},
	{Suit:"クラブ", Rank: "Q", Value: 10},
	{Suit:"クラブ", Rank: "K", Value: 10},

	{Suit:"スペード", Rank: "A", Value: 1},
	{Suit:"スペード", Rank: "2", Value: 2},
	{Suit:"スペード", Rank: "3", Value: 3},
	{Suit:"スペード", Rank: "4", Value: 4},
	{Suit:"スペード", Rank: "5", Value: 5},
	{Suit:"スペード", Rank: "6", Value: 6},
	{Suit:"スペード", Rank: "7", Value: 7},
	{Suit:"スペード", Rank: "8", Value: 8},
	{Suit:"スペード", Rank: "9", Value: 9},
	{Suit:"スペード", Rank: "10", Value: 10},
	{Suit:"スペード", Rank: "J", Value: 10},
	{Suit:"スペード", Rank: "Q", Value: 10},
	{Suit:"スペード", Rank: "K", Value: 10},
}

type ScoreInfo struct{
	Score int
	HasCards []Card
}

func main(){
	player := ScoreInfo{Score: 0}
	dealer := ScoreInfo{}

	fmt.Println("ブラックジャックへようこそ！！")
	fmt.Println("ゲームを開始します")
	fmt.Println("------------------")

	rand.Seed(time.Now().Unix())
	playerResult := rand.Intn(52)
	fmt.Printf("あなたの引いたカードは%sの%vです\n", Cards[playerResult].Suit, Cards[playerResult].Rank)
	player.HasCards = append(player.HasCards, Cards[playerResult])
	player.Score += Cards[playerResult].Value

	playerResult = rand.Intn(52)
	fmt.Printf("あなたの引いたカードは%sの%vです\n", Cards[playerResult].Suit, Cards[playerResult].Rank)
	player.HasCards = append(player.HasCards, Cards[playerResult])
	player.Score += Cards[playerResult].Value

	fmt.Println("------------------")

	dealerResult := rand.Intn(52)
	fmt.Printf("ディーラーの引いたカードは%sの%vです\n", Cards[dealerResult].Suit, Cards[dealerResult].Rank)
	dealer.HasCards = append(dealer.HasCards, Cards[dealerResult])
	dealer.Score += Cards[dealerResult].Value

	dealerResult = rand.Intn(52)
	fmt.Printf("ディーラーの引いたカードは%sの%vです\n", Cards[dealerResult].Suit, Cards[dealerResult].Rank)
	dealer.HasCards = append(dealer.HasCards, Cards[dealerResult])
	dealer.Score += Cards[dealerResult].Value

	fmt.Println("------------------")

	fmt.Printf("あなたの現在の得点は%v点です\n", player.Score)
	fmt.Printf("カードを引きますか？")
	fmt.Printf("引く場合はYを、引かない場合はNを押してください\n")

	
}