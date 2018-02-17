package p2p


//同步本地的区块链
func SyncBlockChain() {
	go UpdateUsefulNode()

}
