# 上家 div 反轉排列 + Playwright 驗證

## 任務清單

### Task 1：style.css — 上家區塊改為 column-reverse 排列
- [/] 檔案範圍：`static/style.css`（#zone-left）
- 摘要：
  仿照 #zone-right 的 flex-direction: column-reverse，
  讓棄牌列出現在視覺上方（靠近桌面中央），
  並調整 justify-content / padding 使牌列對齊正確

### Task 2：Playwright 視覺測試
- [ ] 檔案範圍：Playwright 測試腳本（暫時執行，不提交）
- 摘要：
  啟動網頁版，截圖確認上家/下家的棄牌列位置正確

## 驗收條件
- 上家棄牌列出現在區塊頂部（靠近中央）
- 下家棄牌列維持在底部（靠近中央）
- Playwright 截圖無明顯版面異常
