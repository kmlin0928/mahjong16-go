# 網頁版 log-box 行距與捲動修正

## 狀態規則
[ ] 未開始  [/] 進行中  [v] 完成

## 任務清單

### Task 1：style.css — log-box 行距與排列方向
- [v] 檔案範圍：`static/style.css`
- Commit afdcc7e：fix(style): log-box 行距歸零並改為正向排列
- 摘要：
  1. `#log-box p { margin: 1px 0 }` → `margin: 0`
  2. `#log-box` 移除 `flex-direction: column-reverse`（改為預設 column）

### Task 2：app.js — appendLog 改用 appendChild + 自動捲到底
- [/] 檔案範圍：`static/app.js`
- 摘要：
  1. `box.prepend(p)` → `box.appendChild(p)`
  2. 加上 `box.scrollTop = box.scrollHeight` 確保最新一條完整可見

## 驗收條件
- log 條目之間無額外行距（margin 0）
- 每次新訊息進入後，捲動自動到最底，最新一條文字完整顯示不被截斷
