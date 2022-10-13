#!/bin/sh

cd ../
if ! [ -f gen-crd-api-reference-docs/gen-crd-api-reference-docs ]; then  
    cd gen-crd-api-reference-docs/
    go build
    cd ../
fi

gen-crd-api-reference-docs/gen-crd-api-reference-docs -config scripts/gen-apidocs-config.json \
-out-file $1 \
-api-dir $2 \
-template-dir gen-crd-api-reference-docs/template
cd apis