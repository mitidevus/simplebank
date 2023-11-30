#!/bin/sh

set -e

echo "start the app"
exec "$@" # It means: take all parameters passed to the script and run it

