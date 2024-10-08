export CC_RUNTIME_LANGUAGE=golang
# export CC_SRC_PATH="../chaincode/fabcar/go/"
export CC_SRC_PATH="./chaincode/"
export VERSION=1

echo Vendoring Go dependencies ...
# pushd ../chaincode/fabcar/go
pushd ./chaincode
alias go=$HOME/go/bin/go1.19
export GO111MODULE=on go mod vendor
popd
echo Finished vendoring Go dependencies