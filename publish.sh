#!/usr/bin/env bash

#######################
# Publish Page
# Eray Ates <rytsh@devusage.com>
# MIT license
#######################

BASE_DIR="$(realpath $(dirname "$0"))"
cd $BASE_DIR

read -p "> Do you want to publish [y/N]? " answer
case ${answer:0:1} in
    y|Y)
        PUBLISH_PAGE="Y"
    ;;
    *)
        echo "> Ok than, byee.."
    ;;
esac

# Publish Page
if [[ "${PUBLISH_PAGE}" == "Y" ]]; then
    (
        echo "> Publish page started with ${VERSION}"
        # echo "LATEST_VERSION=${VERSION}" > .env
        npm install && npm run build && npm run publish
    )
fi
