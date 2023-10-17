#!/bin/bash
POSTID=$(go run main.go durov); if [ ! -z "${POSTID}" ]; then echo "const LAST_POST_ID = ${POSTID};" > last_post_id.js; fi;