#!/bin/sh

gen-crd-api-reference-docs/gen-crd-api-reference-docs -config gen-crd-api-reference-docs/example-config.json \
-out-file $1 \
-api-dir $2 \
-template-dir gen-crd-api-reference-docs/template