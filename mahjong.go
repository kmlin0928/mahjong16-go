package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Mahjong struct {
	nHand   int
	players [4]Player
	remain  []int
	sea     []int
}

type Player struct {
	nHand int
	hand  [17]int
	table []int
	see   [3*9 + 4 + 3]int // 數牌與字牌剩下的數量
	cPlay map[int]int      // 打出手牌的機率
	gates map[int]int      // 聽哪些牌
}

func (m *Mahjong) nToChinese(n int) (s string) {
	if n < 0 || n >= 3*9*4+4*4+3*4+8 {
		return "？"
	} else if n < 9*4 { // 餅
		return string(rune('1'+n/4)) + "筒" // 1~9筒 "d" as Dots
	} else if n < 2*9*4 { // 索
		return string(rune('1'-9+n/4)) + "條" // 1~9條 "b" as Bamboos
	} else if n < 3*9*4 { // 萬
		return string(rune('1'-2*9+n/4)) + "萬" // 1~9萬 "c" as Characters
	} else if n < 4*4+3*9*4 { // 風
		return []string{"東", "南", "西", "北"}[-3*9+n/4] // 東南西北 "ewsn"
	} else if n < 3*4+4*4+3*9*4 { // 龍
		return []string{"中", "發", "白"}[-4-3*9+n/4] // 中發白 "*$@"
	} else { // 花 // 春夏秋冬 "uaio", 梅蘭竹菊 "mljz"
		return []string{"春", "夏", "秋", "冬", "梅", "蘭", "竹", "菊"}[n-3*4-4*4-3*9*4]
	}
}

func (m *Mahjong) deal1() (n int) {
	tile := -1
	if len(m.remain) > 0 { // 有牌可發
		tile, m.remain = m.remain[0], m.remain[1:]
	}
	return tile
}

func (m *Mahjong) initDeal() {
	m.remain = rand.Perm(3*9*4 + 4*4 + 3*4 + 8)
	for i := 0; i < m.nHand; i++ {
		for j := 0; j < 4; j++ {
			m.players[j].hand[i] = m.deal1()
		}
	}
	for i := 0; i < 4; i++ {
		m.players[i].nHand = m.nHand
	}
}

func (m *Mahjong) showBonus() { // 補花
	for player := 0; player < 4; player++ {
		p := &(m.players[player])
		fmt.Printf("\n%d", player)
		for i := 0; i < p.nHand; i++ {
			fmt.Printf(" %s", m.nToChinese(p.hand[i]))
			m.iShowBonus(p, i)
		}
		m.initSee(p) // 可摸的牌不含手上的牌
	}
}

func (m *Mahjong) iShowBonus(p *Player, i int) {
	n := p.hand[i]
	for n >= 3*4+4*4+3*9*4 { // 花
		p.table = append(p.table, n)
		if n = m.deal1(); n < 0 {
			return
		}
		fmt.Printf("補 %s", m.nToChinese(n))
	}
	p.hand[i] = n
}

func (m *Mahjong) initSee(p *Player) {
	for _, t := range p.hand[:p.nHand] { // 不含將打出去的牌
		p.addSee(t) // 可摸的牌不含手上的牌
	}
}

func (p *Player) addSee(t int) {
	p.see[t/4]++ // 可摸的牌不含剛打出的牌
}

func (m *Mahjong) decidePlay(p *Player) (n int) {
	// 打後聽最多牌的牌
	key, value := -10, 0
	for k, v := range p.gates {
		if v > value {
			key, value = k, v
		}
	}
	for i, t := range p.hand[:p.nHand+1] {
		if t/4 == key/4 {
			return i // 選打後聽最多的牌
		}
	}

	// 打最常出現的牌
	key, value = -10, 4
	for k, v := range p.cPlay {
		if v <= value {
			key, value = k, v
		}
	}
	for i, t := range p.hand[:p.nHand+1] {
		if t/4 == key/4 {
			return i // 選最常出現的牌
		}
	}
	return rand.Intn(p.nHand + 1) // 隨機選一張牌
}

