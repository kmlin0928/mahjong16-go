# 胡牌結算完整顯示（含花牌）與四家總覽

## 任務清單

- [v] Task 1：修正 main() 部分 return 遺漏 seat_winds（Bug）
  - Commit d667259：fix(terminal): 補齊 main() 中遺漏 seat_winds 的三處 return
  - 檔案：mahjong.py（main() 函式中 3 處 return）
  - 摘要：第 2457、2517、2644 行的 `return player/rob_idx/cand_idx, dealer_idx`
    補上第三項 seat_winds，與宣告的回傳型別一致

- [v] Task 2：Web 結算加入花牌顯示
  - Commit 8f36200：feat(ui): 遊戲結算加入四家花牌顯示
  - 檔案：static/app.js（showGameOver 手牌列渲染邏輯）
  - 摘要：在手牌後加 `state.bonus[i]` 花牌；
    若有花牌則先加分隔線（meld-sep）再渲染花牌 tile（.flower 樣式）

- [v] Task 3：終端機結局顯示四家手牌／面牌／花牌
  - Commit 72931d8：feat(terminal): 遊戲結局印出四家手牌、面牌、花牌摘要
  - 檔案：mahjong.py（main() 函式中每個 return 前加印結局摘要）
  - 摘要：新增 helper _print_player_summary(m, seat_winds) 迴圈印出：
    門風、手牌（已排序）、面牌（各組）、花牌；
    在每個 return 前呼叫（含和局 break 後）

## 狀態規則
[ ]:未開始　[/]:進行中　[v]:完成

## 驗收條件
- 執行終端局後印出四家摘要（手牌+面牌+花牌）
- Web gameover banner 每家顯示：面牌 │ 手牌 │ 花牌
- main() 回傳型別一致（不會 ValueError unpack）
