name="sysinfocollector"
home_dir=`pwd`


# 运行频率 多少分钟一次 默认1分钟一次 
m=1
# 保存天数 默认保存5天
d=5	
	
#if [ "$1" == "" ]; then
#    # 运行频率 多少分钟一次
#    m=1
#else
#    m=$1
#fi
#if [ "$2" == "" ]; then
#    # 保存天数
#    d=5
#else
#    d=$2
#fi

echo """#!/bin/bash

start(){
	chmod +x $home_dir/$name
    nohup $home_dir/$name -m $m -s $d > /dev/null 2>&1 &
}

stop(){
    ps -ef | grep '$home_dir/$name -m $m -s $d' | grep -v "grep" | awk '{print \$2}' | xargs kill 

}

case "\$1" in
start)
    start
    ;;
stop)
    stop
    ;;
*)
    echo "unknow args"
    ;;
esac

""" > /etc/$name.sh
chmod +x /etc/$name.sh
echo """
[Unit]
Description=Media wanager Service
After=network.target
 
[Service]
Type=forking
 
ExecStart=/bin/bash -c '/etc/$name.sh start'
ExecStop=/etc/$name.sh stop
 
[Install]
WantedBy=multi-user.target

""" > /usr/lib/systemd/system/${name}.service

systemctl enable  ${name} --now 
systemctl start ${name}