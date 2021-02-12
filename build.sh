#!/usr/bin/env bash

#######################
# Build and Test script
# Eray Ates <rytsh@devusage.com>
# MIT license
#######################

BASE_DIR="$(realpath $(dirname "$0"))"
cd $BASE_DIR

OUTPUT_FOLDER="${BASE_DIR}/out"

# Droneio tag get
if [[ -z ${DRONE_TAG} ]]; then
    VERSION="$(git describe --tags --abbrev=0)"
else
    VERSION=${DRONE_TAG}
fi

MAINGO="${BASE_DIR}/cmd/indigo/indigo.go"
FLAG_V="indigo/internal/common.Version=${VERSION}"

function usage() {
    cat - <<EOF
Build script for golang
Set PLATFORMS env variable to export
PLATFORMS="windows:amd64,linux:amd64,darwin:amd64" is default
Usage: $0 <OPTIONS>
OPTIONS:
  --run
    Run for dev
  --build
    Build application to various platforms
    --pack
      Pack output
  --clean
    Clean output folder

  --test
    Test code
    --cover
      Export coverage of test

  --build-context
    Build context for image build
  --build-docker
    Build docker image and publish to docker hub
    Auto enabled build-context

  --changelog
    Get change log

  --coveralls
    Run coveralls tool

  -h, --help
    This help page
EOF
}

#######################
# Functions
function build() {
    echo "> Buiding indigo for ${1}-${2}"
    OUTPUT_FOLDER_IN=${OUTPUT_FOLDER}/${1}
    mkdir -p ${OUTPUT_FOLDER_IN}
    CGO_ENABLED=0 GOOS=${1} GOARCH=${2} go build -trimpath -ldflags="-s -w -X ${FLAG_V}" -o ${OUTPUT_FOLDER_IN} ${MAINGO}
    if [[ "${PACK}" == "Y" ]]; then
        (
            cd ${OUTPUT_FOLDER_IN}
            if [[ "${1}" == "windows" ]]; then
                zip ../indigo-${1}-${2}-${VERSION}.zip *
            else
                tar czf ../indigo-${1}-${2}-${VERSION}.tar.gz *
            fi
        )
    fi
}
#######################

#######################
# Run

if [[ -z ${PLATFORMS} ]]; then
    # set default platforms
    PLATFORMS="windows:amd64,linux:amd64,darwin:amd64"
fi

while [[ "$#" -gt 0 ]]; do
    case "${1}" in
    --run)
        shift 1
        set -x
        go run cmd/indigo/indigo.go ${*}
        set +x
        exit 0
        ;;
    --build)
        BUILD="Y"
        shift 1
        ;;
    --pack)
        PACK="Y"
        shift 1
        ;;
    --clean)
        CLEAN="Y"
        shift 1
        ;;
    --build-context)
        CONTEXT_BUILD="Y"
        shift 1
        ;;
    --build-docker)
        BUILD_DOCKER="Y"
        CONTEXT_BUILD="Y"
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
    --changelog)
        CHANGELOG="Y"
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
    set -e
    mkdir -p ${OUTPUT_FOLDER}
    IFS=',' read -ra PLATFORMS_ARR <<< $(echo ${PLATFORMS} | tr -d ' ')
    for PLATFORM_A in "${PLATFORMS_ARR[@]}"; do
        PLATFORM=$(echo ${PLATFORM_A} | cut -d ':' -f 1)
        ARCHS=$(echo ${PLATFORM_A} | cut -d ':' -f 2)
        IFS='-' read -ra ARCHS_ARR <<< ${ARCHS}
        for ARCH in ${ARCHS_ARR[@]}; do
            build ${PLATFORM} ${ARCH}
        done
    done
    set +e
fi

# Context for docker build
if [[ "${CONTEXT_BUILD}" == "Y" ]]; then
    echo "> Context building ${VERSION}"
    echo -n ${VERSION} > "${OUTPUT_FOLDER}/version"
    tar -czf "${OUTPUT_FOLDER}/context.tar.gz" out/linux/indigo ci/run/Dockerfile
fi

# Publish Docker
if [[ "${BUILD_DOCKER}" == "Y" ]]; then
    echo "> Build docker ${VERSION}"
    docker build -t ryts/indigo:${VERSION} -f ci/run/Dockerfile - < "${OUTPUT_FOLDER}/context.tar.gz"
    docker tag ryts/indigo:${VERSION} ryts/indigo:latest
fi

if [[ "${CHANGELOG}" == "Y" ]]; then
    echo "# Changelog"

    while IFS= read -r LOG; do
        # take action on $line #
        LOG_HASH=$(echo ${LOG} | cut -d ' ' -f 1)
        LOG_MESSAGE=$(echo ${LOG} | cut -d ' ' -f 2-)
        echo '`'${LOG_HASH}'`' ${LOG_MESSAGE}
    done <<< $(git tag -l --sort=-refname | awk 'NR==2' | xargs -I {} git log --oneline --no-decorate {}..${VERSION})

    echo ""
    echo "# Docker Images"
    echo '`'docker pull ryts/indigo:${VERSION}'`'
    echo '`'docker pull ryts/indigo:latest'`'

fi
###############
# END
