#!/bin/bash

ENDPOINT="http://127.0.0.1:8888"
CTYPE="Content-Type:application/json; charset=utf-8"

curl -v -X POST -H "${CTYPE}" -d "{\"comment\":\"HelloWorld #${RANDOM}\",\"state\":0}" "${ENDPOINT}/todos"
