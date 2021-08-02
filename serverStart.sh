#!/bin/sh
/go/bin/server&
/go/bin/zendesk&

sleep 5

curl -f -sS -L localhost:8080/handler/zendesk/connect \
    -d '{"url": "circleci.zendesk.com/"}'

# moved to last so that we hang
sleep infinity