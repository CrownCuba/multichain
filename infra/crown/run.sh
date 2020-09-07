#!/bin/bash
ADDRESS=$1
# Print setup
echo "CROWN_ADDRESS=$ADDRESS"

# Start
/app/bin/crownd
sleep 100


# Import the address
/app/bin/crown-cli importaddress $ADDRESS



