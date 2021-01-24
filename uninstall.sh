#! /bin/bash

TARGET=/usr/local/bin/docgen
MESSAGE_START="Removing docgen"
MESSAGE_END="Docgen removed"

echo "$MESSAGE_START"
rm $TARGET
echo "$MESSAGE_END"