func (p *Player) play(n int, hand int) {
	p.hand[n], p.hand[hand] = p.hand[hand], p.hand[n] // 打出第n張牌, 換入摸到的牌
}

func (m *Mahjong) sort(tiles [17]int) (s []int) {
	s = append(tiles[:len(tiles)-1], tiles[len(tiles)-1])
	sort.Ints(s)
	return s
}

func (m *Mahjong) gates(p *Player) map[int]int { // 聽哪些牌
	p.cPlay = map[int]int{} //  沒出過幾張牌
	hist, candidate, g := [3*9 + 4 + 3]int{}, map[int]int{}, map[int]int{}
	for _, t := range p.hand[:p.nHand+1] {
		if t/4 < 3*9+4+3 { // 有牌可補花
			hist[t/4]++
		}
	}

	// 把所有牌面有的每一個數牌都當作候選聽牌
	for i := 1; i < 9; i++ {
		hist[i] += hist[i-1]
		hist[i+9] += hist[i+9-1]
		hist[i+2*9] += hist[i+2*9-1]
	}
	for i := 0; i < 9-1; i++ {
		hist[i] = hist[9-1]
		hist[i+9] += hist[2*9-1]
		hist[i+2*9] += hist[3*9-1]
	}
	for i, count := range p.see {
		if hist[i] > 0 && count < 4 { // 尚未出現所有4張牌
			candidate[i*4] += count // 越出過的牌計算越大
		}
	}

	for i, t := range p.hand[:p.nHand+1] {
		p.cPlay[i] = 4 - p.see[t/4]
		for c := range candidate {
			p.hand[i] = c // 假設摸入第i張牌
			if m.isWin(p) {
				g[(t/4)*4]++
			}
			p.hand[i] = t // 還原第i張牌
		}
	}
	return g
}

func (m *Mahjong) isWin(p *Player) (win bool) { //是否胡牌
	sortedHand, pairs, uniquePairs := m.sort(p.hand), []int{}, map[int]int{}
	suited, honor := m.findPair(sortedHand, &pairs) // 找到眼及牌點與字的分佈

	for _, index := range pairs {
		uniquePairs[sortedHand[index]/4]++
	}

	for n := range uniquePairs {
		if n < 3*9 { // 數牌
			suited[n] -= 2
			if isHonor(honor) && m.isSuit(suited) { // 去掉眼和字牌，找數牌胡牌型
				return true
			}
			suited[n] += 2
		} else { // 字牌
			honor[n-3*9] -= 2
			if isHonor(honor) && m.isSuit(suited) { // 去掉眼和字牌，找數牌胡牌型
				return true
			}
			honor[n-3*9] += 2
		}
	}
	return false
}

func isHonor(honor [4 + 3]int) bool {
	for _, count := range honor {
		if count > 0 && count != 3 {
			return false
		}
	}
	return true
}

func (m *Mahjong) isSuit(suited [3 * 9]int) bool { // 定理一 https://www.thenewslens.com/article/100657
	count := 0
	for t := 0; t < 3; t++ { // 筒, 條, 萬
		for i := t * 9; i < t*9+9; i++ {
			n := suited[i]
			count += n
			if n < 0 {
				return false
			} else if n == 0 {
				continue
			} else if n >= 3 {
				suited[i] -= 3
				return m.isSuit(suited)
			} else if i > t*9+9-3 || suited[i+1] < 1 || suited[i+2] < 1 {
				return false
			} else {
				suited[i]--
				suited[i+1]--
				suited[i+2]--
				return m.isSuit(suited)
			}
		}
	}
	if count == 0 {
		return true
	}
	return false
}

