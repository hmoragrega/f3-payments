#!/bin/bash

HOST=$1
PORT=$2
TIMEOUT=$3

success=0 
retries=$3
checkDB() {
   cat < /dev/tcp/$HOST/$PORT
   success=$?
}

echo "Waiting for DB to be available"

checkDB
while [ "$success" -ne 0 ]; do
    retries=$((retries - 1))
    if [ "$retries" -eq 0 ]; then
        echo "Timeout reached"
        break;
    fi

    echo "."
    sleep 1
    checkDB
done

echo "Last check $success"

exit $success