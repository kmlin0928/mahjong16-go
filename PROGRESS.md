# 任務：AI或人類玩家胡牌時要顯示手牌

## 任務清單

### 1. [v] 五處胡牌路徑均顯示手牌
- 檔案：`mahjong.py`
- AI 自摸：移除 `if not contest_mode:` 條件，改為無論模式均顯示手牌
- 天胡：移除 `if not contest_mode:` 條件
- 人類自摸：在「自摸胡 X」之後新增手牌列印
- 放槍胡：在 score_hand 呼叫前新增手牌列印（_cp.hand 含棄牌）
- 搶槓胡：在 score_hand 呼叫前新增手牌列印（rob_p.hand 含搶到的牌）

## 驗收條件
- `uv run mahjong.py` 執行後，每次任意玩家胡牌均顯示完整手牌
- 非胡牌情況 AI 手牌在 contest_mode 下仍隱藏

## 狀態更新規則
- `[ ]` 未開始
- `[/]` 進行中
- `[v]` 完成
