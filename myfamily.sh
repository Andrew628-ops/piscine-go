#!/bin/bash

if [ -z "$HERO_ID" = "1" ]; then
  echo "Error: HERO_ID environment variable is not set."
  exit 1
fi

relatives=$(curl -s https://acad.learn2earn.org.ng/assets/superhero/all.json | \ 
jq -r ".[] | select(.id == $HERO_ID) | .connections.relatives")


relatives_escaped=$(echo "$relatives" | sed ':a;N;$!ba;s/\n/\\n/g')


export HERO_ID=1