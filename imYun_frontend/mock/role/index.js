const Mock = require('mockjs')
const { deepClone } = require('../utils')
const { asyncRoutes, constantRoutes } = require('./routes.js')

const routes = deepClone([...constantRoutes, ...asyncRoutes])

const roles = [
  {
    key: 'admin',
    name: 'admin',
    description: 'Super Administrator. Have access to view all pages.',
    routes: routes
  },
  {
    key: 'editor',
    name: 'editor',
    description: 'Normal Editor. Can see all pages except permission page',
    routes: routes.filter(i => i.path !== '/permission')// just a mock
  },
  {
    key: 'visitor',
    name: 'visitor',
    description: 'Just a visitor. Can only see the home page and the document page',
    routes: [{
      path: '',
      redirect: 'dashboard',
      children: [
        {
          path: 'dashboard',
          name: 'Dashboard',
          meta: { title: 'dashboard', icon: 'dashboard' }
        }
      ]
    }]
  }
]

module.exports = [
  // mock get all routes form server
  {
    url: '/v1/webs/routes',
    type: 'get',
    response: _ => {
      return {
        code: 1000,
        data: routes
      }
    }
  },

  // mock get all roles form server
  {
    url: '/v1/webs/roles',
    type: 'get',
    response: _ => {
      return {
        code: 1000,
        data: roles
      }
    }
  },

  // add role
  {
    url: '/v1/webs/role',
    type: 'post',
    response: {
      code: 1000,
      data: {
        key: Mock.mock('@integer(300, 5000)')
      }
    }
  },

  // update role
  {
    url: '/v1/webs/role/[A-Za-z0-9]',
    type: 'put',
    response: {
      code: 1000,
      data: {
        status: 'success'
      }
    }
  },

  // delete role
  {
    url: '/v1/webs/role/[A-Za-z0-9]',
    type: 'delete',
    response: {
      code: 1000,
      data: {
        status: 'success'
      }
    }
  }
]
