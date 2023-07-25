# Cobra 脚手架
## 根命令使用
```bash 
go run main.go --help    
```
![img.png](images/img.png)
```bash 
go run main.go --a 18 --b h --c 19 --d hello --d tony --e 98 
```
![img_1.png](images/img_1.png)
```bash
go run main.go version
```
![img_4.png](images/img_4.png)

## 子命令使用
```bash
go run main.go subcmd --help
```
![img_2.png](images/img_2.png)
```bash
go run main.go subcmd --s xiaomi 
```
![img_3.png](images/img_3.png)

## 交叉编译脚本
编译安装在本地
```bash
sh .cross_compile.sh install
```
删除本地安装
```bash
sh .cross_compile.sh uninstall
```
编译多个架构
```bash
sh .cross_compile.sh 
```

