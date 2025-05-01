#!/bin/bash
echo '==============================================================='
echo 'space usage before clenning:'
df -h /
echo
docker system df
echo

echo 'cleaning ...'
docker builder prune -f
echo

echo '==============================================================='
echo 'space usage after clenning: '
docker system df
echo
df -h /

