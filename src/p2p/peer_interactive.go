package p2p

import "net"

//同步本地的区块链
func SyncBlockChain() {
	peerChannel := make(chan *net.TCPAddr, 1)
	go UpdateUsefulNode(peerChannel)
	select {
		peerConn := <-peerChannel

	}
}
