# PROGRESS.md

任務：新增「極度安全」危險等級（第四張已棄出或已被吃/碰/槓的牌）
目標檔案：mahjong.py、README.md

狀態說明：[ ] 未開始 | [/] 進行中 | [v] 完成

---

## 任務清單

### 1. [/] 新增 EXTREMELY_SAFE 等級並更新分類邏輯與文件

- **檔案範圍**：`mahjong.py`（`DangerLevel` 列舉、`classify_danger()`、驗收測試）、`README.md`
- **摘要**：
  - `DangerLevel` 新增 `EXTREMELY_SAFE = 0`（比 VERY_SAFE 更安全）
  - 判定條件：數牌或字牌的全局棄牌次數 == 4（四張全出），或出現在 `chi_tiles`/`pon_tiles`/`kong_tiles`
  - 花牌不在此條件（花牌只有 1 張，無法「四張全棄」）
  - `classify_danger()` 在最開頭加入此判斷（優先於花牌判斷）
  - 在 `__main__` 驗收區塊補上 EXTREMELY_SAFE 的測試
  - README 危險等級表格新增第 0 級
- **驗收**：
  - 某牌面在全局棄牌出現 4 次 → 回傳 `EXTREMELY_SAFE`
  - 花牌最多只能 `VERY_SAFE`（無法達到 EXTREMELY_SAFE）
  - 新等級可正確比較：`EXTREMELY_SAFE < VERY_SAFE < SAFE`

---

## 狀態更新規則
- [ ] 未開始
- [/] 進行中
- [v] 完成（完成後附 Commit 編號前7碼與訊息摘要）
