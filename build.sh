#!/usr/bin/env bash

#######################
# Build and Test script
# Copyright 2020 <rytsh@devusage.com>
# All rights reserved. MIT license
#######################

BASE_DIR="$(realpath $(dirname "$0"))"
OUTPUT_FOLDER="${BASE_DIR}/out"
VERSION="$(git describe --tags --abbrev=0)"
MAINGO="${BASE_DIR}/cmd/indigo/indigo.go"
FLAG_V="indigo/internal/common.Version=${VERSION}"

function usage() {
    cat - <<EOF
Build script for golang
Set PLATFORMS env variable to export
PLATFORMS="windows linux darwin" is default
Usage: $0 <OPTIONS>
OPTIONS:
  --run
    Run for dev
  --build
    Build application to various platforms
  --clean
    Clean output folder

  --test
    Test code
    --cover
      Export coverage of test

  --coveralls
    Run coveralls tool

  -h, --help
    This help page
EOF
}

#######################
# Functions
function build() {
    echo "> Buiding indigo for ${1}"
    OUTPUT_FOLDER_IN=${OUTPUT_FOLDER}/${1}
    mkdir -p ${OUTPUT_FOLDER_IN}
    CGO_ENABLED=0 GOOS=${1} GOARCH=amd64 go build -trimpath -ldflags="-s -w -X ${FLAG_V}" -o ${OUTPUT_FOLDER_IN} ${MAINGO}
    (
	    cd ${OUTPUT_FOLDER_IN}
        if [[ "${1}" == "windows" ]]; then
            zip ../indigo-${1}-amd64-${VERSION}.zip *
        else
            tar czf ../indigo-${1}-amd64-${VERSION}.tar.gz *
        fi
    )
}
#######################

#######################
# Run

if [[ -z ${PLATFORMS} ]]; then
    # set default platforms
    PLATFORMS="windows linux darwin"
fi

while [[ "$#" -gt 0 ]]; do
    case "${1}" in
    --run)
        go run cmd/indigo/indigo.go test/ex.json
        exit 0
        ;;
    --build)
        BUILD="Y"
        shift 1
        ;;
    --clean)
        CLEAN="Y"
        shift 1
        ;;
    --test)
        TEST="Y"
        shift 1
        ;;
    --cover)
        COVER="Y"
        shift 1
        ;;
    --coveralls)
        COVERALLS="Y"
        TEST="Y"
        COVER="Y"
        shift 1
        ;;
    -h | --help)
        usage
        exit 0
        ;;
    *)
        usage >&2
        exit 1
        ;;
    esac
done

cd $BASE_DIR

# Clean output folder
if [[ "${CLEAN}" == "Y" ]]; then
    echo "> Cleaning builded files..."
	rm -rf ${OUTPUT_FOLDER}/* 2> /dev/null
fi

# Test
if [[ "${TEST}" == "Y" ]]; then
    echo "> Test started"
    [[ "${COVER}" == "Y" ]] && COVERAGE="-coverprofile=${OUTPUT_FOLDER}/cover.out"
	CGO_ENABLED=0 go test -v ./... ${COVERAGE}
fi

# Send to coveralls
if [[ "${COVERALLS}" == "Y" ]]; then
    echo "> Coveralls Test started"
    goveralls -coverprofile=./out/cover.out -service=drone.io
fi

# Build packages
if [[ "${BUILD}" == "Y" ]]; then
    mkdir -p ${OUTPUT_FOLDER}
    for PLATFORM in ${PLATFORMS}; do
        build ${PLATFORM}
    done
fi

###############
# END
