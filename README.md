# imYun 一个简化打印店工作流程的系统
[![](https://img.shields.io/badge/license-MIT-green)](https://github.com/run-nerver/imYun_wx/blob/main/LICENSE)
[![](https://img.shields.io/badge/Go-1.15-brightgreen)](https://golang.org/)
[![](https://img.shields.io/badge/vue-2.6.10-brightgreen.svg?style=flat-square)](https://github.com/vuejs/vue)
[![](https://img.shields.io/badge/vue--element--admin-4.3.1-brightgreen)](https://panjiachen.github.io/vue-element-admin-site/zh/)
## 传统打印店流程:
用户到店U盘（微信、QQ）发送文件&rarr;店主接收&rarr;打印  
## imYun打印流程:
用户通过小程序上传文件&rarr;店主在线预览(下载)&打印  
## 系统截图:
![后台](https://github.com/run-nerver/imYun/blob/main/images/%E5%89%8D%E7%AB%AF%E9%A6%96%E9%A1%B5.png)  
![小程序](https://github.com/run-nerver/imYun/blob/main/images/%E5%B0%8F%E7%A8%8B%E5%BA%8F.jpg)
## 使用说明:
### - 本机测试(此种方式仅供测试使用，如果开店用请移步至下面公网部署)
1、确保机器安装docker及docker-compose  
2、下载源码，修改imYun_backend/config.repo.yaml下的wechat的appid和secret为自己小程序的。注:如果仅测试后台，不需要小程序上传，此步可省略。  
3、
```
cd docker
docker-compose up
```
访问http://127.0.0.1:9527即可访问后台，账号:admin，密码:123456  

### - 公网部署(此方式开店可用，需要公网服务器)
1、确保机器安装docker及docker-compose  
2、下载源码，修改imYun_backend/config.repo.yaml下的wechat的appid和secret为自己小程序的。  
3、将imYun_frontend/src/api/global_variable.js中的127.0.0.1修改为服务器IP或域名(后面的:5000不要修改)    
4、修改[小程序端](https://github.com/run-nerver/imYun_wx)api/url.js下的URL_SERVER的127.0.0.1为服务器IP，按照微信官方步骤将小程序发布(如果仅测试，使用小程序开发工具打开)  
5、
```
cd docker
docker-compose up
```
访问http://公网IP:9527即可访问后台，账号:admin，密码:123456  
## 其他说明:
1、有任何问题或者建议欢迎提issue或者联系sailor0913@qq.com
