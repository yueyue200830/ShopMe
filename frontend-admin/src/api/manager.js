import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/manager/login',
    method: 'post',
    data
  })
}

export function getInfo(token) {
  return request({
    url: '/manager/info',
    method: 'get',
    params: { token }
  })
}

export function logout() {
  return request({
    url: '/manager/logout',
    method: 'post'
  })
}

export function getManagers(query) {
  return request({
    url: '/managers',
    method: 'get',
    params: query
  })
}

export function createManager(data) {
  return request({
    url: '/manager',
    method: 'post',
    data
  })
}

export function updateManager(data) {
  return request({
    url: '/manager',
    method: 'put',
    data
  })
}

export function deleteManager(query) {
  return request({
    url: '/manager',
    method: 'delete',
    params: query
  })
}

export function resetManagerPassword(data) {
  return request({
    url: '/manager/password/reset',
    method: 'put',
    data
  })
}

export function validManagerName(query) {
  return request({
    url: '/manager/name/check',
    method: 'get',
    params: query
  })
}
