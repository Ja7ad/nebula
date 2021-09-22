// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"strconv"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/strmangle"
)

// M type is for providing columns and column values to UpdateAll.
type M map[string]interface{}

// ErrSyncFail occurs during insert when the record could not be retrieved in
// order to populate default value information. This usually happens when LastInsertId
// fails or there was a primary key configuration that was not resolvable.
var ErrSyncFail = errors.New("models: failed to synchronize data after insert")

type insertCache struct {
	query        string
	retQuery     string
	valueMapping []uint64
	retMapping   []uint64
}

type updateCache struct {
	query        string
	valueMapping []uint64
}

func makeCacheKey(cols boil.Columns, nzDefaults []string) string {
	buf := strmangle.GetBuffer()

	buf.WriteString(strconv.Itoa(cols.Kind))
	for _, w := range cols.Cols {
		buf.WriteString(w)
	}

	if len(nzDefaults) != 0 {
		buf.WriteByte('.')
	}
	for _, nz := range nzDefaults {
		buf.WriteString(nz)
	}

	str := buf.String()
	strmangle.PutBuffer(buf)
	return str
}

// Enum values for crawl_state
const (
	CrawlStateStarted   = "started"
	CrawlStateCancelled = "cancelled"
	CrawlStateFailed    = "failed"
	CrawlStateSucceeded = "succeeded"
)

// Enum values for dial_error
const (
	DialErrorIoTimeout               = "io_timeout"
	DialErrorConnectionRefused       = "connection_refused"
	DialErrorProtocolNotSupported    = "protocol_not_supported"
	DialErrorPeerIDMismatch          = "peer_id_mismatch"
	DialErrorNoRouteToHost           = "no_route_to_host"
	DialErrorNetworkUnreachable      = "network_unreachable"
	DialErrorNoGoodAddresses         = "no_good_addresses"
	DialErrorContextDeadlineExceeded = "context_deadline_exceeded"
	DialErrorNoPublicIP              = "no_public_ip"
	DialErrorMaxDialAttemptsExceeded = "max_dial_attempts_exceeded"
	DialErrorUnknown                 = "unknown"
	DialErrorMaddrReset              = "maddr_reset"
)
