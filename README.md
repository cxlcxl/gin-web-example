## gin-web-example

基于 go 语言的 gin 框架项目骨架，做了基础的封装，方便 web 开发快速启动，快速提供 api 开发。

### 集成模块
1. logrus 日志
2. redis
3. gorm mysql
4. redis 数据库缓存
5. config 配置获取
6. air 热更新

### 使用

下载项目源码
```perl
git clone https://github.com/cxlcxl/gin-web-example.git
```

进入项目目录


初始化模块
```perl
go mod init xxxx
```

下载依赖
```perl
go mod tidy
```

启动（windows），mac 或其他操作系统需要修改 .air.conf 配置文件
```perl
air
```