# go memcache 使用
    centos7 
    sudo yum install libevent memcached telnet
    sudo service memcached start
    sudo systemctl enable memcached
    sudo systemctl status memcached
    sudo chkconfig memcached on
    telnet localhost 11211
    quit

    开启防火墙端口
    sudo vim /etc/sysconfig/iptables
    添加如下
    -A INPUT -p tcp -m tcp --dport 11211 -j ACCEPT

    sudo service iptables restart
    sudo service network restart

    go run main.go

# 参考文档

    https://github.com/golang/groupcache
