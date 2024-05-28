#! /bin/bash

jq -r '.[] | select(.id == 170) | .name, .powerstats.power, .appearance.gender' data.json
