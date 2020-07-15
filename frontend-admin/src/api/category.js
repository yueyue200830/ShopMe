import request from '@/utils/request'

export function getCategories(query) {
  return request({
    url: '/categories',
    method: 'get',
    params: query
  })
}

export function createCategory(data) {
  return request({
    url: '/category',
    method: 'post',
    data
  })
}

export function updateCategory(data) {
  return request({
    url: '/category',
    method: 'put',
    data
  })
}

export function deleteCategory(query) {
  return request({
    url: '/category',
    method: 'delete',
    params: query
  })
}
