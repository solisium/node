#!/bin/bash

set -e

exec /node/build/myst/myst \
  --config-dir=/etc/mysterium-node \
  --script-dir=/etc/mysterium-node \
  --log-dir= --data-dir=/var/lib/mysterium-node \
  --runtime-dir=/var/run/mysterium-node \
  --tequilapi.address=0.0.0.0 \
  --discovery.ping=10s \
  --discovery.fetch=10s \
  --log-level=debug \
  --payments.mystscaddress=0x4D1d104AbD4F4351a0c51bE1e9CA0750BbCa1665 \
  --ip-detector=http://ipify:3000/?format=json \
  --location.type=manual \
  --location.country=e2e-land \
  --broker-address=broker \
  --firewall.protected.networks= \
  --api.address=http://discovery:8080/api/v3 \
  --ether.client.rpc=http://ganache:8545 \
  --transactor.registry-address=0xbe180c8CA53F280C7BE8669596fF7939d933AA10 \
  --transactor.channel-implementation=0x599d43715DF3070f83355D9D90AE62c159E62A75 \
  --hermes.hermes-id=0xf2e2c77D2e7207d8341106E6EfA469d1940FD0d8 \
  --transactor.address=http://transactor:8888/api/v1 \
  --quality.address=http://morqa:8085/api/v2 \
  --keystore.lightweight service \
  --agreed-terms-and-conditions \
  --identity=0xd1a23227bd5ad77f36ba62badcb78a410a1db6c5 \
  --identity.passphrase=localprovider \
  --openvpn.port=3000 \
  openvpn,wireguard
