/** When your routing table is too long, you can split it into small modules**/

import Layout from '@/layout'

const chartsRouter = {
  path: '/user',
  component: Layout,
  redirect: 'noRedirect',
  children: [
    {
      path: 'index',
      component: () => import('@/views/user/index'),
      name: 'user',
      meta: { title: '用户管理', icon: 'el-icon-user-solid', affix: true }
    }
  ]
}

export default chartsRouter
