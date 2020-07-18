import request from '@/utils/request'

export function getUsers(query) {
  return request({
    url: '/users',
    method: 'get',
    params: query
  })
}

export function getUserNumber() {
  return request({
    url: '/user/number',
    method: 'get',
  })
}
