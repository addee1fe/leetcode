#!/usr/bin/env bash

res=$(./solution.sh)

expected='the 4
is 3
sunny 2
day 1'

[[ ${res} == ${expected} ]] &&  echo "Passed" || echo "Failed"
