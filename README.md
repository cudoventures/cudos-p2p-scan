# CUDOS P2P SCAN

The tool checks if given node p2p port is responsive, by connecting, doing the tendermint handshake. If successful, the target node information will be displayed and process exit code will be zero. On fail the exit code will be non-zero.

## Build

```
cd apps/cudos-p2p-scan
go build
```


## Run

```
./cudos-p2p-scan TargetHost:26656
```

Example output:

```
{
  "Status": "open",
  "ID": "f93e129f120fd1de3e9d60d2bd376ae96af325dd",
  "Moniker": "cudos-seed-node-03",
  "Network": "cudos-testnet-public-2"
}
```