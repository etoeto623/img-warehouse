> 本项目目的是开发一个图床工具，通过alist，使用各种网盘做为图片的存储库
# 开发计划
开发计划如下：
- [x] 实现图片的显示
- [x] 读取剪切板的图片 [clipboard](https://pkg.go.dev/golang.design/x/clipboard)
- [x] server可指定启动端口和接口前缀
- [ ] 实现Alist服务器等信息可配置
- [ ] client mode实现服务端信息配置
- [ ] server mode添加日志

# 接口列表
- 上传文件
``` bash
/api/fs/put

Content-Length: 70524
Content-Type: image/png
File-Path: %2xxx.png
Password: xxxx
```

# 环境依赖
## ubuntu
ubuntu上需要安装x11的开发依赖
``` bash
sudo apt install libx11-dev
```

# 使用方法
## 做为服务端
``` bash
imghouse -s
```
## 做为客户端
``` bash
imghouse [-fn fileName] [-f local image file path]
```
## 配置示例
``` json
{
    "port": 8888,
    "image_view_api": "view",
    "alist_url": "http://xxx.xxx:5244",
    "alist_password": "xxxx",
    "server_url": "http://xxx.xxx",
    "aes_key": "xvvsdsdf"
}
```