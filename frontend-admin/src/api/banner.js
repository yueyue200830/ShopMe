import request from '@/utils/request'

export function getBannerList(query) {
  return request({
    url: '/banners',
    method: 'get',
    params: query
  })
}

export function deleteBanner(query) {
  return request({
    url: '/banner',
    method: 'delete',
    params: query
  })
}

export function getProductNameList() {
  return request({
    url: '/productNames',
    method: 'get'
  })
}
