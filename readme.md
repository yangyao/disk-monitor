## Ubuntu 磁盘监控

用途说明 : 家里的服务器使用 USB 外接了一个移动硬盘。外接的硬盘非常不稳定，总是出现掉盘的现象。
本脚本，会定时检查磁盘是否被系统正确识别，如果找不到磁盘了，那么发送一条server 酱微信通知

### 使用方法

- clone 项目
- 新建一个 .env 文件
- 在系统上运行 `blkid` 找到你磁盘的UUID 
- 编辑 .env 文件，增加一行 UUID=<你的移动硬盘的 UUID>
- 编辑 .env 文件，增加一行 SC_KEY=<server酱的 API key>
- 在系统上运行 `crontab -e` 新增一条 cron 
- cron 执行的命令 `go run main.go`