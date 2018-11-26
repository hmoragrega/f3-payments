#!/bin/sh

host=$1
port=$2
timeout=$3

echo "Waiting for DB to be available"

return=0
while ! nc -z $host $port; do   
  retries=$((retries - 1))
  if [ "$retries" -eq 0 ]; then
    return=-1
    break;
  fi
  echo "."
  sleep 1
done

exit $success