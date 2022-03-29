# VideoStation

West2 Assignment 4

[Postman auto doc](https://documenter.getpostman.com/view/16949749/UVsHS7Cv#intro)

[project address](https://github.com/Anxiu0101/VideoStation)

## Project Structure

![VideoStation](README.assets/VideoStation.svg)

### 用户模块 (user)

#### 用户结构体



- [x] username 用户名
- [x] password 密码
- [x] avatars 头像
- [x] 个人资料
    - [x] 性别 gender
    - [x] 年龄 age
    - [x] 邮箱 email
- [x] 状态(是否封号) state
- [ ] 收藏列表

#### 方法

- [x] 用户登录
- [x] 用户注册
- [ ] 修改密码
- [x] 获取用户资料
- [x] 修改用户资料
- [ ] 获取收藏列表
- [ ] 拉黑用户

### 视频模块 (video)

#### 视频结构体

- [x] 视频文件主体

- [x] 状态
    - 待审核
    - 已审核
    - 草稿
- [ ] 弹幕列表
- [ ] 点赞列表
- [x] 评论列表
- [ ] 转发数
- [ ] 收藏数

#### 弹幕结构体

- [x] 发送用户ID
- [x] 弹幕内容
- [x] 弹幕位置（在视频时间线的位置）

#### 评论结构体

- [x] 发送用户ID
- [x] 评论内容
- [x] 回复对象ID
    - [x] 置零为不回复，为视频评论

#### 收藏结构体

- [x] 用户ID UID
- [x] 视频ID VID
- [x] 收藏组

#### 方法

- [x] 用户上传视频

- [x] 用户点赞

- [x] 用户评论
    
    > 设定了一个默认用户，id为0，当用户评论视频时，receiver id 就是0。若用户回复某位用户的评论时，receiver id 就是被回复的用户
    
- [x] 用户收藏

- [ ] 用户转发

- [ ] 用户发送弹幕

### 管理员模块 (admin)

#### 管理员结构体

- [ ] adminName 管理员名
- [ ] adminPassword 管理员密码
- [ ] adminImg 管理员头像

#### 方法

- [ ] 获取待审核视频列表

  > 如何解决数个管理员同时管理一个视频的冲突情况？

- [ ] 封禁用户账号


## Problems

- 部分接口获取用户ID需要参数，而没有从 token 中解析

- 文件上传未处理
    - 文件重名问题 
    
- [x] favorite 功能不依赖 Interactive 结构体而是有自己的结构体，未并入

    > 现已并入

- Upload 不能解析 introduction 字段，问题未排查，从传入时就有问题。

- Like and Favorite 数据收集是，gorm count 返回的类型是 int64，被我强转了，这个问题需要解决

- 在统一类型上遇到了很多问题，id的类型是uint，但是传入的数字类型一般都是int，在数据库操作时，gorm的部分函数返回格式又是int64

- 到底什么时候需要错误处理？



## 接口思路

### Show Video

Show video 这个接口可以分为两个部分，

- Video info 包含视频标题，视频介绍，用户等，从 mysql 中获取数据
- Video data 包含点击量，收藏关系，点赞关系，从 redis 中获取数据

现在的问题是，interactive 表创建失效，需要寻找几个模型的关系，但是这个问题只要关注 gorm 语句和模型的代码即可。



