import request from '@/utils/request'

export function searchUser(name) {
  return request({
    url: '/v1/webs/search/user',
    method: 'get',
    params: { name }
  })
}

export function transactionList(query) {
  return request({
    url: '/v1/webs/transaction/list',
    method: 'get',
    params: query
  })
}
