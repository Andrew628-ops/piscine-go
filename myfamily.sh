 #!/bin/bash
 if [ -z "$HERO_ID" ] || [ "HERO_ID" -ne 1]; then
  echo "Error: HERO_ID environment variableis not set."
  exit 1
 fi

 relatives=$(curl -s https://acad.learn2earn.org.ng/assets/superhero/all.json | \
 jq -r ".[] | select(.id == $HERO_ID) | .connections.relatives")

 relatives_escaped=$(echo "$relatives" | sed ':a;N;$!ba;s/\n/\\n/g')

 if [ -z "$relatives" ] || [ "$relatives" = "null" ]; then
   echo "No relatives found for HERO_ID=$HERO_ID"
 else
   echo "$relatives_escaped"
 fi

 export HERO_ID=1
 