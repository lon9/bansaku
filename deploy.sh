#!/bin/sh

SCRIPT_DIR=`dirname $0`
GO_PATH=`which go`
NPM_PATH=`which npm`
NGINX=`which nginx`
if [ $GO_PATH = "" ]; then
  echo "Go is not installed."
  exit 1
fi

if [ $NGINX = "" ]; then
  echo "Nginx is not installed."
  exit 1
fi

echo "Stopping app server..."
killall zepher-bansaku

echo "Stopping nginx..."
sudo nginx -s stop

echo "Getting sources..."
go get -v -u github.com/Rompei/zepher-bansaku

cd $SCRIPT_DIR
echo "Building src..."
$GO_PATH build

cd $SCRIPT_DIR/workers/backup/
echo "installing node dependencies..."
$NPM_PATH install

echo "Copy to www directory"
if [ -e $HOME/www/zepher-bansaku ]; then
  rm -rf $HOME/www/zepher-bansaku
  mkdir -p $HOME/www/zepher-bansaku
else
  mkdir -p $HOME/www/zepher-bansaku
fi
cp -r $SCRIPT_DIR/* $HOME/www/zepher-bansaku/

echo "Starting nginx..."
sudo nginx

echo "Starting app server..."
cd $HOME/www/zepher-bansaku 
nohup ./zepher-bansaku &
