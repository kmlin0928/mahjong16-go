package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Player 代表一位玩家的所有狀態。
type Player struct {
	nHand int               // 初始手牌張數（固定 16）
	hand  [17]int           // hand[0..nHand-1] 手牌；hand[nHand] 為剛摸入的牌
	table []int             // 桌面牌（已補花）
	see   [3*9 + 4 + 3]int // 各牌面已見過的數量（供 AI 估算剩餘牌數）
	cPlay map[int]int       // 手牌各位置的「打牌機率」輔助資料（由 gates() 填寫）
	gates map[int]int       // 打出某牌後可等待的聽牌清單（由 gates() 填寫）
}

// initSee 初始化已見牌記錄，排除目前手上的牌。
func (p *Player) initSee() {
	for _, t := range p.hand[:p.nHand] {
		p.addSee(t)
	}
}

// addSee 將牌 t 標記為已見（自己手牌、他人打出的牌均呼叫此函式）。
func (p *Player) addSee(t int) {
	p.see[t/4]++
}

// play 打出手牌第 n 張：將其與剛摸入的第 hand 張互換位置，
// 使打出的牌移到 hand[nHand]（供呼叫端取出並加入海底）。
func (p *Player) play(n int, hand int) {
	p.hand[n], p.hand[hand] = p.hand[hand], p.hand[n]
}

// playAI 呼叫 AI.go 的決策函式自動選擇打出哪張牌，回傳手牌索引。
// 注意：p.gates 須已由外部（通常是 main 的遊戲迴圈）呼叫 s.gates(p) 填好，
// 此函式直接使用已計算好的結果，不重複計算。
func (p *Player) playAI(s *Server) int {
	return s.decidePlay(p)
}

// playManual 顯示目前手牌（含剛摸入的牌）並由使用者輸入要打出的牌索引（0-based）。
// 輸入不合法時會重複提示。
func (p *Player) playManual(s *Server) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("\n選擇打出的牌（")
		for i, t := range p.hand[:p.nHand+1] {
			fmt.Printf("%d:%s ", i, s.nToChinese(t))
		}
		fmt.Printf("）：")
		line, _ := reader.ReadString('\n')
		n, err := strconv.Atoi(strings.TrimSpace(line))
		if err == nil && n >= 0 && n <= p.nHand {
			return n
		}
		fmt.Println("輸入無效，請重新輸入。")
	}
}

// main 執行四人麻將自動對局（測試範例：全部玩家呼叫 AI 自動打牌）。
//
// 若要改為手動操作某位玩家，可將對應的
//
//	n := p.playAI(s)
//
// 替換為：
//
//	n := p.playManual(s)
func main() {
	s := NewServer(16) // 十六張麻將
	players := [4]*Player{}
	for i := range players {
		players[i] = &Player{}
	}

	s.initDeal(players)   // 每人發 16 張初始手牌
	s.showBonus(players)  // 開局補花並初始化已見牌記錄
	fmt.Println()

	for player := 0; len(s.remain) > 0; player = (player + 1) % 4 {
		p := players[player]

		// ── 摸牌 ──
		p.hand[p.nHand] = s.deal1()
		fmt.Printf("\n%d摸 %s", player, s.nToChinese(p.hand[p.nHand]))
		s.iShowBonus(p, p.nHand) // 摸到花牌則補牌
		p.addSee(p.hand[p.nHand])

		// ── 計算聽牌並顯示 ──
		p.gates = s.gates(p)
		fmt.Printf(" 打後聽牌:")
		for gate, chance := range p.gates {
			fmt.Printf(" %s=%d", s.nToChinese(gate), chance)
		}

		// ── 胡牌 / 和局判定 ──
		if len(s.remain) <= 0 {
			fmt.Printf("\n和局")
			break
		} else if s.isWin(p) {
			fmt.Printf("\n%d胡", player)
			for _, t := range p.hand[:p.nHand] {
				fmt.Printf(" %s", s.nToChinese(t))
			}
			break
		}

		// ── 打牌（全部玩家使用 AI 自動打牌）──
		n := p.playAI(s)
		p.play(n, p.nHand)
		pTile := p.hand[p.nHand] // 打出的牌已換至 hand[nHand]

		fmt.Printf("\n%d打 %s_", player, s.nToChinese(pTile))
		for _, t := range p.hand[:p.nHand] {
			fmt.Printf(" %s", s.nToChinese(t))
		}
		for _, t := range p.table {
			fmt.Printf("|%s", s.nToChinese(t))
		}

		// ── 更新牌局狀態 ──
		s.sea = append(s.sea, pTile)
		for other := 1; other < 4; other++ {
			players[(player+other)%4].addSee(pTile)
		}
		p.hand[p.nHand] = -1 // 打出的牌移出玩家
	}
}
