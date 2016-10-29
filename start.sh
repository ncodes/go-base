#!/usr/local/bin/fish

export APP_ENV="development"					
export PORT=4001

# refresh -c ./refresh.conf
go build -v; and ./base --port 4001