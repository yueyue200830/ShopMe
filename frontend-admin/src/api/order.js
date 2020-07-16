import request from '@/utils/request'

export function getOrders(query) {
  return request({
    url: '/orders/all',
    method: 'get',
    params: query
  })
}

export function cancelOrder(data) {
  return request({
    url: '/cancelOrder',
    method: 'post',
    data
  })
}

export function finishOrder(data) {
  return request({
    url: '/order/finish',
    method: 'post',
    data
  })
}

export function getOrderDetail(query) {
  return request({
    url: '/order',
    method: 'get',
    params: query
  })
}
