# ⚙️ Hyperledger Fabric Code Setup

## Download Fabric Samples and Binaries

```bash
curl -sSLO https://raw.githubusercontent.com/hyperledger/fabric/main/scripts/install-fabric.sh
chmod +x install-fabric.sh
./install-fabric.sh d s b
```

This downloads:

- fabric-samples/
- Binaries: peer, orderer, configtxgen, etc.
- Docker images: fabric-peer, fabric-orderer, etc.

This script installs binaries to your current directory and also pulls Docker images.


## Set Up the Test Network

```bash
cd fabric-samples/test-network
```

Start the network with Certificate Authorities (CA), a channel (mychannel), and CouchDB state DB:

```bash
./network.sh up createChannel -ca -c mychannel -s couchdb
```

⚠️ Make sure Docker is running!


## Write Chaincode

1. Create a directory for your chaincode:

```bash
mkdir -p ../chaincode/mitigation
cd ../chaincode/mitigation
```

2. Initialize Go module:

```bash
go mod init mitigation
vi mitigation.go
```
- Press `I` symobl on the keyboard
- Copy the content from `chaincode/mitigation.go` and paste here.
- Press `esc`
- Press `Shift:` on the keyboard
- Type `wq!` to save the content of the file.


3. Install dependencies:

```bash
go mod tidy
go mod vendor
```

## Deploy the Chaincode

1. Go back to fabric-samples/test-network:

```bash
cd ../../test-network
```

2. Deploy the chaincode using:

```bash
./network.sh deployCC -ccn mitigation -ccp ../chaincode/mitigation -ccl go
```

- ccn: Chaincode name (e.g., mitigation)
- ccp: Path to chaincode folder
- ccl: Chaincode language (go)


## Environment Variables for Peer CLI

```bash
export PATH=${PWD}/../bin:$PATH
export FABRIC_CFG_PATH=$PWD/../config/

export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_TLS_ROOTCERT_FILE2=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051
export ORDERER_CA=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
```
   
## Test Chaincode Functions

1. Verify whether the chaincode is properly setup or not by running the following commands.

```bash
peer lifecycle chaincode querycommitted --channelID mychannel --name mitigation
peer lifecycle chaincode queryapproved --channelID mychannel --name mitigation
```

2. Initialize the ledger

```bash
peer chaincode invoke -o localhost:7050 --tls --cafile $ORDERER_CA -C mychannel -n mitigation --peerAddresses localhost:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses localhost:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE2 -c '{"Args":["InitLedger"]}'
```

3. Query active devices

```bash
peer chaincode query -C mychannel -n mitigation -c '{"Args":["QueryDevicesByState","active"]}'
```

4. Block malicious device

```bash
peer chaincode invoke -o localhost:7050 --tls --cafile $ORDERER_CA -C mychannel -n mitigation --peerAddresses localhost:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses localhost:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE2 -c '{"Args":["BlockDevice","192.168.0.11"]}'

peer chaincode invoke -o localhost:7050 --tls --cafile $ORDERER_CA -C mychannel -n mitigation --peerAddresses localhost:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses localhost:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE2 -c '{"Args":["BlockDevice","192.168.0.12"]}'

peer chaincode invoke -o localhost:7050 --tls --cafile $ORDERER_CA -C mychannel -n mitigation --peerAddresses localhost:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses localhost:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE2 -c '{"Args":["BlockDevice","192.168.0.13"]}'

peer chaincode invoke -o localhost:7050 --tls --cafile $ORDERER_CA -C mychannel -n mitigation --peerAddresses localhost:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses localhost:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE2 -c '{"Args":["BlockDevice","192.168.0.14"]}'

peer chaincode invoke -o localhost:7050 --tls --cafile $ORDERER_CA -C mychannel -n mitigation --peerAddresses localhost:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses localhost:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE2 -c '{"Args":["BlockDevice","192.168.0.4"]}'

peer chaincode invoke -o localhost:7050 --tls --cafile $ORDERER_CA -C mychannel -n mitigation --peerAddresses localhost:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses localhost:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE2 -c '{"Args":["BlockDevice","192.168.0.7"]}'

peer chaincode invoke -o localhost:7050 --tls --cafile $ORDERER_CA -C mychannel -n mitigation --peerAddresses localhost:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses localhost:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE2 -c '{"Args":["BlockDevice","192.168.0.3"]}'

peer chaincode invoke -o localhost:7050 --tls --cafile $ORDERER_CA -C mychannel -n mitigation --peerAddresses localhost:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses localhost:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE2 -c '{"Args":["BlockDevice","192.168.0.20"]}'

peer chaincode invoke -o localhost:7050 --tls --cafile $ORDERER_CA -C mychannel -n mitigation --peerAddresses localhost:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses localhost:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE2 -c '{"Args":["BlockDevice","192.168.0.18"]}'

peer chaincode invoke -o localhost:7050 --tls --cafile $ORDERER_CA -C mychannel -n mitigation --peerAddresses localhost:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses localhost:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE2 -c '{"Args":["BlockDevice","192.168.0.9"]}'
```

