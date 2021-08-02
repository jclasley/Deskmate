#!/bin/sh
/go/bin/server&
/go/bin/zendesk&

now=$(date +"%T")
echo "Before sleep: $now"
sleep 5
now=$(date +"%T")
echo "After sleep: $now"

curl -f -v -L localhost:8080/handler/ # test any curl at all
curl -f -v -L localhost:8080/handler/zendesk/connect \
    -d '{"url": "http://localhost:8090/"}'

# moved to last so that we hang
sleep infinity