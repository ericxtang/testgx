package testgx

import (
	"os"

	ma "github.com/multiformats/go-multiaddr"

	"github.com/golang/glog"
)

func main() {
	glog.Infof("Print")
	addr, err := ma.NewMultiaddr("/ip4/0.0.0.0/tcp/" + os.Args[2])
	if err != nil {
		glog.Errorf("Error")
	}
	glog.Infof("Addr: %v", addr)
}
