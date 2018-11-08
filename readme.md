## 步骤
```
cd config
cp config.yaml.local config.yaml
cd ..
go run *.go
```
## govendor
```
govendor init
govendor add +external
```
## 添加包
```
govendor add +outside
```
## 同步包
```
govendor sync
```
## 添加额外包
```
govendor add gopkg.in/yaml.v2
```