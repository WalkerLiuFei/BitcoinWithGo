package utils

import (
	"configs"
	"github.com/spf13/viper"
	"net"
	"os"
)

//获取本地的服务地址，因为Bitcoin要求IP地址的byte数组的长度需满足32字节的长度
func GetLocalServiceAddr() *net.TCPAddr {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops:" + err.Error())
		os.Exit(1)
	}
	ip := make([]byte, 12)
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = append(ip, ipnet.IP.To4()...)
				break
			}
		}
	}
	return &net.TCPAddr{
		IP:   ip,
		Port: viper.GetInt(configs.PORT),
	}
}

func checkerr(e error) {
	logger := GetSugarLogger()
	if e != nil {
		logger.Error(e.Error())
	}
}
