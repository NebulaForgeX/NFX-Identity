# User Occupations Table

## 概述 / Overview

`directory.user_occupations` 表用于存储用户的工作/就业经历，一个用户可以有多个工作记录。

The `directory.user_occupations` table stores user work/employment history, one user can have multiple occupation records.

## 使用场景 / Use Cases

### 1. 添加工作经历
**场景**: 用户添加工作经历
**例子**:
- 用户添加"Google 软件工程师"工作经历
- 操作：`INSERT INTO directory.user_occupations (user_id, company, position, department, start_date, end_date, is_current, employment_type) VALUES ('zhangsan-uuid', 'Google', '软件工程师', '搜索部门', '2020-01-01', '2022-12-31', false, 'full-time')`
- 结果：工作经历添加

### 2. 添加当前工作
**场景**: 用户添加当前工作
**例子**:
- 用户添加"Microsoft 高级软件工程师（当前）"
- 操作：`INSERT INTO directory.user_occupations (user_id, company, position, start_date, is_current, employment_type) VALUES ('zhangsan-uuid', 'Microsoft', '高级软件工程师', '2023-01-01', true, 'full-time')`
- 结果：当前工作添加

