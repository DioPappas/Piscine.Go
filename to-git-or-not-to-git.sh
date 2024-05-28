#! /bin/bash

curl -o to-git-or-not-to-git.sh  https://platform.zone01.gr/assets/superhero/all.json
jq -r '.[] | select(.id == 170) | .name, .powerstats.power, .appearance.gender' data.json