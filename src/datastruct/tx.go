package datastruct


import "common"

type Tx interface{
	Init(input *common.BitcoinInput)

	GetBytes() []byte
}

type TxOut struct {
	//Transaction Value
	Value int64

	//Usually contains the public key as a Bitcoin script setting up condition to claim this output
	PubKeyHash []byte
}

type TxIn struct {
	PreviousOutput OutPoint
	//签名
	SigScript []byte

	//uint32, Transaction version as defined by the sender. Intended for
	//"replacement" of transactions when information is updated before
	// inclusion into a block.

	Sequence uint32
}

func (tx *TxOut) Init(input *common.BitcoinInput) {
	input.ReadNum(&tx.Value)
	scriptLen,err  := input.ReadVarInt()
	if err != nil{
		logger.Error(err.Error())
		return
	}
	tx.PubKeyHash = make([]byte,scriptLen)
	input.ReadBytes(tx.PubKeyHash)
}

func (tx *TxOut) GetBytes()  []byte{
	output := common.BitcoinOuput{}
	output.WriteNum(tx.Value).
		WriteBytes(tx.PubKeyHash)
	return output.Buffer.Bytes()
}

func (tx *TxIn) Init(input *common.BitcoinInput) {
	tx.PreviousOutput = OutPoint{}
	tx.PreviousOutput.Init(input)
	signatureLen, err := input.ReadVarInt()
	if err != nil{
		logger.Error(err.Error())
		return
	}
	tx.SigScript = make([]byte, signatureLen)
	input.ReadBytes(tx.SigScript)
	input.ReadNum(&tx.Sequence)
}

func (tx *TxIn) GetBytes() []byte {
	output := common.BitcoinOuput{}
	output.WriteBytes(tx.PreviousOutput.Hash).
		WriteNum(tx.PreviousOutput.Index).
		WriteBytes(tx.SigScript).
		WriteNum(tx.Sequence)
	return output.Buffer.Bytes()
}
