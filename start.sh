#!/bin/bash

./mindlines-backend &

node ./server/index.mjs &
wait -n

exit $?
