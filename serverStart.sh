#!/bin/sh
/go/bin/zendesk&
/go/bin/server

# somehowWait 3
# need to hit connect endpoints of each server wiht a `curl`

# curl localhost:8080/handler/slack/connect
# curl localhost:8080/handler/zendesk/connect
