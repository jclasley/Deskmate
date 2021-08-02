#!/bin/sh
/go/bin/server&

sleep 5

curl localhost:8080/handler/slack/connect
curl localhost:8080/handler/zendesk/connect \
    -d "{url: 'localhost:8090'}"

# moved to last so that we hang
/go/bin/zendesk
