#!/bin/bash
ADDRESS=$1
PK=$2
# Print setup
echo "OOOOOOeEEEEEEEEE CROWN_ADDRESS=$ADDRESS"
echo "PRIVATE_KEY=$PK"
# Start
/app/bin/crownd
sleep 200


# Import the address
/app/bin/crown-cli importaddress $ADDRESS

/app/bin/crown-cli importprivkey $PK
