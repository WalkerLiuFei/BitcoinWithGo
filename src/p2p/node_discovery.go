package p2p

import (
	"net"
	"configs"
	"github.com/spf13/viper"
)

var DNS_SEEDS = []string{"bitseed.xf2.org", "dnsseed.bluematt.me", "seed.bitcoin.sipa.be", "dnsseed.bitcoin.dashjr.org", "seed.bitcoinstats.com"}

func ConsumeDNSSeed() [](*net.TCPAddr) {
	result := make([](*net.TCPAddr), len(DNS_SEEDS))
	for _, dnsAddr := range DNS_SEEDS {
		ipArray, err := net.LookupIP(dnsAddr)
		checkerr(err, ConsumeDNSSeed)
		for _, ip := range ipArray {
			tcpAddr := &net.TCPAddr{
				IP:   ip,
				Port: viper.GetInt(configs.PORT),
			}
			result = append(result, tcpAddr)
		}
	}
	return result
}
