# 網頁版 log 顯示：補花、棄牌、吃/碰/槓

## 狀態規則
[ ] 未開始  [/] 進行中  [v] 完成

## 任務清單

### Task 1：修正 `_game_loop` log 清除時機（mahjong.py）
- [/] 檔案範圍：`mahjong.py`（`_game_loop` 函式）
- 摘要：
  1. 初始補花改用 `self._L()` 直接輸出（line 1986），移除 `_init_bonus_logs` 機制
  2. 移除 while 迴圈頂端的 `self._log_clear()`（line 2009）與 re-injection block（lines 2010-2013）
  3. 在所有 8 個 `yield` 回應後加入 `self._log_clear()`，確保兩次人類決策之間的所有事件累積後才清空

### Task 2：前端棄牌訊息強調顯示（app.js）
- [ ] 檔案範圍：`static/app.js`（`_MELD_KEYWORDS`）
- 摘要：在關鍵字列表加入 `'打'`，使「東打 X」「南打 X」等棄牌訊息套用 `meld-action` 高亮樣式

## 驗收條件
- 手測：玩幾輪後，中間 log 應依序顯示如：「東補花 春」、「你打 1筒」、「南碰 1筒」、「南打 7萬」等
- 不再只看到提示卡（胡/碰/吃），而是能看到所有本回合發生的事件
