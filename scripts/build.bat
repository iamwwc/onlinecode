cd ../client-side
npm install && npm run build
cd ..
xcopy client-side\dist\*.* /YESQ .
REM rm client-side /s
docker build --tag 0.0.1-alpine3.7 .