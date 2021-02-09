import request from '@/utils/request'

export function getRoutes() {
  return request({
    url: '/v1/webs/routes',
    method: 'get'
  })
}

export function getRoles() {
  return request({
    url: '/v1/webs/roles',
    method: 'get'
  })
}

export function addRole(data) {
  return request({
    url: '/v1/webs/role',
    method: 'post',
    data
  })
}

export function updateRole(id, data) {
  return request({
    url: `/vue-element-admin/role/${id}`,
    method: 'put',
    data
  })
}

export function deleteRole(id) {
  return request({
    url: `/vue-element-admin/role/${id}`,
    method: 'delete'
  })
}
