#!/bin/sh
/go/bin/server&
/go/bin/zendesk&

now=$(date +"%T")
echo "Before sleep: $now"
sleep 5
now=$(date +"%T")
echo "After sleep: $now"

curl -f -sS -L localhost:8080/handler/zendesk/connect \
    -d '{"url": "circleci.zendesk.com/"}'

# moved to last so that we hang
sleep infinity