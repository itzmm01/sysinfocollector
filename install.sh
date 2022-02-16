name="sysinfocollector"
home_dir=`pwd`

if [ "$1" == "" ]; then
    m=1
else
    m=$1
fi
if [ "$2" == "" ]; then
    d=5
else
    d=$2
fi

echo """#!/bin/bash

start(){
    cd $home_dir
	chmod +x $name
    nohup ./$name -m $m -s $d > /dev/null 2>&1 &
}

stop(){
    ps -ef | grep "./$name -m $m -s $d" | grep -v "grep" | awk '{print \$2}' | xargs kill 

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

""" > $home_dir/$name.sh
chmod +x $home_dir/$name.sh
echo """
[Unit]
Description=Media wanager Service
After=network.target
 
[Service]
Type=forking
 
ExecStart=/bin/bash -c '$home_dir/${name}.sh start'
ExecStop=$home_dir/${name}.sh stop
 
[Install]
WantedBy=multi-user.target

""" > /usr/lib/systemd/system/${name}.service

systemctl enable  ${name} --now 
systemctl start ${name}