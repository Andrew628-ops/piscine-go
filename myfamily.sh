#!/bin/bash
if [ -z "$HERO_ID=1" ]; then
  echo "Error: HERO_ID environment variable is not set."
  exit 1
fi
URL="https://acad.learn2earn.org.ng/assets/superhero/all.json"
DATA=$(curl -s "$URL")
RELATIVES=$(echo "$DATA" | jq -r --arg ID "$HERO_ID" '.[] | select(.id == ($ID | tonumber)) | .relatives')
if [ -z "$RELATIVES" ]; then
  echo "No relatives found for HERO_ID=$HERO_ID"
else
  echo "Relatives of HERO_ID=$HERO_ID:"
  echo "$RELATIVES"
fi
