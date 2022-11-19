#!/bin/sh

if [ "$1" = "bootnode" ]
then
  bootnode -nodekey=/nodedata/boot.key -vmodule p2p=4
elif [ "$1" = "aqchain" ]
then
  # 初始化创世区块
  aqchain init /nodedata/genesis.json
  # 导入账户
  aqchain account import /nodedata/node.key --password /nodedata/password.txt
  # 获取账户地址
  account=$(aqchain account list)
  account=${account##*--}
  # 启动链
  aqchain --syncmode 'full' --port 30303 --rpc --rpcaddr '0.0.0.0' --rpcport 8545 --rpccorsdomain="*" --rpcapi 'personal,db,eth,net,web3,txpool,miner,net' --ws --wsaddr '0.0.0.0' --wsport 8546 --wsorigins '*' --wsapi 'personal,db,eth,net,web3,txpool,miner,net' --bootnodes 'enode://5847e67fae8f77aa9965e794e304429278039b741c3859219af15b6615450d5d08dc92a8d315ef62a14d86706ed90e6acce0e6d6c72982f050916dc91c182212@[172.20.0.10]:30301'  --networkid 1000 -unlock "$account" --password /nodedata/password.txt --vmodule p2p=4 --nodekey=/nodedata/node.key --mine
else
  echo "aqchain"
fi

