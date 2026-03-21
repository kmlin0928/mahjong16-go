# 首局莊家隨機化

## 狀態規則
[ ] 未開始  [/] 進行中  [v] 完成

## 任務清單

### Task 1：mahjong.py — 首局莊家改為隨機（_game_loop + main）
- [/] 檔案範圍：`mahjong.py`（_game_loop 約 line 1969；main() 約 line 2375）
- 摘要：
  當 dealer_idx_override 為 None 時，改為 `random.randrange(4)` 取代固定的 `0`；
  兩處均需修改（GameSession._game_loop 與 main() 函式）

## 驗收條件
- 多次 `uv run mahjong.py` 可看到莊家為不同玩家（非固定人類）
- GameSession 64 組合測試仍通過（測試均有傳入 dealer_idx_override，不受影響）
