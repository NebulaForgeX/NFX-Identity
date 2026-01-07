# User Educations Table

## 概述 / Overview

`directory.user_educations` 表用于存储用户的教育经历，一个用户可以有多个教育记录。

The `directory.user_educations` table stores user education history, one user can have multiple education records.

## 使用场景 / Use Cases

### 1. 添加教育经历
**场景**: 用户添加教育经历
**例子**:
- 用户添加"清华大学计算机科学学士"教育经历
- 操作：`INSERT INTO directory.user_educations (user_id, school, degree, major, start_date, end_date, is_current) VALUES ('zhangsan-uuid', '清华大学', 'Bachelor', '计算机科学', '2015-09-01', '2019-06-30', false)`
- 结果：教育经历添加，可以在个人资料显示

### 2. 添加在读教育
**场景**: 用户添加当前在读的教育
**例子**:
- 用户添加"北京大学计算机科学硕士（在读）"
- 操作：`INSERT INTO directory.user_educations (user_id, school, degree, major, start_date, end_date, is_current) VALUES ('zhangsan-uuid', '北京大学', 'Master', '计算机科学', '2020-09-01', NULL, true)`
- 结果：在读教育添加

