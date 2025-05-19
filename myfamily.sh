
#!/bin/bash

export HERO_ID=${HERO_ID:-70}

result=$(curl -s htts://acad.learn2earn.org.ng/assets/superhero/all.json | \
jq -r --arg id "$HERO_ID" '
.[] |
select(.id == ($id | tonumber)) | 
.connections.relatives
')

# format with literal \n
echo "$result" | sed ':a;N;$!ba;s/\n/\\n/g'