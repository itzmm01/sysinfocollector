## 安装

解压压缩文件后执行`bash install.sh`

```shell
# 运行频率 多少分钟一次和保存天数 可修改install.sh脚本 默认1分钟一次 保存5天
bash install.sh  
```

安装完成后会自动启动，程序日志为`/var/log/monitor/monitor.log`以日期命名

收集的日志数据会存放在`/var/log/monitor`目录下,以日期命名

## 相关命令

```shell
# 查看运行状态
systemctl status  sysinfocollector
# 停止
systemctl stop  sysinfocollector
# 启动
systemctl restart  sysinfocollector
```

