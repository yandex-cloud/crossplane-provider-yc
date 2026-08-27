#!/bin/sh

set -eu

directory=".work/terraform-provider-yandex/docs/resources"

if [ ! -d "$directory" ]; then
    exit 2
fi

for file in "$directory"/*.md; do
    resource_name="yandex_$(basename "$file" .md)"
    temporary_file="${file}.tmp"

    awk -v page_title="Yandex: ${resource_name}" -v resource_name="$resource_name" '
        NR == 1 && $0 == "---" {
            in_front_matter = 1
            print
            next
        }

        in_front_matter && /^subcategory:/ {
            has_subcategory = 1
        }

        in_front_matter && /^page_title:/ {
            has_page_title = 1
        }

        in_front_matter && /^description:/ {
            has_description = 1
        }

        in_front_matter && $0 == "---" {
            if (!has_subcategory) {
                print "subcategory: \"unknown\""
            }
            if (!has_page_title) {
                print "page_title: \"" page_title "\""
            }
            if (!has_description) {
                print "description: |-"
                print "  Manages the " resource_name " resource."
            }
            in_front_matter = 0
        }

        { print }

        END {
            if (in_front_matter) {
                exit 3
            }
        }
    ' "$file" > "$temporary_file"
    mv "$temporary_file" "$file"
done
