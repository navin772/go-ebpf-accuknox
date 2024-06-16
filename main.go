package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/rlimit"
)

const (
	ifaceName = "lo" // interface name
	xdpFlags  = 0
)

func main() {
	// Parse command line argument for the port number, pass the port via -port flag, eg: -port 7000
	portPtr := flag.Int("port", 4040, "Port number to drop TCP packets on")
	flag.Parse()

	port := uint16(*portPtr)

	// Allow the current process to lock memory for eBPF
	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatalf("RemoveMemlock: %v", err)
	}

	// Load the compiled eBPF program
	spec, err := ebpf.LoadCollectionSpec("bpf/xdp_prog.o")
	if err != nil {
		log.Fatalf("LoadCollectionSpec: %v", err)
	}

	coll, err := ebpf.NewCollection(spec)
	if err != nil {
		log.Fatalf("NewCollection: %v", err)
	}
	defer coll.Close()

	// Look up the XDP program by name
	prog, found := coll.Programs["xdp_drop_tcp"]
	if !found {
		log.Fatalf("Program 'xdp_drop_tcp' not found")
	}

	// Get the interface index
	iface, err := net.InterfaceByName(ifaceName)
	if err != nil {
		log.Fatalf("InterfaceByName: %v", err)
	}

	// Attach the program to the network interface
	link, err := link.AttachXDP(link.XDPOptions{
		Program:   prog,
		Interface: iface.Index,
		Flags:     xdpFlags,
	})
	if err != nil {
		log.Fatalf("AttachXDP: %v", err)
	}
	defer link.Close()

	// Retrieve the map for updating the port number
	portMap, found := coll.Maps["port_map"]
	if !found {
		log.Fatalf("Map 'port_map' not found")
	}

	// Set the port number
	key := uint32(0)
	if err := portMap.Put(key, port); err != nil {
		log.Fatalf("Put: %v", err)
	}

	fmt.Printf("eBPF program loaded and attached. Dropping TCP packets on port %d. Press Ctrl+C to exit.\n", port)
	select {}
}
