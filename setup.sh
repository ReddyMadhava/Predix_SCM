#!/bin/bash
UP_DOWN=$1

function printHelp {
        echo "Usage: ./setup.sh <up|down>"
}

function validateArgs {
        if [ -z "${UP_DOWN}" ]; then
                echo "up/down not mentioned"
                printHelp
                exit 1
        fi
}

validateArgs

#Create the netowk install node and up the app
if [ "${UP_DOWN}" == "up" ]; then
        docker-compose -f docker-compose.yaml up -d
        sleep 10
        npm install
        sleep 10
        echo
        echo "Start the application"
        ##start node application
        node product_shipment_temp_tracking.js
		node PayloadService.js
elif [ "${UP_DOWN}" == "down" ]; then ##Clean up the network
        docker-compose -f docker-compose.yaml down
        rm -rf keyValStore chaincodeID
else
        printHelp
        exit 1
fi

