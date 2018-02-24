package datastruct

import "time"

type TimeStampNetworkAddress struct {
	//last communicate time
	Timestamp time.Time

	NetworkAddress *NetworkAddress
}
