<!--
 * @Github: https://github.com/shijf
 * @Author: shijf
 * @Date: 2020-12-26 09:19:23
 * @LastEditTime: 2020-12-26 09:29:42
 * @LastEditors: shijf
 * @FilePath: /golang.learnku.com/note.md
 * @Description: 
-->
# 注意点

## linux(wsl)下配置环境变量
```shell
export GOROOT=/usr/local/go # 安装go的地方
export GOPATH=$HOME/go # go env 下查看 GoPath 
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin ## 导出环境变量
```

## 本地查看 godoc
`godoc -http=:6060`

## godoc 不管用

`go get golang.org/x/tools/cmd/godoc`

## 设置Go Proxy


`go env -w  GOPROXY=https://goproxy.cn`

## 热重载

`GO111MODULE=on  go get -u github.com/cosmtrek/air`

> 最前面的 GO111MODULE=on 是只为当前命令启用 Go Module，开启以后我们才能使用 Go Proxy 进行加速。

## vscode 下 golang 语法提示慢

打开 VS Code 的setting, 搜索 go.useLanguageServe, 并勾选上.
默认情况下, 会提示叫你reload，重新打开之后，右下角会自动弹出下载的框框，点击 install即可。
