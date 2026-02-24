package main

import "sort"

// decidePlay 決定打出哪張手牌（含剛摸入的第 nHand 張）：
//  1. 打後聽牌數最多的牌（p.gates 須已由 gates() 填好）
//  2. 若無聽牌機會，打出「已出現最多張」的牌（最稀有的留著）
//  3. 最後備援：隨機打牌
func (s *Server) decidePlay(p *Player) int {
	// 策略一：打後聽牌最多
	key, value := -10, 0
	for k, v := range p.gates {
		if v > value {
			key, value = k, v
		}
	}
	for i, t := range p.hand[:p.nHand+1] {
		if t/4 == key/4 {
			return i
		}
	}

	// 策略二：打最常出現的牌
	key, value = -10, 4
	for k, v := range p.cPlay {
		if v <= value {
			key, value = k, v
		}
	}
	for i, t := range p.hand[:p.nHand+1] {
		if t/4 == key/4 {
			return i
		}
	}

	// 策略三：隨機
	return RandomPlay(p.nHand)
}

// gates 計算打出每張牌後的聽牌清單，同時填寫 p.cPlay（打牌機率輔助資料）。
// 回傳 map[牌面*4]聽牌數（候選牌以 count 表示可摸到的數量）。
func (s *Server) gates(p *Player) map[int]int {
	p.cPlay = map[int]int{}
	hist, candidate, g := [3*9 + 4 + 3]int{}, map[int]int{}, map[int]int{}

	// 統計手牌（含剛摸入的牌）各點數出現次數
	for _, t := range p.hand[:p.nHand+1] {
		if t/4 < 3*9+4+3 {
			hist[t/4]++
		}
	}

	// 把所有牌面有的每一個數牌都當作候選聽牌（累積前綴和後鋪平）
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
		if hist[i] > 0 && count < 4 { // 尚未出現所有 4 張
			candidate[i*4] += count // 出現越多的牌計算越大
		}
	}

	// 對每張手牌嘗試替換成候選牌，若能胡牌則記錄
	for i, t := range p.hand[:p.nHand+1] {
		p.cPlay[i] = 4 - p.see[t/4]
		for c := range candidate {
			p.hand[i] = c // 假設摸入 c
			if s.isWin(p) {
				g[(t/4)*4]++
			}
			p.hand[i] = t // 還原
		}
	}
	return g
}

// isWin 判斷玩家目前手牌（hand[0..nHand] 共 nHand+1 張）是否構成合法胡牌。
func (s *Server) isWin(p *Player) bool {
	sorted, pairs, uniquePairs := s.sortHand(p.hand), []int{}, map[int]int{}
	suited, honor := s.findPair(sorted, &pairs)

	for _, index := range pairs {
		uniquePairs[sorted[index]/4]++
	}

	for n := range uniquePairs {
		if n < 3*9 { // 數牌眼
			suited[n] -= 2
			if isHonor(honor) && s.isSuit(suited) {
				return true
			}
			suited[n] += 2
		} else { // 字牌眼
			honor[n-3*9] -= 2
			if isHonor(honor) && s.isSuit(suited) {
				return true
			}
			honor[n-3*9] += 2
		}
	}
	return false
}

// isHonor 驗證字牌部分：每種字牌（東南西北中發白）的剩餘數量必須為 0 或 3。
func isHonor(honor [4 + 3]int) bool {
	for _, count := range honor {
		if count > 0 && count != 3 {
			return false
		}
	}
	return true
}

// isSuit 驗證數牌部分是否可完整拆解成刻子與順子（貪婪遞迴，對應論文命題 4.2）。
// 定理保證：若最小點數牌數 ≥ 3，必可先拆成刻子再遞迴，結果正確。
func (s *Server) isSuit(suited [3 * 9]int) bool {
	count := 0
	for t := 0; t < 3; t++ { // 筒、條、萬
		for i := t * 9; i < t*9+9; i++ {
			n := suited[i]
			count += n
			if n < 0 {
				return false
			} else if n == 0 {
				continue
			} else if n >= 3 { // 先拆刻子
				suited[i] -= 3
				return s.isSuit(suited)
			} else if i > t*9+9-3 || suited[i+1] < 1 || suited[i+2] < 1 {
				return false // 無法拆順子
			} else { // 拆順子
				suited[i]--
				suited[i+1]--
				suited[i+2]--
				return s.isSuit(suited)
			}
		}
	}
	return count == 0
}

