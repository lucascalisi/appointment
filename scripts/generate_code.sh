#!/bin/bash

if_error_exit() {
    if [ $? -ne 0 ]
    then
        echo "ERROR: $1" >&2
        exit 1
    fi
}

# Generate server code
swagger generate server --spec ./http/swagger/swagger.yml --target ./http/
if_error_exit "could not generate server code"

# Generate client libraries
swagger generate client -f ./http/swagger/swagger.yml --target ./http/ --client-package appointment_client
echo "INFO: code generation finished OK"
