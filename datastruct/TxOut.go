package datastruct

type TxOut struct {
	//Transaction Value
	value int64

	//Usually contains the public key as a Bitcoin script setting up condition to claim this output
	PubKeyHash []byte
}