import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/token',
    method: 'post',
    data
  })
}

export function getInfo(token) {
  return request({
    url: '/mine',
    method: 'get',
    params: { token }
  })
}

export function logout() {
  return request({
    url: '/logout',
    method: 'post'
  })
}


export function getList(data) {
  return request({
    url: '/user/list',
    method: 'get',
    params: data
  })
}


export function addUser(data) {
  return request({
    url: '/user/add',
    method: 'post',
    data
  })
}

export function editUser(id, data) {
  return request({
    url: '/user/edit/'+id,
    method: 'put',
    data
  })
}

export function deleteUser(id) {
  return request({
    url: '/user/delete/'+id,
    method: 'delete'
  })
}
