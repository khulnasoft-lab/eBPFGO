// Copyright (c) 2020 khulnasoft-lab, Inc.

package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/khulnasoft-lab/ebpfgo"
	"github.com/vishvananda/netlink"
)

func AttachXdp(x ebpfgo.Program, iList []string, devMap ebpfgo.Map) error {
	for _, intf := range iList {
		ifData, err := netlink.LinkByName(intf)
		if err != nil {
			fmt.Printf("Could not find link %s!\n", intf)
			return err
		}
		ifIndex := ifData.Attrs().Index
		if err := x.Attach(&ebpfgo.XdpAttachParams{Interface: intf, Mode: ebpfgo.XdpAttachModeSkb}); err != nil {
			fmt.Printf("Failed attaching xdp prog to %s!\n", intf)
			return err
		}
		if err := devMap.Upsert(ifIndex, ifIndex); err != nil {
			return err
		}
		fmt.Printf("XDP program attached to %s. to detach, use `ip -f link set dev %s xdp off`\n", intf, intf)
	}
	return nil
}

func FetchMaps(prog ebpfgo.System) (devMap ebpfgo.Map, err error) {
	devMap = prog.GetMapByName("if_redirect")
	if devMap == nil {
		err = fmt.Errorf("failed fetching map if_redirect from program.\n")
	}
	return
}

func main() {
	var xdpProgFile, xdpIntf string

	flag.StringVar(&xdpProgFile, "file", "", "xdp binary to attach")
	flag.StringVar(&xdpIntf, "i", "", "interfaces to attach xdp code to. comma separated")
	flag.Parse()

	if xdpProgFile == "" {
		panic("Please enter a valid filename.")
	}
	if xdpIntf == "" {
		panic("Please enter a valid interface name.")
	}
	intfList := strings.Split(strings.Replace(xdpIntf, " ", "", 0), ",")

	bpf := ebpfgo.NewDefaultEbpfSystem()
	if err := bpf.LoadElf(xdpProgFile); err != nil {
		panic(err)
	}
	xdpProg := bpf.GetProgramByName("xdp_test")
	if xdpProg == nil {
		panic(fmt.Sprintf("Could not find xdp_test program in %s", xdpProgFile))
	}
	devMap, err := FetchMaps(bpf)
	if err != nil {
		panic(err)
	}
	if err := xdpProg.Load(); err != nil {
		panic(err)
	}
	fmt.Printf("Attaching program %s\n", xdpProg.GetName())
	if err := AttachXdp(xdpProg, intfList, devMap); err != nil {
		panic(err)
	}
}
