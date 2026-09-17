# 螭吻·九九 Story Engine PoC 输入文件

版本：v0.3\
用途：Novel Studio Story Engine 第一阶段 PoC\
平台：抖音

## 核心原则

-   AI负责提出可能性，人负责决定什么是真的。
-   IDEA ≠ FACT
-   PROPOSAL ≠ CANON
-   SIMULATION ≠ HISTORY
-   DRAFT ≠ ACCEPTED EPISODE
-   AI ACCEPT ≠ HUMAN APPROVAL
-   "继续、试试、模拟、深化、先跑一版"均不等于Canon批准。

## Canon等级

-   **LEVEL 0 / HARD CANON**：正式确定，AI不得修改或违反。
-   **LEVEL 1 / APPROVED CANON**：已经确认，修改必须人工批准。
-   **LEVEL 2 / WORKING ASSUMPTION**：仅供PoC临时使用。
-   **TBD**：当前未知，AI不得自动补全，只能提出Proposal。

所有AI新增设定默认 `PROPOSED`，不得污染Accepted Character
State、Knowledge、Relationship、Timeline、Foreshadow、World State或Canon
RAG。

## 螭吻·九九

### HARD CANON

-   正式名称：螭吻·九九；简称九九。
-   神话原型：螭吻，龙生九子中的第九子。
-   喜水，具有灭火属性。
-   神兽本体核心结构：**龙头鱼身**，鱼类躯干为主体。
-   不得写成麒麟、普通四足龙、西方龙、单纯蛇形东方龙、陆生神兽或机械载具。

### APPROVED CANON

-   人格：好奇、好胜、爱面子、有担当。
-   主要行动动力：主动探索。
-   核心价值底线：守护。
-   缺点可由好胜、爱面子、希望证明自己引发显摆、逞强、嘴硬、过度尝试和错误判断，并产生真实后果。

### 能力边界

已确认与水高度相关、喜水、具有灭火属性。降雨、吞火、守护可继续验证。能力强度、次数、来源、代价、升级、战斗等级、额珠关系及形态差异均为TBD。

## 九九与渔岛

九九IP与渔岛海洋度假区存在明确联系。"独角龙一夜开海口"是重要地方文化来源；如何纳入线上连续剧情为TBD。

## 泼水节

正式名称为**泼水节**；"扣水节"为错误名称。九九已经参与渔岛泼水节，相关基础概念包括水、祈福、互动、喷水祈福。完整神话解释为TBD。

## 世界方向

东方 / 中式 / 奇幻 / 海岸 / 海洋 / 度假。避免整体世界观向西式魔幻偏移。

## 真九九、机甲九九、凌墨

三者必须区分。线上是否继续采用既有元神、机甲、合体机制为TBD。

## 九尾狐

### HARD CANON

-   神兽原型：九尾狐。
-   神兽本体必须是真正九尾狐，普通狐狸、单尾狐狸、狐耳人形均不能替代。

### TBD

正式姓名、性别、年龄感、性格、身份、能力、Want、Need、Misbelief、Secret、目标、与九九/凤凰关系及人物弧均未确定。新增内容必须以Proposal形式提出。

## 羽族凤凰

### HARD CANON

-   神兽原型：凤凰，不是青鸾。
-   神兽本体必须是真正、完整的凤凰神鸟形态。
-   人形女性+翅膀、天使、西式翼人、狮鹫、普通鸟类等不得等同凤凰本体。
-   属于东方神话/东方奇幻体系。

### APPROVED CANON / DIRECTION

-   与羽族存在明确关联，但羽族的种族/文明/势力性质仍为TBD。
-   人形视觉探索方向：白、金、浅蓝、东方凤冠、凤凰纹样、舒展长尾羽、东方云上宫阙。
-   神兽凤凰不要求颜色、宝石与九九一致。

### TBD边界

-   是否拥有正式化人能力
-   正式性别（现有人形仅为女性视觉表现）
-   正式姓名、年龄、身份地位
-   核心性格、Want、Need、Current Goal、Misbelief、Secret
-   完整能力体系；不得自动赋予火焰、涅槃、重生、治愈、不死等传统凤凰能力
-   与火的关系
-   与九九、九尾狐的初始关系
-   CP关系
-   世界秘密与历史知识

### 凤凰Knowledge起点

采用 **MINIMUM-KNOWLEDGE
INITIALIZATION**：没有明确Canon证明凤凰知道的事情，不得进入Known Facts。

当前可确认： - KNOWN：自身凤凰身份；自身与羽族存在关联。 -
SUSPECTED：无已批准项目。 - FALSE_BELIEF：无已批准项目。 -
FORBIDDEN：其他角色私有秘密、未批准未来剧情、Arc
Rehearsal未来结果、其他Proposal分支内容。

凤凰是否认识九九、是否知道九九是螭吻、是否知道九九能力、独角龙传说、泼水节、现实渔岛、九尾狐及世界重大秘密，全部为TBD。

Knowledge统一状态：KNOWN / SUSPECTED / BELIEVED / FALSE_BELIEF / UNKNOWN
/ FORBIDDEN。

任何知识状态变化必须记录来源事件、证据、Episode和story
time。模型推理本身不能让UNKNOWN直接变成KNOWN。

## Phase 0范围

本轮禁止Season Planning，只完成： 1. 整理Hard / Approved Canon； 2.
建立九九Initial Character State；缺失项标记TBD或PROPOSAL_NEEDED； 3.
九尾狐生成FOX-A/B/C三套隔离Proposal； 4.
凤凰生成PHOENIX-A/B/C三套隔离Proposal； 5. 完成Knowledge Boundary、Hard
Canon、Character Logic和Proposal Isolation测试； 6. STOP，等待人工审核。

不得自动选择Proposal、Commit新Canon、规划Season或写Episode。
