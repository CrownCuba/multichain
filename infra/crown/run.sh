#!/bin/bash

# Run the crown sandbox
python2.7 /root/sandbox.py -d=/root/ -f=/root/devnet-template.json.in start 

# Keep the container alive
while :
do
    sleep 10
done
