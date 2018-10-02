#/usr/bin/bash

cd ../client-side
npm install && npm run build
cd ..
cp client-side/dist/* . -R
#rm -Rf client-side
docker build --tag 0.0.1-alpine3.7 .