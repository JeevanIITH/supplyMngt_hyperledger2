export PATH=${PWD}/bin:${PWD}:$PATH
cryptogen generate --config=./organizations/cryptogen/crypto-config-producer.yaml --output="organizations"
cryptogen generate --config=./organizations/cryptogen/crypto-config-supplier.yaml --output="organizations"
cryptogen generate --config=./organizations/cryptogen/crypto-config-wholeseller.yaml --output="organizations"
cryptogen generate --config=./organizations/cryptogen/crypto-config-orderer1.yaml --output="organizations"
export COMPOSE_PROJECT_NAME=net
export DOCKER_SOCK=/var/run/docker.sock
IMAGE_TAG=latest docker-compose -f compose/compose-test-net.yaml -f compose/docker/docker-compose-test-net.yaml up