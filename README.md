# Library-Api

这是一个作业, 正在施工

### 功能实现

基础功能

- 用户的登录与注册(token+refresh_token)
- 用户信息的更改以及密码修改
- 通过名字查找书籍
- 通过标签查找书籍
- 收藏书籍(focus)
- 显示用户收藏的书籍
- 点赞书籍(praise)
- 获得书籍的所有书评
- 书评的查增删改
- 媒体资源访问(gin静态文件服务)

其他功能

- 实现RAS+SHA256的JWT接口
- 使用系统调用监听文件事件, 重启服务(便于调试)
- 数据库操作重试机制(增加可靠性)
- 书评的嵌套(递归)
- 匿名的书评(隐藏用户名)