package bitcoinWithGo

import (
	"configs"
	"p2p"
)

func main() {
	configs.InitConfigs()
	p2p.UpdateUsefulNode()

}
