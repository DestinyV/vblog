import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/vLog/user/login',
    method: 'post',
    data
  })
}

export function getInfo(token) {
  return request({
    url: '/vLog/user/getInfo',
    method: 'get',
    params: { token }
  })
}

export function logout() {
  return request({
    url: '/vLog/user/logout',
    method: 'post'
  })
}
