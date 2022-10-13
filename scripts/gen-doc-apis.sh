#!/bin/sh

for dir in $(ls ".")
do
    # if file, then skip
    if [ -f $dir ]; then  
        continue
    fi

    # if it is v1alpha1, then skip
    if [ $dir = v1alpha1 ]; then
        continue
    fi

    # if file exists, then skip
    if [ -f $dir/"v1alpha1"/"doc.go" ]; then
        rm $dir/"v1alpha1"/"doc.go"
    fi

    export FORDOC=$dir".yandex-cloud.jet.crossplane.io"
    export service=$dir
    cat ../scripts/doc.go.tmpl | envsubst > $dir/"v1alpha1"/"doc.go"

    #cp "v1alpha1/doc.go" $dir/"v1alpha1"
    echo "COPIED doc.go to " $dir/"v1alpha1"
done