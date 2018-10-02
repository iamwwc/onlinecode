#/usr/bin/bash

cd ../client-side
npm install && npm run build
cd ..
cp -R client-side/dist .
#rm -rf client
docker build --tag 0.0.1-alpine .