#!/bin/bash
#ADDRESS=$1
#PK=$2
# Print setup
#echo "OOOOOOeEEEEEEEEE CROWN_ADDRESS=$ADDRESS"
#echo "PRIVATE_KEY=$PK"
# Prepare the directory
python2.7 /app/scripts/sandbox.py -d=/app/sandbox/ -b=/app/bin/ -f=/app/scripts/devnet-template.json.in prepare
python3.5 /app/scripts/addRpcUserRpcAllow.py /app/sandbox/
# Run the directory
python2.7 /app/scripts/sandbox.py -d=/app/sandbox/ -b=/app/bin/ -f=/app/scripts/devnet-template.json.in start

#/app/scripts/sb.sh.in cmd miner getinfo
#echo "==================="
##/app/bin/crown-cli generatetoaddress 101 $ADDRESS
#ifconfig

# Simulate mining
while true
do
    sleep 1
done