5. Query blocked devices

```bash
peer chaincode query -C mychannel -n mitigation -c '{"Args":["QueryDevicesByState","blocked"]}'
```

6. Delete the malicious device

```bash
peer chaincode invoke -o localhost:7050 --tls --cafile $ORDERER_CA -C mychannel -n mitigation --peerAddresses localhost:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses localhost:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE2 -c '{"Args":["DeleteDevice","192.168.0.11"]}'

peer chaincode invoke -o localhost:7050 --tls --cafile $ORDERER_CA -C mychannel -n mitigation --peerAddresses localhost:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses localhost:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE2 -c '{"Args":["DeleteDevice","192.168.0.12"]}'

peer chaincode invoke -o localhost:7050 --tls --cafile $ORDERER_CA -C mychannel -n mitigation --peerAddresses localhost:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses localhost:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE2 -c '{"Args":["DeleteDevice","192.168.0.13"]}'

peer chaincode invoke -o localhost:7050 --tls --cafile $ORDERER_CA -C mychannel -n mitigation --peerAddresses localhost:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses localhost:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE2 -c '{"Args":["DeleteDevice","192.168.0.14"]}'

peer chaincode invoke -o localhost:7050 --tls --cafile $ORDERER_CA -C mychannel -n mitigation --peerAddresses localhost:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses localhost:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE2 -c '{"Args":["DeleteDevice","192.168.0.4"]}'

peer chaincode invoke -o localhost:7050 --tls --cafile $ORDERER_CA -C mychannel -n mitigation --peerAddresses localhost:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses localhost:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE2 -c '{"Args":["DeleteDevice","192.168.0.7"]}'

peer chaincode invoke -o localhost:7050 --tls --cafile $ORDERER_CA -C mychannel -n mitigation --peerAddresses localhost:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses localhost:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE2 -c '{"Args":["DeleteDevice","192.168.0.3"]}'

peer chaincode invoke -o localhost:7050 --tls --cafile $ORDERER_CA -C mychannel -n mitigation --peerAddresses localhost:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses localhost:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE2 -c '{"Args":["DeleteDevice","192.168.0.20"]}'

peer chaincode invoke -o localhost:7050 --tls --cafile $ORDERER_CA -C mychannel -n mitigation --peerAddresses localhost:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses localhost:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE2 -c '{"Args":["DeleteDevice","192.168.0.18"]}'

peer chaincode invoke -o localhost:7050 --tls --cafile $ORDERER_CA -C mychannel -n mitigation --peerAddresses localhost:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses localhost:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE2 -c '{"Args":["DeleteDevice","192.168.0.9"]}'
```

7. Query deleted devices

```bash
peer chaincode query -C mychannel -n mitigation -c '{"Args":["QueryDevicesByState","deleted"]}'
```

8. Create device

```bash
peer chaincode invoke -o localhost:7050 --tls --cafile $ORDERER_CA -C mychannel -n mitigation --peerAddresses localhost:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses localhost:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE2 -c '{"Args":["RegisterDevice","192.168.1.111","test-sensor","Sensor"]}'
```

### Shut Down the Network

Once you're done testing:

```bash
./network.sh down
```