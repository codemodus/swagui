package swagui

//go:generate git clone -q https://github.com/swagger-api/swagger-ui

//o:generate git --git-dir=swagger-ui/.git --work-tree=swagger-ui checkout -q v1.1.15
//o:generate withdraw swagger-ui/dist/swagger-ui.js
//o:generate go-bindata -pkg=bindata1 -prefix=swagger-ui/dist -o=bindata1/bindata.go swagger-ui/dist/...

//o:generate git --git-dir=swagger-ui/.git --work-tree=swagger-ui checkout -q v2.2.10
//o:generate withdraw swagger-ui/dist/swagger-ui.js
//o:generate go-bindata -pkg=bindata2 -prefix=swagger-ui/dist -o=bindata2/bindata.go swagger-ui/dist/...

//go:generate git --git-dir=swagger-ui/.git --work-tree=swagger-ui checkout -q v3.21.0
//go:generate withdraw swagger-ui/dist/swagger-ui-bundle.js.map swagger-ui/dist/swagger-ui.css.map swagger-ui/dist/swagger-ui.js.map swagger-ui/dist/swagger-ui-standalone-preset.js.map
//go:generate go-bindata -pkg=bindata3 -prefix=swagger-ui/dist -o=bindata3/bindata.go swagger-ui/dist/...

//go:generate withdraw swagger-ui/flavors/swagger-ui-react/dist/.npmrc
//go:generate git clean -qdff swagger-ui
