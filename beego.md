### bee 工具
* bee migrate

* bee generate

* bee run

* bee help [xxx]

* https://beego.me/docs/install/bee.md

* install golang.org/x/crypto/bcrypt
```
mkdir -p $GOPATH/src/golang.org/x/
cd $GOPATH/src/golang.org/x/
git clone https://github.com/golang/crypto
```

### bee 案例
* 生成模型: `bee generate model tag -fields="id:int,name:string,created_at:datetime,deleted_at:datetime"`

* 生成控制器: `bee generate controller Tag`

* migrate 暂时没有使用, 不如手动创建来得快


### 路由
* https://beego.me/docs/mvc/controller/router.md

* 路由分组 嵌套

### ab 测试 (含 header token)
```
ab -n 1000 -c 100 -H "Content-Type: application/json" -H "Authorization: xxx" http://localhost:8080/v1/tag/1
```


### TODO
* 如何热更新： https://beego.me/docs/module/grace.md

* cache 使用 

* vue 端