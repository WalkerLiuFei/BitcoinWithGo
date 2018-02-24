// Copyright (c) 2013-2014 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package network

import (
	"datastruct"
	"time"
	"golang.org/x/text/unicode/bidi"
)

type NodeAddress struct {
	address *datastruct.TimeStampNetworkAddress

	attempts int16

	lastAttempts time.Time

	lastSuccess time.Time

	tried bool
}

/**
+ It claims to be from the future
+ It hasn't been seen in over a month
+ It has failed at least three times and never succeeded
+ It has failed ten times in the last week
**/
func (na *NodeAddress) isValid() bool {
	if na.lastAttempts.After(time.Now().Add(-1 * time.Minute)) {
		return false
	}
	//From Future
	if na.address.Timestamp.After(time.Now().Add(1 * time.Minute)) {
		return true
	}

	if na.address.Timestamp.Before(time.Now().Add(-1 * A_MONTH_SECONDS)) {
		return false
	}

	if na.lastSuccess.IsZero() && na.attempts >= 3 {
		return false
	}
	if na.lastSuccess.Before(time.Now().Add(-1*A_WEEK_SECONDS)) && na.attempts >= 10 {
		return false
	}
	return true
}
