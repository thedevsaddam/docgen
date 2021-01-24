#! /bin/bash

exec_curl(){
  echo "Found docgen latest version: $VERSION"
  echo "Download may take few minutes depending on your internet speed"
  echo "Downloading docgen to $2"
  curl -L --silent --connect-timeout 30 --retry-delay 5 --retry 5 -o "$2" "$1"
}

OS=`uname`
ARCH=`uname -m`
VERSION=$1
URL=https://github.com/thedevsaddam/docgen
TARGET=/usr/local/bin/docgen
MESSAGE_START="Installing docgen"
MESSAGE_END="Installation complete"

if [ "$VERSION" == "" ]; then
  LATEST_RELEASE=$(curl -L -s -H 'Accept: application/json' $URL/releases/latest)
  VERSION=$(echo $LATEST_RELEASE | sed -e 's/.*"tag_name":"\([^"]*\)".*/\1/')
fi

if [ "$OS" == "Darwin" ]; then
    exec_curl $URL/releases/download/$VERSION/mac_amd64 $TARGET
    echo "$MESSAGE_START"
    chmod +x $TARGET
    echo "$MESSAGE_END"
    docgen
elif [ "$OS" == "Linux" ]; then
  if [ "$ARCH" == "x86_64" ]; then
    exec_curl $URL/releases/download/$VERSION/linux_amd64 $TARGET
    echo "$MESSAGE_START"
    chmod +x $TARGET
    echo "$MESSAGE_END"
    docgen
  fi

  if [ "$ARCH" == "i368" ]; then
    exec_curl $URL/releases/download/$VERSION/linux_386 $TARGET
    chmod +x $TARGET
    docgen
  fi
fi

