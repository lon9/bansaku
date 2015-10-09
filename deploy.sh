#!/bin/sh

SCRIPT_DIR=`dirname $0`
GO_PATH=`which go`
if [ $GO_PATH = "" ]; then
  echo "Go is not installed."
fi

cd $SCRIPT_DIR
pwd
echo "Build src"
$GO_PATH build

echo "Copy to www directory"
if [ -e $HOME/www/zepher-bansaku ]; then
  mkdir -p $HOME/www/zepher-bansaku
fi
cp zepher-bansaku $HOME/www/zepher-bansaku/

cd $HOME/www/zepher-bansaku 
#./zepher-bansaku
