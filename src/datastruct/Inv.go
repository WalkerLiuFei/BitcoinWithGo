package datastruct

//Allows a node to advertise its knowledge of one or more objects.
// It can be received unsolicited, or in reply to getblocks.
//Payload (maximum 50,000 entries, which is just over 1.8 megabytes):

//允许节点广播一个或者多个对象，它可以接受主动的推送消息，也可用来作为 getblocks的返回值类型
//容量 最大为 500000 个对象  约为1.8 M
type Inv struct {
	//Number of inventory entries
	count uint32

	//Inventory vectors
	inventory []InvVect
}