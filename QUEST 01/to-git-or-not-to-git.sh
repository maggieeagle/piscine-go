#!/bin/bash 
curl -s https://01.kood.tech/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"maggieeagle\"}}){id}}"}' | egrep -oE "[0-9]{1,}"