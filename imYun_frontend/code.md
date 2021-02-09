## vue-elemnt-admin开发笔记

### 全局状态码

| 状态           | 状态码 |
| -------------- | ------ |
| 访问正常       | 1000   |
| 表单错误       | 1001   |
| 验证出错       | 1002   |
| 数据库出错     | 1003   |
| 非法访问       | 1004   |
| 内部服务器错误 | 1005   |






#### 请求链接微`localhost:9257`后面拼接的是自己的服务器地址
原因：没有加请求协议，默认为字符串，凭借在地址后面

解决方法：加上请求协议，代码显示如下
```
ENV = 'development'
# base api
# VUE_APP_BASE_API = '/dev-api'
VUE_APP_BASE_API = 'http://192.168.50.232'
```



#### `Value of key 'ID' is not a string!`

> 翻译警告

解决方法

````javascript
// lang/index.js
const i18n = new VueI18n({
  // set locale
  // options: en | zh | es
  locale: getLanguage(),
  silentTranslationWarn: true,
  // set locale messages
  messages
})


````





```
fileId   文件id
fileTitle	文件标题
fileColor	打印颜色
fileDirection	打印方向
fileNum			打印数量
fileIsSingle	单双面
Serial			编号
fileNote		打印备注
userNames		用户名
userPhone		用户手机
fileUrl			下载链接
```


```
文件地址
https://al-yun-oss.oss-cn-hangzhou.aliyuncs.com/PrintYun/UploadFile/ + fileName
```
### Element UI表格批量删除

> 问题产生原因
  数据表格没有下标index,无法获取选中的数据表格
> 解决方法
  外层循环原数组，内层批量选中的数组

> 代码演示
```
deleteAll(row) {
      if (this.listChecked.length === 0) {
        this.$notify({
          title: '失败',
          message: 'Sorry,你还没有选择要删除的文件哦',
          type: 'error',
          duration: 3000
        })
        return
      }
      for (let i = this.list.length; i > 0; i--) {
        for (let j = 0; j < this.listChecked.length; j++) {
          if (
            this.list[i - 1] === this.listChecked[j]
          ) {
            this.list.splice(i - 1, 1)
          }
        }
      }
  }
```