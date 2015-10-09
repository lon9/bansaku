#!/bin/sh

SCRIPT_DIR=`dirname $0`
GO_PATH=`which go`
NGINX=`which nginx`
if [ $GO_PATH = "" ]; then
  echo "Go is not installed."
  exit 1
fi

if [ $NGINX = "" ]; then
  echo "Nginx is not installed."
  exit 1
fi

echo "Stopping nginx..."
sudo nginx -s stop

echo "Stopping app server..."
killall zepher-bansaku

cd $SCRIPT_DIR
pwd
echo "Building src..."
$GO_PATH build

echo "Copy to www directory"
if [ -e $HOME/www/zepher-bansaku ]; then
  mkdir -p $HOME/www/zepher-bansaku
fi
cp zepher-bansaku $HOME/www/zepher-bansaku/

echo "Starting app server..."
cd $HOME/www/zepher-bansaku 
./zepher-bansaku &

echo "Starting nginx..."
sudo nginx
