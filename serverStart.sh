#!/bin/sh
/go/bin/server&

sleep 5

curl localhost:8080/handler/slack/connect
curl localhost:8080/handler/zendesk/connect

# moved to last so that we hang
/go/bin/zendesk