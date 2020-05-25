package main

import (
	"log"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

/**
centos7 
sudo yum install libevent memcached telnet
sudo service memcached start
sudo chkconfig memcached on
telnet localhost 11211
quit

开启防火墙端口
sudo vim /etc/sysconfig/iptables
添加如下
-A INPUT -p tcp -m tcp --dport 11211 -j ACCEPT

sudo service iptables restart
sudo service network restart
*/

func main() {
	// support server list.
	mc := memcache.New("192.168.0.11:11211")
	mc.Timeout = 2 * time.Second
	mc.MaxIdleConns = 300

	// mc.Delete("foo")

	it, err := mc.Get("foo")

	if err != nil {
		if err != memcache.ErrCacheMiss {
			log.Println("mc error: ", err)
			return
		}

		err = mc.Set(&memcache.Item{Key: "foo", Value: []byte("my value"), Expiration: 5 * 60})
		if err != nil {
			log.Println("mc set error: ", err.Error())
			return
		}

		log.Println("mc set success")
	} else {
		log.Println(string(it.Value), err)
	}
}