// sortHand 將手牌陣列（含摸入的第 nHand 張）複製並排序後回傳切片。
// 重新命名以避免與標準庫 sort 套件衝突。
func (s *Server) sortHand(tiles [17]int) []int {
	sl := make([]int, len(tiles))
	copy(sl, tiles[:])
	sort.Ints(sl)
	return sl
}

// findPair 在已排序手牌中找出所有候選眼的索引，同時統計數牌/字牌分布。
// 對應論文定理二（findSuitPair 用模 3 餘數快速定位眼所在的組）。
func (s *Server) findPair(sorted []int, pairs *[]int) (suited [3 * 9]int, honor [4 + 3]int) {
	i, j := 0, 0
	for ; i < len(sorted) && sorted[i] < 4*9; i++ {
		suited[sorted[i]/4]++
	} // 筒
	s.findSuitPair(0, sorted[:i], pairs)

	for j = i; i < len(sorted) && sorted[i] < 2*4*9; i++ {
		suited[sorted[i]/4]++
	} // 條
	s.findSuitPair(j, sorted[j:i], pairs)

	for j = i; i < len(sorted) && sorted[i] < 3*4*9; i++ {
		suited[sorted[i]/4]++
	} // 萬
	s.findSuitPair(j, sorted[j:i], pairs)

	if i >= len(sorted) || sorted[i]/4 >= 3+4+3*9 {
		return suited, honor
	}
	j = i
	for _, n := range sorted[j:] {
		honor[n/4-3*9]++
	} // 字牌
	s.findHonorPair(j, sorted[j:], pairs)
	return suited, honor
}

// findHonorPair 找字牌中的所有候選眼（直接逐一枚舉）。
func (s *Server) findHonorPair(pad int, tiles []int, pairs *[]int) {
	for i := range tiles {
		s.findiPair(pad, tiles, i, pairs)
	}
}

// findiPair 判斷 tiles[i] 是否為候選眼起始位置（相鄰兩張牌面相同且不與前一張重複）。
func (s *Server) findiPair(pad int, tiles []int, i int, pairs *[]int) {
	j := i + 1
	if j >= len(tiles) || tiles[i]/4 != tiles[j]/4 || (i > 0 && tiles[i-1]/4 == tiles[i]/4) {
		return
	}
	*pairs = append(*pairs, pad+i)
}

// findSuitPair 利用模 3 餘數定理快速定位數牌中眼所在的組（對應論文定理二）。
//
// 數學原理：
//   - 刻子（同點數 3 張）使某組 +3，三組模 3 餘數不變。
//   - 順子（連續 3 張）使三組各 +1，三組餘數同步改變，差值不變。
//   - 眼（同點數 2 張）使某組 +2，該組餘數與另兩組產生差異。
//
// 因此，三組中「餘數與另兩組不同的那組」必然包含眼。
func (s *Server) findSuitPair(pad int, tiles []int, pairs *[]int) {
	bin := [3]int{}
	indexBin := map[int][]int{}
	for i, t := range tiles {
		bin[(t/4)%3]++
		indexBin[(t/4)%3] = append(indexBin[(t/4)%3], i)
	}
	b := 2 // 預設：眼在第 2 組
	if bin[0]%3 != bin[1]%3 {
		if bin[0]%3 != bin[2]%3 {
			b = 0 // 眼在第 0 組
		} else {
			b = 1 // 眼在第 1 組
		}
	}
	for _, i := range indexBin[b] {
		s.findiPair(pad, tiles, i, pairs)
	}
}
