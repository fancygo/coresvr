# srv
#游戏简介
* core服务器实现 

#服务器实现
* 作为中心服务器
* 数据库建库, 落地全部使用gorm进行操作
* 所有内部连接使用tcp
* 协议通信使用的proto格式
* 并发模型使用的多线程, 每个客户端连接会创建一个连接线程, 消息收发通过channel送到不同的接受和发送线程

#各文件夹介绍
* admin 配置脚本
* bin 可执行文件
* conf 服务器配置文件, 暂时只有mainserver和mysql是起作用的
* log 游戏运行日志
* pkg 编译生成的中间文件
* script 启动脚本等
* src 源代码, 各文件作用, 见注释

#编译
./admin make生成core_server和mysql两个可执行文件
* mysql 初始化数据库
* core_server 游戏主服务
