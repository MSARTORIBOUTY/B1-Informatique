
#!/bin/bash
curl -s https://ytrack.learn.ynov.com/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{githubLogin:{_eq:\"MSARTORIBOUTY\"}}){id}}"}' | jq ' .data.user[0].id '
