import request from '@/utils/request'

export function deleteOrder(orderId) {
  return request({
    url: '/v1/webs/order/odelete',
    method: 'delete',
    params: { orderId }
  })
}

export function deleteOrders(orderIdList) {
  return request({
    url: '/v1/webs/order/odelete',
    method: 'delete',
    params: orderIdList
  })
}

export function downloadFile(filename) {
  return request({
    url: '/v1/webs/order/getLocalFile',
    method: 'get',
    params: { ReFileName: filename },
    responseType: 'blob'
  })
}
