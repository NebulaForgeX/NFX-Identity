# User Badges Table

## 概述 / Overview

`directory.user_badges` 表用于用户和徽章的多对多关系，记录用户获得的徽章。

The `directory.user_badges` table represents the many-to-many relationship between users and badges.

## 使用场景 / Use Cases

### 1. 授予用户徽章
**场景**: 系统或管理员授予用户徽章
**例子**:
- 用户完成某个成就，获得"早期用户"徽章
- 操作：`INSERT INTO directory.user_badges (user_id, badge_id, level, description, earned_at) VALUES ('zhangsan-uuid', (SELECT id FROM directory.badges WHERE name = 'early_user'), 1, '感谢您成为早期用户', NOW())`
- 结果：用户获得徽章，可以在个人页面显示

### 2. 查询用户徽章
**场景**: 查看用户的所有徽章
**例子**:
- 查询用户徽章：`SELECT b.name, b.description, b.icon_url, ub.level, ub.earned_at FROM directory.user_badges ub JOIN directory.badges b ON ub.badge_id = b.id WHERE ub.user_id = 'zhangsan-uuid' ORDER BY ub.earned_at DESC`
- 结果：返回用户的所有徽章

### 3. 升级徽章等级
**场景**: 用户徽章等级提升
**例子**:
- 用户徽章从 1 级升级到 2 级
- 操作：`UPDATE directory.user_badges SET level = 2, updated_at = NOW() WHERE user_id = 'zhangsan-uuid' AND badge_id = 'badge-uuid'`
- 结果：徽章等级更新

