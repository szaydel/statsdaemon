// Copyright 2015 RackTop Systems Inc. and/or its affiliates.
// http://www.racktopsystems.com
//
// The methods and techniques utilized herein are considered TRADE SECRETS
// and/or CONFIDENTIAL unless otherwise noted. REPRODUCTION or DISTRIBUTION
// is FORBIDDEN, in whole and/or in part, except by express written permission
// of RackTop Systems.

package main
import (
	"flag"
)
/*
 Global Variables
*/


/*
Flags
*/
var (
	serviceAddress    = flag.String("address", ":8125", "UDP service address")
	tcpServiceAddress = flag.String("tcpaddr", "", "TCP service address, if set")
	maxUdpPacketSize  = flag.Int64("max-udp-packet-size", 1472, "Maximum UDP packet size")
	postFlushCmd      = flag.String("post-flush-cmd", "stdout", "Command to run on each flush")
	graphiteAddress   = flag.String("graphite", "127.0.0.1:2003", "Graphite service address (or - to disable)")
	flushInterval     = flag.Int64("flush-interval", 10, "Flush interval (seconds)")
	debug             = flag.Bool("debug", false, "print statistics sent to graphite")
	showVersion       = flag.Bool("version", false, "print version string")
	deleteGauges      = flag.Bool("delete-gauges", true, "don't send values to graphite for inactive gauges, as opposed to sending the previous value")
	persistCountKeys  = flag.Int64("persist-count-keys", 60, "number of flush-intervals to persist count keys")
	receiveCounter    = flag.String("receive-counter", "", "Metric name for total metrics received per interval")
	percentThreshold  = Percentiles{}
	prefix            = flag.String("prefix", "", "Prefix for all stats")
	postfix           = flag.String("postfix", "", "Postfix for all stats")
)
