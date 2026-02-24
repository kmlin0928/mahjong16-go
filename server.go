package main

// Server 持有牌堆狀態，透過 random.go 的 Shuffle() 初始化全副牌序。
type Server struct {
	nHand  int
	remain []int
	sea    []int
}

// NewServer 建立新的遊戲伺服器，nHand 為每位玩家的初始手牌張數。
// 牌序由 random.go 的 Shuffle() 產生。
func NewServer(nHand int) *Server {
	return &Server{
		nHand:  nHand,
		remain: Shuffle(),
	}
}

// nToChinese 將牌號（0～151）轉為中文名稱。
// 編碼規則：每種牌面有 4 張，牌面 = n/4。
//
//	0–35   : 1～9 筒（餅）
//	36–71  : 1～9 條（索）
//	72–107 : 1～9 萬
//	108–123: 東南西北
//	124–135: 中發白
//	136–143: 春夏秋冬梅蘭竹菊
func (s *Server) nToChinese(n int) string {
	if n < 0 || n >= 3*9*4+4*4+3*4+8 {
		return "？"
	} else if n < 9*4 { // 餅
		return string(rune('1'+n/4)) + "筒" // 1~9筒
	} else if n < 2*9*4 { // 索
		return string(rune('1'-9+n/4)) + "條" // 1~9條
	} else if n < 3*9*4 { // 萬
		return string(rune('1'-2*9+n/4)) + "萬" // 1~9萬
	} else if n < 4*4+3*9*4 { // 風
		return []string{"東", "南", "西", "北"}[-3*9+n/4]
	} else if n < 3*4+4*4+3*9*4 { // 三元
		return []string{"中", "發", "白"}[-4-3*9+n/4]
	}
	// 花
	return []string{"春", "夏", "秋", "冬", "梅", "蘭", "竹", "菊"}[n-3*4-4*4-3*9*4]
}
