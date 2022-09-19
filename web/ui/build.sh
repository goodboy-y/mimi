#！/bin/bash

npm install --registry=http://registry.npmmirror.com
npm run build

rm -rf ../src/main/resources/static/assets
cp dist/index.html ../src/main/resources/templates/home
cp -r dist/assets ../src/main/resources/static/