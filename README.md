# CUDOS P2P SCAN

The tool checks if given node p2p port is responsive, by connecting and doing the tendermint handshake. If successful, the target node information will be displayed and process exit code will be zero. On fail the exit code will be non-zero.

## Build

### Local

```bash
go build
```

### Direct from GitHub repo

```bash
go install -v github.com/CudoVentures/cudos-p2p-scan@latest
```

### Binary

Binary will be written to ```~/go/bin/cudos-p2p-scan```

## Run

```bash
~/go/bin/cudos-p2p-scan <Target Host>:<Target port usually 26656>
```

Example output:

```json
{
  "Status": "open",
  "ID": "f93e129f120fd1de3e9d60d2bd376ae96af325dd",
  "Moniker": "cudos-seed-node-03",
  "Network": "cudos-testnet-public-2"
}
```
