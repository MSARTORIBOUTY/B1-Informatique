curl -s https://ytrack.learn.ynov.com/assets/superhero/all.json | jq --arg HERO_ID "$HERO_ID" '.[] | select( .id == ($HERO_ID|tonumber))|.connections.relatives' | sed 's/"//g'
