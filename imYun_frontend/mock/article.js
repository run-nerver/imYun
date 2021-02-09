const Mock = require('mockjs')

const List = []
const count = 100

const image_url = 'https://wpimg.wallstcn.com/e4558086-631c-425c-9430-56ffb46e70b3'
// const pdf_url = 'http://weiliangdemo.cn1.utools.club/pdf%E6%B5%8B%E8%AF%95%E6%96%87%E6%A1%A3.pdf'
// const execl_url = 'http://weiliangdemo.cn1.utools.club/execl%E6%B5%8B%E8%AF%95%E6%96%87%E6%A1%A3.xlsx'
// const ppt_url = 'http://weiliangdemo.cn1.utools.club/ppt%E6%B5%8B%E8%AF%95%E6%96%87%E6%A1%A3.pptx'
// const word_url = 'http://weiliangdemo.cn1.utools.club/word%E6%B5%8B%E8%AF%95%E6%96%87%E6%A1%A3.docx'

for (let i = 0; i < count; i++) {
  List.push(Mock.mock({
    id: '@increment',
    nickName: '@cname(2,3)',
    'fileName|1': ['社会主义核心价值观.doc', '工作总结.pdf', '每日数据.xls', '哈士奇.jpg', '未知文件.nk', '北斗七星.psd', 'bmptupian.bmp'],
    'fileColor|1': ['黑白', '彩色'],
    'direction|1': ['横向', '纵向'],
    'type|1': ['DOC', 'PDF', 'EXECL', 'TXT', '其他'],
    'singleSide|1': ['单面', '双面'],
    code: '@integer(1000, 9999)',
    'fileNum|1-10': 10,
    'remarks|1': ['加急', '调整后再打印', '明天来取'],
    'fileStatus|1': ['已打印', '未打印', '已删除', '已下载'],
    'userPhone|1': ['13531544954', '13632250649', '15820292420', '15999905612'],
    // 'fileUrl|1': [pdf_url, image_url, execl_url, ppt_url, word_url]
    'fileUrl|1': [image_url]

  }))
}

module.exports = [
  {
    url: '/v1/webs/article/list',
    type: 'get',
    response: config => {
      const { importance, type, title, page = 1, limit = 20, sort } = config.query

      let mockList = List.filter(item => {
        if (importance && item.importance !== +importance) return false
        if (type && item.type !== type) return false
        if (title && item.title.indexOf(title) < 0) return false
        return true
      })

      if (sort === '-id') {
        mockList = mockList.reverse()
      }

      const pageList = mockList.filter((item, index) => index < limit * page && index >= limit * (page - 1))

      return {
        code: 1000,
        data: {
          total: mockList.length,
          items: pageList
        }
      }
    }
  },

  {
    url: '/v1/webs/article/detail',
    type: 'get',
    response: config => {
      const { id } = config.query
      for (const article of List) {
        if (article.id === +id) {
          return {
            code: 1000,
            data: article
          }
        }
      }
    }
  },

  {
    url: '/v1/webs/article/pv',
    type: 'get',
    response: _ => {
      return {
        code: 1000,
        data: {
          pvData: [
            { key: 'PC', pv: 1024 },
            { key: 'mobile', pv: 1024 },
            { key: 'ios', pv: 1024 },
            { key: 'android', pv: 1024 }
          ]
        }
      }
    }
  },

  {
    url: '/v1/webs/article/create',
    type: 'post',
    response: _ => {
      return {
        code: 1000,
        data: 'success'
      }
    }
  },

  {
    url: '/v1/webs/article/update',
    type: 'post',
    response: _ => {
      return {
        code: 1000,
        data: 'success'
      }
    }
  }
]

