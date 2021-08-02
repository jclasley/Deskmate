#!/bin/sh
/go/bin/server&
/go/bin/zendesk&

sleep 5

curl localhost:8080/handler/zendesk/connect \
    -d '{"url": "localhost:8090/"}'

# moved to last so that we hang
sleep infinity