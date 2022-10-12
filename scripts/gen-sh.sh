#!/bin/sh

cd ../
gen-crd-api-reference-docs/gen-crd-api-reference-docs -config config/config_gen_doc.json \
-out-file $1 \
-api-dir $2 \
-template-dir gen-crd-api-reference-docs/template
cd apis