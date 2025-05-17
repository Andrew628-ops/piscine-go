#!/bin/bash
curl -s https://acad.learn2earn.org.ng/assets/superhero/all.json \ | jq '.[] | select(.id == 70).name'


