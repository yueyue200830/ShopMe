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

export function getCategory() {
  return request({
    url: '/allCategories',
    method: 'get',
  })
}

export function createProduct(data) {
  return request({
    url: '/product',
    method: 'post',
    data
  })
}

export function updateProduct(data) {
  return request({
    url: '/product',
    method: 'put',
    data
  })
}

export function deleteProduct(query) {
  return request({
    url: '/product',
    method: 'delete',
    params: query
  })
}

export function getProductDetails(query) {
  return request({
    url: '/productDetails',
    method: 'get',
    params: query
  })
}

export function updateProductDetails(data) {
  return request({
    url: '/productDetail',
    method: 'post',
    data
  })
}
