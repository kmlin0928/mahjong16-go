# 胡牌時顯示四家手牌

## 任務清單

- [v] Task 1：GameState 新增 all_hands 欄位，game_over 時填入四家手牌
  - Commit 791e9b4：feat(state): GameState 新增 all_hands，game_over 時填入四家手牌
  - 檔案：mahjong.py（GameState dataclass、_snapshot 方法）
  - 摘要：新增 all_hands: list[list[str]] | None = None；
    phase=="game_over" 時填入四家 sorted hand 的牌名列表，
    其他 phase 保持 None

- [v] Task 2：前端 game_over 時渲染四家手牌
  - Commit 5c94900：feat(ui): 遊戲結束時顯示四家手牌
  - 檔案：static/app.js（showGameOver 函式）
          static/index.html（gameover-banner 加入手牌區塊）
          static/style.css（hand reveal 樣式）
  - 摘要：showGameOver 讀取 state.all_hands；
    在台數下方逐家顯示「東：1萬 2萬 …」；
    使用與面牌相同的 tile 樣式

## 狀態規則
[ ]:未開始　[/]:進行中　[v]:完成

## 驗收條件
- 遊戲結束（胡牌或和局）後，gameover-banner 內出現四家手牌
- 手牌顯示排序後牌名與 Unicode emoji
- 遊戲進行中不顯示 AI 手牌（all_hands = null）
