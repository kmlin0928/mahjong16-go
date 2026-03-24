---
allowed-tools: Read, Edit, Write, Bash(git branch:*), Bash(git add:*), Bash(git status:*), Bash(git commit:*), Bash(git log:*), Bash(git diff:*), Bash(git push:*), Bash(gh pr:*)
argument-hint: <功能名稱或任務描述>
description: 規劃＞實作＞斷點提交＞自動PR的工作流 (一次一個功能，提交訊息標準化，含歷程記錄)
---

# 規劃階段 (先不要改碼, 使用 plan mode)

在開始規劃前, 先檢查當前目錄是否存在 PROGRESS.md

- 如果存在 PROGRESS.md: 讀取它
  - 對於 [ ] 的項目: 照常使用, 等待後續實作
  - 對於 [/] 的項目: 問我是否要「重新規劃」這一項目, 或者是「繼續上次的進度」
    - 若選擇重新規劃: 覆寫該項的摘要與檔案範圍, 狀態保持 [/]
    - 若選擇繼續: 沿用原有內容, 維持原規劃

- 如果不存在 PROGRESS.md: 根據 $ARGUMENTS 的任務內容, 自動生成新的 PROGRESS.md 草稿, 內容須包含:
  - 任務清單 ([ ]狀態)
  - 每項任務的檔案範圍
  - 每項任務的摘要 (簡述要修改什麼)
  - 狀態更新規則 ([ ]:未開始; [/]:進行中; [v]:完成)

在規劃時, 確保任務數量控制在 3-8 項, 並附上驗收條件 (可手測或最小化測試)

> 在輸出草稿後, 請等待我回覆 (y/n):
- 「y」: 正式寫入 PROGRESS.md, 並立即執行以下初始化步驟:
  1. `git add PROGRESS.md`
  2. `git commit -m "chore: 新增任務規劃 PROGRESS.md"`
  3. `git push -u origin <當前分支名稱>` (將分支推送至 GitHub, 確保後續步驟可正常參照)
  4. 自動執行 `gh pr create --title "<任務名稱>" --body "## 任務說明\n$ARGUMENTS\n\n## 進度追蹤\n詳見 PROGRESS.md\n\n此 PR 由 auto-fix-step 自動建立，可於此 PR 繼續後續操作。"` 建立 Pull Request
  5. 顯示 PR 連結, 並提示: 「✅ PR 已建立, 您可於 PR 頁面繼續後續操作與審閱。」
- 「n」: 根據我的補充修改重新規劃

# 實作階段 (一次只做一項)

在開始實作前, 請先讀取 @PROGRESS.md, 找到第一個狀態為 [] 或 [/] 的項目

流程如下:
1. 根據 PROGRESS.md 的檔案範圍與摘要, 先描述預計修改內容
2. 將該項目狀態由 [ ] 改為 [/] (表示進行中)
3. **斷點提交 (進入中斷點)**: 立即執行一次階段提交, 保留進行中狀態:
   - `git add PROGRESS.md`
   - `git commit -m "chore: 標記 [任務名稱] 為進行中 [/]"`
4. 使用 `git diff` 比較修改前後差異, 並解釋細節
5. 詢問我是否要套用修改 (y/n)
   - **n**: 將該項目保持為 [/], 暫停修改
     - 如果我輸入「leave」: 直接退出指令, 下次再執行時會檢查 PROGRESS.md 裡的 [/] 項目, 並詢問是否要重新規劃或繼續進行
     - 如果我輸入「continue」: 參考我的建議再重新提交一次修改提案
     - 其他輸入: 與我討論修改方針, 直到輸入「continue」
   - **y**: 執行修改
6. **斷點提交 (完成中斷點)**: 修改套用後, 立即依照 @.claude/commands/pack-zh.md 的流程執行完整提交
7. 更新 @PROGRESS.md:
   - 將該項目狀態改為 [v] (已完成)
   - 在該項目下方新增一行, 紀錄「Commit 編號(前7碼)」與「提交訊息摘要」
8. **斷點提交 (更新進度)**: 提交 PROGRESS.md 的進度更新:
   - `git add PROGRESS.md`
   - `git commit -m "chore: 更新 PROGRESS.md - [任務名稱] 已完成 [v]"`
9. 若所有項目均為 [v]: 表示任務全部完成, 進入**歷程記錄生成**階段

> 若我回覆「next」, 繼續下一個項目; 若我回覆「end」, 則結束當前指令

# 歷程記錄生成階段 (所有任務完成後自動執行)

當所有 PROGRESS.md 項目均為 [v] 時, 自動執行以下步驟:

1. 取得當天日期 (格式: YYYYMMDD), 將 PROGRESS.md 重新命名為 `YYYYMMDD_revised_content.md`
2. 確認 `log_history` 資料夾:
   - 若有 `log_history` 資料夾: 將重新命名後的檔案移入其中
   - 若無: 自動建立 `log_history` 資料夾, 再將檔案移入
3. 在 `log_history` 資料夾中自動生成 `log_history_YYYYMMDD.md`, 內容包含:
   - 任務名稱與執行日期
   - 每項任務的完成摘要 (從 PROGRESS.md 擷取)
   - 對應 Commit 編號 (前7碼) 與提交訊息
   - 本次修改影響的檔案清單 (使用 `git diff --name-only` 產生)
4. 最後執行最終提交:
   - `git add log_history/`
   - `git commit -m "docs: 新增歷程記錄 log_history_YYYYMMDD.md，任務全部完成"`
   - `git push origin <當前分支名稱>`
5. 顯示完成訊息: 「✅ 所有任務已完成, 歷程記錄已儲存至 log_history/log_history_YYYYMMDD.md。請至 PR 頁面進行最終審閱與合併。