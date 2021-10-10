git clone https://github.com/desmos-labs/desmos
cd desmos
git clone https://github.com/desmos-labs/mainnet
git checkout v1.0.4
export DESMOS_P2P_LADDR=tcp://0.0.0.0:2000
export DESMOS_RPC_LADDR=tcp://127.0.0.1:2001
export DESMOS_GRPC_ADDRESS=127.0.0.1:2002
export DESMOS_API_ADDRESS=tcp://127.0.0.1:2003
export DESMOS_P2P_PERSISTENT_PEERS="c9709f150e9d09a2e5b7270257d7362f8f362e25@75.119.151.203:26656,ea5d1deee37fb04833efd730ecebc7e54a944f81@194.163.176.148:26656,fbb4b9b9d7cba4e539f8240df237ac65d03122ba@78.47.198.90:26656,ea5d1deee37fb04833efd730ecebc7e54a944f81@194.163.176.148:26656,f759a2ea33b68f2024eff5f751cc3ec68cb8b399@135.181.79.106:26656,650ec704b5a91c6488c579be2e0f777265d45f6e@dsms2.kysenpool.io:26656"
export DESMOS_NODE=tcp://127.0.0.1:2001
export DESMOS_SNAPSHOT_INTERVAL=1000
export DESMOS_P2P_MAX_NUM_OUTBOUND_PEERS=10
make install
desmos init notional --chain-id desmos-mainnet
cp mainnet/genesis.json ~/.desmos/config/
desmos start --pruning nothing
