package flags

import "flag"

var (
	// commandline flags
	// -l --listen  the address http server listen for.
	ListenAddr string
)

var (
	// default commandline flag values
	defaultListenAddr = "127.0.0.1:12345"
)

func init() {
	flag.StringVar(&ListenAddr, "l", defaultListenAddr,
		"-l --listen  the address http server listen for. default is "+defaultListenAddr)
}
