import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/v1/webs/user/login',
    method: 'post',
    data
  })
}

export function getInfo(token) {
  return request({
    url: '/v1/webs/user/info',
    method: 'get',
    params: { token }
  })
}

export function logout() {
  return request({
    url: '/v1/webs/user/logout',
    method: 'post'
  })
}
