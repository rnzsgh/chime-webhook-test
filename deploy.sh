#!/bin/bash

GOOS=linux go build

zip handler.zip ./chime-webhook-test

aws lambda update-function-code --function-name ChimeBot --zip-file fileb://handler.zip --publish