func (m *Mahjong) findPair(s []int, pairs *[]int) (suited [3 * 9]int, honor [4 + 3]int) { // 定理二 https://www.thenewslens.com/article/100657
	i, j := 0, 0
	for ; i < len(s) && s[i] < 4*9; i++ {
		suited[s[i]/4]++
	} // 筒
	m.findSuitPair(0, s[:i], pairs)
	for j = i; i < len(s) && s[i] < 2*4*9; i++ {
		suited[s[i]/4]++
	} // 條
	m.findSuitPair(j, s[j:i], pairs)
	for j = i; i < len(s) && s[i] < 3*4*9; i++ {
		suited[s[i]/4]++
	} // 萬
	m.findSuitPair(j, s[j:i], pairs)

	if i >= len(s) || s[i]/4 >= 3+4+3*9 {
		return suited, honor
	}
	j = i
	for _, n := range s[j:] {
		honor[n/4-3*9]++
	} // 字
	m.findHonorPair(j, s[j:], pairs)
	return suited, honor
}

func (m *Mahjong) findHonorPair(pad int, s []int, pairs *[]int) { // 字的眼
	for i := range s {
		m.findiPair(pad, s, i, pairs)
	}
}

func (m *Mahjong) findiPair(pad int, s []int, i int, pairs *[]int) {
	j := i + 1
	if j >= len(s) || s[i]/4 != s[j]/4 || (i > 0 && s[i-1]/4 == s[i]/4) {
		return
	} // 每張牌有4個複製, 只取第一次看到的複製
	// fmt.Printf(" 眼%s", m.nToChinese(s[i]))
	*pairs = append(*pairs, pad+i)
}

func (m *Mahjong) findSuitPair(pad int, s []int, pairs *[]int) { // 順的眼 // 定理二 https://www.thenewslens.com/article/100657
	// filterMin := t * 4 * 9 // filterMax := filterMin + 4*9
	bin := [3]int{0, 0, 0}
	indexBin := map[int][]int{}
	for i, t := range s {
		bin[(t/4)%3]++
		indexBin[(t/4)%3] = append(indexBin[(t/4)%3], i)
	}
	b := 2 // 不在第0堆也不在第1堆, 那必在第2堆
	if bin[0]%3 != bin[1]%3 {
		if bin[0]%3 != bin[2]%3 { // 眼在第0堆
			b = 0
		} else { // 眼在第1堆
			b = 1
		}
	}
	for _, i := range indexBin[b] {
		m.findiPair(pad, s, i, pairs)
	}
}

func main() {
	m := Mahjong{}
	m.nHand = 16 // 十六張麻將
	m.initDeal()
	m.showBonus()

	fmt.Println()

	for player := 0; len(m.remain) > 0; player = (player + 1) % 4 {
		p := &m.players[player]
		p.hand[p.nHand] = m.deal1()
		fmt.Printf("\n%d摸 %s", player, m.nToChinese(p.hand[p.nHand]))
		m.iShowBonus(p, p.nHand)
		// p.hand = [17]int{44, 0, 1, 2, 4, 8, 12, 16, 20, 24, 28, 32, 33, 34, 35, 41, 40}
		pTile := p.hand[p.nHand]
		p.addSee(pTile) // 記錄摸到的牌

		p.gates = m.gates(p)
		fmt.Printf(" 打後聽牌:")
		for gate, chance := range p.gates {
			fmt.Printf(" %s=%d", m.nToChinese(gate), chance)
		}
		if len(m.remain) <= 0 {
			fmt.Printf("\n和局")
			break
		} else if m.isWin(p) {
			fmt.Printf("\n%d胡", player)
			for _, t := range p.hand[:p.nHand] {
				fmt.Printf(" %s", m.nToChinese(t))
			}
			break
		}
		p.play(m.decidePlay(p), p.nHand) // 將打出的牌與摸到的牌交換
		pTile = p.hand[p.nHand]
		fmt.Printf("\n%d打 %s_", player, m.nToChinese(pTile))

		for _, t := range p.hand[:p.nHand] {
			fmt.Printf(" %s", m.nToChinese(t))
		}
		for _, t := range p.table {
			fmt.Printf("|%s", m.nToChinese(t))
		}

		m.sea = append(m.sea, pTile)         // 海底加上打出的牌
		for other := 1; other < 4; other++ { // 其他三家記錄打出的牌
			(&m).players[(player+other)%4].addSee(pTile)
		}
		p.hand[p.nHand] = -1 // 打出的牌移出玩家
	}
}
