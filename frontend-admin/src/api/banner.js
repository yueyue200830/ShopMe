import request from '@/utils/request'

export function getBannerList(query) {
  return request({
    url: '/banners',
    method: 'get',
    params: query
  })
}
