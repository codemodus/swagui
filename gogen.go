package swagui

//go:generate git clone -q https://github.com/swagger-api/swagger-ui

//go:generate git --git-dir=swagger-ui/.git --work-tree=swagger-ui checkout -q v1.1.15
//go:generate go-bindata -pkg=bindata1 -prefix=swagger-ui/dist -ignore=swagger-ui.js -o=bindata1/bindata.go swagger-ui/dist/...

//go:generate git --git-dir=swagger-ui/.git --work-tree=swagger-ui checkout -q master
//go:generate go-bindata -pkg=bindata2 -prefix=swagger-ui/dist -ignore=swagger-ui.js -o=bindata2/bindata.go swagger-ui/dist/...

//go:generate git clean -qdff swagger-ui
