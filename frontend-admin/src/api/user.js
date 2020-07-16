import request from '@/utils/request'

export function getUsers(query) {
  return request({
    url: '/users',
    method: 'get',
    params: query
  })
}
