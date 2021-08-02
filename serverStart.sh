#!/bin/sh
/go/bin/zendesk&
/go/bin/server&

sleep 5

curl localhost:8080/handler/slack/connect
curl localhost:8080/handler/zendesk/connect

