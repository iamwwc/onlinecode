#/usr/bin/bash

cd ../client
npm install && npm run build
cd ..
cp -R client/dist .
#rm -rf client
docker build --tag 0.0.1-alpine .