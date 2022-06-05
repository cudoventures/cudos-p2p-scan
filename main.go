package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	trans "github.com/CudoVentures/cudos-p2p-scan/transport"
	"github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/version"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Example usage: cudos-p2p-scan host:port")
		os.Exit(1)
	}

	nodeKey := p2p.NodeKey{
		PrivKey: ed25519.GenPrivKey(),
	}
	nodeInfo := testNodeInfo(nodeKey.ID(), "node-1")
	transport := trans.NewMultiplexTransport(nodeInfo, nodeKey, p2p.MConnConfig(config.DefaultP2PConfig()))

	tcpAddr, err := net.ResolveTCPAddr("tcp", os.Args[1])
	if err != nil {
		fmt.Printf("Resolving host failed: %s", err.Error())
		os.Exit(1)
	}

	addr := p2p.NewNetAddress(p2p.PubKeyToID(ed25519.GenPrivKey().PubKey()), tcpAddr)

	p, err := transport.Dial(*addr, trans.PeerConfig{})
	if err != nil {
		output := Output{Status: "closed"}
		out, err := json.MarshalIndent(output, "", "  ")
		if err != nil {
			fmt.Printf("Marshal '%+v' into JSON failed: %s", output, err.Error())
		} else {
			fmt.Println(string(out))
		}

		os.Exit(1)
	}

	transport.Cleanup(p)

	info := p.NodeInfo().(p2p.DefaultNodeInfo)

	output := Output{
		Status:  "open",
		ID:      string(info.DefaultNodeID),
		Moniker: info.Moniker,
		Network: info.Network,
	}

	out, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		fmt.Printf("Marshal '%+v' into JSON failed: %s", output, err.Error())
		os.Exit(1)
	}

	fmt.Println(string(out))

	os.Exit(0)
}

type Output struct {
	Status  string
	ID      string
	Moniker string
	Network string
}

var defaultProtocolVersion = p2p.NewProtocolVersion(
	version.P2PProtocol,
	version.BlockProtocol,
	0,
)

func testNodeInfo(id p2p.ID, name string) p2p.NodeInfo {
	return testNodeInfoWithNetwork(id, name, "cudos-p2p-scan")
}

func testNodeInfoWithNetwork(id p2p.ID, name, network string) p2p.NodeInfo {
	return p2p.DefaultNodeInfo{
		ProtocolVersion: defaultProtocolVersion,
		DefaultNodeID:   id,
		ListenAddr:      "127.0.0.1:1337",
		Network:         network,
		Version:         "1.2.3-rc0-deadbeef",
		Channels:        []byte{0x01},
		Moniker:         name,
		Other: p2p.DefaultNodeInfoOther{
			TxIndex:    "on",
			RPCAddress: "127.0.0.1:31337",
		},
	}
}
