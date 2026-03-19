# PROGRESS.md

任務：擴充 EXTREMELY_DANGEROUS 判定 — 刻子三張牌與對子兩張牌亦為湊牌極度危險
目標檔案：mahjong.py
狀態說明：[ ] 未開始 | [/] 進行中 | [v] 完成

---

## 任務清單

### 1. [v] 新增 `find_hand_pungs` — 偵測手牌中的「刻子候選」（2 張以上同花同點）

- **檔案範圍**：`mahjong.py`（`find_hand_chows` 正下方）
- **摘要**：
  - 輸入 `suited: list[int]`（長度 27 的數牌計數）
  - 輸出各牌面種類索引中張數 ≥ 2 者（刻子候選），以 `list[tuple[int, int]]` 回傳
    （每個 tuple 為 `(kind_idx, count)`，count 為 2 或 3）
  - 刻子已成（3 張）→ 極度危險；刻子半成（2 張）→ 同樣是湊牌極度危險
- **驗收**：函式存在，可被 `__main__` 呼叫並通過新增測試
- `42ac900` feat(mahjong-py): 新增 find_hand_pungs 偵測刻子候選

### 2. [v] 新增 `find_hand_pairs` — 偵測手牌中的「對子候選」（2 張以上同花同點）

- **檔案範圍**：`mahjong.py`（honor 牌面也適用，輸入為全牌面 seen list 或 honor list）
- **摘要**：
  - 輸入為字牌計數 `honor: list[int]`（長度 HONOR_KINDS）
  - 輸出 honor 中張數 ≥ 2 的牌面種類索引，以 `list[int]` 回傳
  - 對子（含將牌候選）屬湊牌中，應標記為 EXTREMELY_DANGEROUS
- **驗收**：函式存在，可被 `__main__` 呼叫並通過新增測試
- `073ea3d` feat(mahjong-py): 新增 find_hand_pairs 偵測字牌對子候選

### 3. [ ] 更新 `DangerLevel.EXTREMELY_DANGEROUS` docstring，說明三種判定來源

- **檔案範圍**：`mahjong.py`（`DangerLevel` enum docstring）
- **摘要**：
  - 補充「刻子候選（find_hand_pungs）」與「對子候選（find_hand_pairs）」
  - 更新 README.md 第 5 級說明，列出三種來源
- **驗收**：`uv run mahjong.py` 執行無誤，README 第 5 列說明更新

### 4. [ ] `__main__` 驗收測試：find_hand_pungs 與 find_hand_pairs

- **檔案範圍**：`mahjong.py`（`__main__` 測試區塊末尾）
- **摘要**：
  - `find_hand_pungs`：純刻子手牌（回傳含 kind_idx）、混合手牌、空手牌
  - `find_hand_pairs`：有對子（回傳 kind_idx）、全孤張（空列表）、刻子不算對子
- **驗收**：所有 assert 通過，`uv run mahjong.py` 執行無誤

---

## 狀態更新規則
- [ ] 未開始
- [/] 進行中
- [v] 完成（完成後附 Commit 編號前7碼與訊息摘要）
