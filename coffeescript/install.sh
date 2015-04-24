#!/bin/bash

#sudo apt-get install -y coffeescript
sudo apt-get install -y npm
sudo apt-get install -y editorconfig
mkdir -p ~/bin
sudo ln -s /usr/bin/nodejs /usr/bin/node

npm install coffee-script
ln -s ~/node_modules/.bin/coffee ~/bin/coffee
ln -s ~/node_modules/.bin/cake ~/bin/cake

npm install coffeelint
ln -s ~/node_modules/.bin/coffeelint ~/bin/coffeelint

npm install gulp
ln -s ~/node_modules/.bin/gulp ~/bin/gulp
npm install gulp-util
npm install gulp-coffee
npm install gulp-coffeelint

# generate coffeelint config with
# coffeelint --makeconfig coffeelint.json
# change indent from 2 to 4
