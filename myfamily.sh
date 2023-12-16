#!/bin/bash

HERO_ID=${HERO_ID:-1}

JSON_DATA=$(curl -s "https://zero.academie.one/assets/superhero/all.json")

echo "$JSON_DATA" | jq --arg hero_id "$HERO_ID" '.[] | select(.id == ($hero_id|tonumber)).connections.relatives' | tr -d '"'
