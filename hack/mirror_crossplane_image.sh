#!/usr/bin/env bash

set -euo pipefail

if [[ $# -ne 5 ]]; then
  echo "Usage: $0 <helm> <yq> <chart> <target-registry> <values-file>" >&2
  exit 2
fi

helm_bin="$1"
yq_bin="$2"
chart="$3"
target_registry="${4%/}"
values_file="$5"
docker_bin="${DOCKER:-docker}"

if [[ -z "${target_registry}" ]]; then
  echo "Target registry is empty" >&2
  exit 1
fi

chart_metadata="$("${helm_bin}" show chart "${chart}")"
chart_values="$("${helm_bin}" show values "${chart}")"

image_repository="$(printf '%s\n' "${chart_values}" | "${yq_bin}" eval -r '.image.repository // ""' -)"
image_tag="$(printf '%s\n' "${chart_values}" | "${yq_bin}" eval -r '.image.tag // ""' -)"
ignore_tag="$(printf '%s\n' "${chart_values}" | "${yq_bin}" eval -r '.image.ignoreTag // false' -)"
app_version="$(printf '%s\n' "${chart_metadata}" | "${yq_bin}" eval -r '.appVersion // ""' -)"

if [[ -z "${image_repository}" ]]; then
  echo "Chart ${chart} does not define image.repository" >&2
  exit 1
fi

if [[ "${ignore_tag}" == "true" ]]; then
  echo "Chart ${chart} uses image.ignoreTag=true, which is not supported by the E2E image mirror" >&2
  exit 1
fi

if [[ -z "${image_tag}" ]]; then
  if [[ -z "${app_version}" ]]; then
    echo "Chart ${chart} defines neither image.tag nor appVersion" >&2
    exit 1
  fi
  image_tag="v${app_version}"
fi

source_image="${image_repository}:${image_tag}"
target_repository="${target_registry}/crossplane"
target_image="${target_repository}:${image_tag}"

echo "Mirroring Crossplane image ${source_image} to ${target_image}"
"${docker_bin}" pull --platform linux/amd64 "${source_image}"
"${docker_bin}" tag "${source_image}" "${target_image}"
"${docker_bin}" push "${target_image}"

mkdir -p "$(dirname "${values_file}")"
TARGET_REPOSITORY="${target_repository}" TARGET_TAG="${image_tag}" \
  "${yq_bin}" eval -n \
  '.image.repository = strenv(TARGET_REPOSITORY) | .image.tag = strenv(TARGET_TAG) | .image.ignoreTag = false' \
  > "${values_file}"

echo "Wrote Crossplane image override to ${values_file}"
