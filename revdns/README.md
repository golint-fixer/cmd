revdns
======

revdns performs reverse DNS lookups on provided IP-addresses.

Installation
------------

	go get github.com/mewkiz/cmd/revdns

Documentation
-------------

Documentation provided by GoDoc.

- [revdns][]: Perform reverse DNS lookups on provided IP-addresses.

[revdns]: http://godoc.org/github.com/mewkiz/cmd/revdns

Usage
-----

	revdns IP...

Examples
--------

1. Lookup of external IP-address.

	revdns 208.80.152.201
	// Output:
	// 208.80.152.201 = wikipedia-lb.pmtpa.wikimedia.org.
