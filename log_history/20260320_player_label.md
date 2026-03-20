# 任務：玩家編號改為相對稱謂（下家／對家／上家）

## 說明
新增輔助函式 `player_label(player)` 回傳相對於 `HUMAN_PLAYER=0` 的稱謂：
- 自己 → `你`
- offset 1 → `下家`
- offset 2 → `對家`
- offset 3 → `上家`

## 任務清單

### 1. [v] 新增 player_label() 輔助函式
- 檔案：`mahjong.py`（`main()` 正上方）
- 摘要：新增 `player_label(player: int) -> str`，以 `HUMAN_PLAYER` 為基準回傳 你/下家/對家/上家

### 2. [v] 替換遊戲迴圈中所有玩家編號顯示
- 檔案：`mahjong.py`（`main()` 內）
- 摘要：
  - `{player}摸` / `{player}胡` / `{player}自摸胡` / `{player}天胡` / `{player}打` / `{player}加槓` / `{player}出牌`
  - `{cand_idx}胡！` / `{rob_idx}搶槓胡！` / `{cand_idx}碰` / `{next_idx}吃`
  - 棄牌列表的 `P{opp}(你)` 改為 `player_label(opp)`
  - 移除多餘的 `{you}` / `{you_mark}` 後綴（`player_label` 已涵蓋）

## 驗收條件
- 執行 `uv run mahjong.py` 後，出牌／摸牌／胡牌訊息顯示「下家」「對家」「上家」「你」，不再出現數字 1 2 3
- 人類玩家顯示「你」

## 狀態更新規則
- `[ ]` 未開始
- `[/]` 進行中
- `[v]` 完成
