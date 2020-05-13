#! /bin/bash

exec_curl(){
  echo "Downloading docgen to $2"
  curl --connect-timeout 30 --retry-delay 5 --retry 5 -o "$2" "$1"
}

OS=`uname`
ARCH=`uname -m`
VERSION=$1
URL=https://raw.githubusercontent.com/thedevsaddam/docgen-bin/master
TARGET=/usr/local/bin/docgen
MESSAGE_START="Installing docgen"
MESSAGE_END="Installation complete"

if [ "$VERSION" == "" ]; then
  VERSION="latest"
fi

if [ "$OS" == "Darwin" ]; then
    exec_curl $URL/$VERSION/mac_amd64 $TARGET
    echo "$MESSAGE_START"
    chmod +x $TARGET
    echo "$MESSAGE_END"
    docgen
elif [ "$OS" == "Linux" ]; then
  if [ "$ARCH" == "x86_64" ]; then
    exec_curl $URL/$VERSION/linux_amd64 $TARGET
    echo "$MESSAGE_START"
    chmod +x $TARGET
    echo "$MESSAGE_END"
    docgen
  fi

  if [ "$ARCH" == "i368" ]; then
    exec_curl $URL/$VERSION/linux_386 $TARGET
    chmod +x $TARGET
    docgen
  fi
fi

