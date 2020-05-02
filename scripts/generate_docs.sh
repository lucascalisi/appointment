#!/bin/bash

swagger-codegen generate -l html2 -i ./http/swagger/swagger.yml -o docs/http_api
