package main

import (
	"fmt"
	"math/rand"
)

// Shuffle 回傳全部 144 張牌（3×9×4 數牌 + 4×4 風牌 + 3×4 三元牌 + 8 花牌）的隨機排列。
// 供 server.go 的 NewServer() 呼叫以建立初始牌堆。
func Shuffle() []int {
	return rand.Perm(3*9*4 + 4*4 + 3*4 + 8)
}

// deal1 從牌堆頂抽取 1 張牌；若牌堆已空則回傳 -1。
func (s *Server) deal1() int {
	if len(s.remain) == 0 {
		return -1
	}
	tile := s.remain[0]
	s.remain = s.remain[1:]
	return tile
}

// initDeal 依序發初始手牌：先給玩家 0 發第 1 張，再給玩家 1，…，直到每人都有 nHand 張。
func (s *Server) initDeal(players [4]*Player) {
	for i := 0; i < s.nHand; i++ {
		for j := 0; j < 4; j++ {
			players[j].hand[i] = s.deal1()
		}
	}
	for i := 0; i < 4; i++ {
		players[i].nHand = s.nHand
	}
}

// showBonus 開局補花：印出每位玩家手牌並對花牌遞迴補牌，最後初始化已見牌記錄。
func (s *Server) showBonus(players [4]*Player) {
	for i, p := range players {
		fmt.Printf("\n%d", i)
		for j := 0; j < p.nHand; j++ {
			fmt.Printf(" %s", s.nToChinese(p.hand[j]))
			s.iShowBonus(p, j)
		}
		p.initSee() // 可摸的牌不含手上的牌
	}
}

// iShowBonus 對手牌第 i 張遞迴補花：若為花牌則放入桌面並重新抽牌。
func (s *Server) iShowBonus(p *Player, i int) {
	n := p.hand[i]
	for n >= 3*4+4*4+3*9*4 { // 花牌範圍
		p.table = append(p.table, n)
		if n = s.deal1(); n < 0 {
			return
		}
		fmt.Printf("補 %s", s.nToChinese(n))
	}
	p.hand[i] = n
}

// RandomPlay 回傳 0 至 nHand（含）之間的隨機整數，作為 AI 的隨機後備打牌索引。
func RandomPlay(nHand int) int {
	return rand.Intn(nHand + 1)
}
