package swagui

//go:generate git clone -q https://github.com/swagger-api/swagger-ui

//NOgo:generate git --git-dir=swagger-ui/.git --work-tree=swagger-ui checkout -q v1.1.15
//NOgo:generate withdraw swagger-ui/dist/swagger-ui.js
//NOgo:generate go-bindata -pkg=suidata1 -prefix=swagger-ui/dist -o=suidata1/bindata.go swagger-ui/dist/...

//NOgo:generate git --git-dir=swagger-ui/.git --work-tree=swagger-ui checkout -q v2.2.10
//NOgo:generate withdraw swagger-ui/dist/swagger-ui.js
//NOgo:generate go-bindata -pkg=suidata2 -prefix=swagger-ui/dist -o=suidata2/bindata.go swagger-ui/dist/...

//go:generate git --git-dir=swagger-ui/.git --work-tree=swagger-ui checkout -q v3.21.0
//go:generate withdraw swagger-ui/dist/swagger-ui-bundle.js.map swagger-ui/dist/swagger-ui.css.map swagger-ui/dist/swagger-ui.js.map swagger-ui/dist/swagger-ui-standalone-preset.js.map
//go:generate go-bindata -pkg=suidata3 -prefix=swagger-ui/dist -o=suidata3/bindata.go swagger-ui/dist/...

//go:generate withdraw swagger-ui/flavors/swagger-ui-react/dist/.npmrc
//go:generate git clean -qdff swagger-ui
