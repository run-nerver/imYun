import request from '@/utils/request'

export function fetchList(query) {
  return request({
    url: '/v1/webs/printfile/list',
    method: 'get',
    params: query
  })
}

export function fetchArticle(id) {
  return request({
    url: '/v1/webs/printfile/detail',
    method: 'get',
    params: { id }
  })
}

export function fetchPv(pv) {
  return request({
    url: '/v1/webs/printfile/pv',
    method: 'get',
    params: { pv }
  })
}

export function createArticle(data) {
  return request({
    url: '/v1/webs/printfile/create',
    method: 'post',
    data
  })
}

export function updateArticle(data) {
  return request({
    url: '/v1/webs/printfile/update',
    method: 'post',
    data
  })
}
