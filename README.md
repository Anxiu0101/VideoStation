# VideoStation

West2 Assignment 4

## Project Structure

![VideoStation](README.assets/VideoStation.svg)

### 用户模块 (user)

#### 用户结构体
- username 用户名
- password 密码
- avatars 头像
- 个人资料
    - 性别 gender
    - 年龄 age
    - 邮箱 email
- 状态(是否封号) state
- 收藏列表

#### 方法

- 用户登录
- 用户注册
- 修改密码
- 获取用户资料
- 修改用户资料
- 获取收藏列表
- 拉黑用户

### 视频模块 (video)

#### 视频结构体

- 视频文件主体

- 状态
    - 待审核
    - 已审核
    - 草稿
- 弹幕列表
- 点赞列表
- 评论列表
- 转发数
- 收藏数

#### 弹幕结构体

- 发送用户ID
- 弹幕内容
- 弹幕位置（在视频时间线的位置）

#### 评论结构体

- 发送用户ID
- 评论内容
- 回复对象ID
    - 置零为不回复，为视频评论

#### 方法

- 用户上传视频
- 用户点赞
- 用户评论
    - 楼中楼功能？没思路
- 用户收藏
- 用户转发
- 用户发送弹幕

### 管理员模块 (admin)

#### 管理员结构体

- adminName 管理员名
- adminPassword 管理员密码
- adminImg 管理员头像

#### 方法

- 获取待审核视频列表

  > 如何解决数个管理员同时管理一个视频的冲突情况？

- 封禁用户账号

- 