#!/usr/bin
# shellcheck shell=sh

mkdir -p /mnt/website
while true
 do
    VAR=$(shuf -i 3600-10800 -n 1)
    wget -O /mnt/website/index.html https://en.wikipedia.org/wiki/Special:Randoms
    sleep $VAR
 done
