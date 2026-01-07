# Badges Table

## 概述 / Overview

`directory.badges` 表用于存储徽章定义，用于用户成就和认可。

The `directory.badges` table stores badge definitions for user achievements and recognition.

## 使用场景 / Use Cases

### 1. 创建徽章
**场景**: 系统创建新徽章
**例子**:
- 创建"早期用户"徽章
- 操作：`INSERT INTO directory.badges (name, description, icon_url, color, category, is_system) VALUES ('early_user', '早期用户', 'https://example.com/icons/early_user.png', '#FFD700', 'achievement', true)`
- 结果：徽章创建，可以授予给用户

### 2. 授予用户徽章
**场景**: 系统或管理员授予用户徽章
**例子**:
- 授予用户"早期用户"徽章
- 操作：`INSERT INTO directory.user_badges (user_id, badge_id, level, description) VALUES ('zhangsan-uuid', (SELECT id FROM directory.badges WHERE name = 'early_user'), 1, '感谢您成为早期用户')`
- 结果：用户获得徽章，可以在个人页面显示

