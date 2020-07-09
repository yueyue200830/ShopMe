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

export function getProductList() {
  return request({
    url: '/allProducts',
    method: 'get'
  })
}

export function updateBanner(data) {
  return request({
    url: '/banner',
    method: 'post',
    data
  })
}
