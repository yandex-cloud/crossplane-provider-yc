#!/bin/sh

if [ -d "docs" ]; then  
    mkdir "docs"
fi

for dir in $(ls ".")
do
    # if file, then skip
    if [ -f $dir ]; then  
        continue
    fi

    if [ -f "../docs/"$dir"_doc.html" ]; then
        continue
    fi

    if [ $dir = "v1alpha1" ]; then
        ../scripts/gen-sh.sh "./docs/"$dir"_doc.html" ./apis/v1alpha1
        continue
    fi

    ../scripts/gen-sh.sh "./docs/"$dir"_doc.html" ./apis/$dir/v1alpha1
    #echo "GENERATED " "docs/"$dir"_doc.html"
done