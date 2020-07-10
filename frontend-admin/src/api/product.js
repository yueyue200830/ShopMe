import request from '@/utils/request'

export function getProducts(query) {
  return request({
    url: '/products',
    method: 'get',
    params: query
  })
}

export function getProductNumber() {
  return request({
    url: '/productNumber',
    method: 'get',
  })
}
