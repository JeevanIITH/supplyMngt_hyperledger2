
Following commands are in commands folder

```
export PATH=${PWD}/bin:${PWD}:$PATH
cryptogen generate --config=./organizations/cryptogen/crypto-config-producer.yaml --output="organizations"
cryptogen generate --config=./organizations/cryptogen/crypto-config-supplier.yaml --output="organizations"
cryptogen generate --config=./organizations/cryptogen/crypto-config-wholeseller.yaml --output="organizations"
cryptogen generate --config=./organizations/cryptogen/crypto-config-orderer1.yaml --output="organizations"
export COMPOSE_PROJECT_NAME=net
export DOCKER_SOCK=/var/run/docker.sock
IMAGE_TAG=latest docker-compose -f compose/compose-test-net.yaml -f compose/docker/docker-compose-test-net.yaml -f compose/compose-ca.yaml up
```
Open another terminal and run .

```
export PATH=${PWD}/bin:${PWD}:$PATH
export FABRIC_CFG_PATH=${PWD}/configtx
export CHANNEL_NAME=channel1

configtxgen -profile AllOrg -outputBlock ./channel-artifacts/${CHANNEL_NAME}.block -channelID $CHANNEL_NAME
cp ./config/core.yaml ./configtx/.

export ORDERER_CA=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer1.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
export ORDERER_ADMIN_TLS_SIGN_CERT=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer1.example.com/tls/server.crt
export ORDERER_ADMIN_TLS_PRIVATE_KEY=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer1.example.com/tls/server.key
osnadmin channel join --channelID $CHANNEL_NAME --config-block ./channel-artifacts/${CHANNEL_NAME}.block -o localhost:7053 --ca-file "$ORDERER_CA" --client-cert "$ORDERER_ADMIN_TLS_SIGN_CERT" --client-key "$ORDERER_ADMIN_TLS_PRIVATE_KEY"

export ORDERER_CA=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer2.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
export ORDERER_ADMIN_TLS_SIGN_CERT=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer2.example.com/tls/server.crt
export ORDERER_ADMIN_TLS_PRIVATE_KEY=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer2.example.com/tls/server.key
osnadmin channel join --channelID $CHANNEL_NAME --config-block ./channel-artifacts/${CHANNEL_NAME}.block -o localhost:7153 --ca-file "$ORDERER_CA" --client-cert "$ORDERER_ADMIN_TLS_SIGN_CERT" --client-key "$ORDERER_ADMIN_TLS_PRIVATE_KEY"

export ORDERER_CA=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer3.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
export ORDERER_ADMIN_TLS_SIGN_CERT=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer3.example.com/tls/server.crt
export ORDERER_ADMIN_TLS_PRIVATE_KEY=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer3.example.com/tls/server.key
osnadmin channel join --channelID $CHANNEL_NAME --config-block ./channel-artifacts/${CHANNEL_NAME}.block -o localhost:7253 --ca-file "$ORDERER_CA" --client-cert "$ORDERER_ADMIN_TLS_SIGN_CERT" --client-key "$ORDERER_ADMIN_TLS_PRIVATE_KEY"


source ./scripts/setOrgPeerContext.sh 1
peer channel join -b ./channel-artifacts/${CHANNEL_NAME}.block

source ./scripts/setOrgPeerContext.sh 2
peer channel join -b ./channel-artifacts/${CHANNEL_NAME}.block

source ./scripts/setOrgPeerContext.sh 3
peer channel join -b ./channel-artifacts/${CHANNEL_NAME}.block

source ./scripts/setOrgPeerContext.sh 4
peer channel join -b ./channel-artifacts/${CHANNEL_NAME}.block

source ./scripts/setOrgPeerContext.sh 5
peer channel join -b ./channel-artifacts/${CHANNEL_NAME}.block

source ./scripts/setOrgPeerContext.sh 6
peer channel join -b ./channel-artifacts/${CHANNEL_NAME}.block



source ./scripts/setOrgPeerContext.sh 1
docker exec cli ./scripts/setAnchorPeer.sh 1 $CHANNEL_NAME

source ./scripts/setOrgPeerContext.sh 3
docker exec cli ./scripts/setAnchorPeer.sh 3 $CHANNEL_NAME

source ./scripts/setOrgPeerContext.sh 5
docker exec cli ./scripts/setAnchorPeer.sh 5 $CHANNEL_NAME





source ./scripts/setChaincodeContext.sh


source ./scripts/setOrgPeerContext.sh 1
peer lifecycle chaincode package supply1.tar.gz --path ${CC_SRC_PATH} --lang ${CC_RUNTIME_LANGUAGE} --label supply1_${VERSION}

peer lifecycle chaincode install supply1.tar.gz


source ./scripts/setOrgPeerContext.sh 2
peer lifecycle chaincode install supply1.tar.gz

source ./scripts/setOrgPeerContext.sh 3
peer lifecycle chaincode install supply1.tar.gz
source ./scripts/setOrgPeerContext.sh 4
peer lifecycle chaincode install supply1.tar.gz
source ./scripts/setOrgPeerContext.sh 5
peer lifecycle chaincode install supply1.tar.gz
source ./scripts/setOrgPeerContext.sh 6
peer lifecycle chaincode install supply1.tar.gz

peer lifecycle chaincode queryinstalled 2>&1 | tee outfile

source ./scripts/setPackageID.sh outfile

source ./scripts/setOrgPeerContext.sh 1
peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer1.example.com --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA --channelID $CHANNEL_NAME --name supply1 --version ${VERSION} --init-required --package-id ${PACKAGE_ID} --sequence ${VERSION}


peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL_NAME --name supply1 --version ${VERSION} --sequence ${VERSION} --output json --init-required
source ./scripts/setOrgPeerContext.sh 3
peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer1.example.com --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA --channelID $CHANNEL_NAME --name supply1 --version ${VERSION} --init-required --package-id ${PACKAGE_ID} --sequence ${VERSION}
source ./scripts/setOrgPeerContext.sh 5
peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer1.example.com --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA --channelID $CHANNEL_NAME --name supply1 --version ${VERSION} --init-required --package-id ${PACKAGE_ID} --sequence ${VERSION}


source ./scripts/setOrgPeerContext.sh 1
peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL_NAME --name supply1 --version ${VERSION} --sequence ${VERSION} --output json --init-required


source ./scripts/setPeerConnectionParam.sh 1 2 3 4 5 6
peer lifecycle chaincode commit -o localhost:7050 --ordererTLSHostnameOverride orderer1.example.com --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA --channelID $CHANNEL_NAME --name supply1 $PEER_CONN_PARAMS --version ${VERSION} --sequence ${VERSION} --init-required

source ./scripts/setOrgPeerContext.sh 1
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer1.example.com --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C channel1 -n supply1 $PEER_CONN_PARAMS --isInit -c '{"function":"InitLedger","Args":[]}'
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer1.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer1.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C channel1 -n supply1 $PEER_CONN_PARAMS -c '{"function":"GetAllAssets","Args":[]}'
```
Generating CCP files.
```
./organizations/ccp-generate.sh
```
Copy 
-organizations/peerOrganizations/producer.example.com/connection-producer.json
-organizations/peerOrganizations/supplier.example.com/connection-supplier.json
-organizations/peerOrganizations/wholeseller.example.com/connection-wholeseller.json
Paste in fabricSDK/certs folder.


Node js SDK .
```
cd fabricSDK
node app.js
```

APIs
Created by cryptogen , So import User identity, certs into wallets by sending following post req 
-POST : `localhost:3000:/api/registerUser` 
    {
      "userId": "User1",
      "affiliation": "test" 
    }
-POST : `localhost:3000:/api/product`
  {
    "productid" : "iPhone19",
    "name" : "iPhone18Pro",
    "description" : "New iPhone ",
    "manufacturingDate": "20-3-3000",
    "batchNumber" : "2341234",
    "userId" : "User1"
  }


