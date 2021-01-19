# Generated 2021-01-19 11:42:11 by go-framework v1.2.1-1-ge87b524
MODULE="github.com/AtricoSoftware/peg-solitaire"
export OUTPUT_NAME="peg-solitaire"
TARGET_DIR=release
TARGET_PLATFORMS="darwin windows linux"

if [[ ! -z "$1" ]]
then
  VERSION=$1
else
  VERSION=$(git describe --tags --dirty)
fi

export CGO_ENABLED=0
export GOARCH="amd64"

# setup details
# built
BUILT_ON=$(date)
BUILT_BY=$(whoami)
# git
GIT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
GIT_COMMIT=$(git rev-parse HEAD)

DETAILS="{\"Built\":{\"On\":\"$BUILT_ON\", \"By\":\"$BUILT_BY\"},\"Git\":{ \"Repository\":\"$MODULE\",\"Branch\":\"$GIT_BRANCH\",\"Commit\":\"$GIT_COMMIT\"} }"
# Setup ldflags
LDFLAGS="-s -w"
LDFLAGS=$LDFLAGS" -X '$MODULE/pkg.Version=$VERSION'"
LDFLAGS=$LDFLAGS" -X '$MODULE/pkg.BuildDetails=$DETAILS'"


mkdir -p $TARGET_DIR
for GOOS in $TARGET_PLATFORMS; do
    export GOOS
    export EXT=""
    if [[ ${GOOS} == "windows" ]]
    then
      export EXT=".exe"
    fi
    export TARGET="$TARGET_DIR/$VERSION-$GOOS-$GOARCH"
    mkdir -p $TARGET
    go build -v -ldflags="$LDFLAGS" -o $TARGET/$OUTPUT_NAME$EXT

done

cd $TARGET_DIR
find . ! -path . -type d |  cut -d "/" -f2 | awk -v name="$OUTPUT_NAME" '{ print name "_" $1 ".zip -r ./" $1 "/"  }' | xargs -L1 zip -j
#find . ! -path . -type d | xargs -L1 rm -rf

