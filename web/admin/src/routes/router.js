import Main from '@/components/main/main';
export default [{
  path: '/',
  redirect: '/dashboard',
  component: Main,
  children: [
    {
      path: '/home',
      name: 'home',
      meta: {
        hideInMenu: true,
        title: '看板',
        notCache: true,
        icon: 'md-home'
      },
      component: () => import('@/view/home/group')
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      meta: {
        hideInMenu: true,
        title: '看板',
        notCache: true,
        icon: 'md-home'
      },
      component: () => import('@/view/dashboard/index')
    },
    {
      path: '/group',
      name: 'group',
      meta: {
        hideInMenu: true,
        title: '群管理',
        notCache: true,
        icon: 'md-home'
      },
      component: () => import('@/view/home/group')
    },
    {
      path: '/group/create',
      name: 'group_create',
      redirect: '/group/create_nav1',
      meta: {
        hideInMenu: true,
        title: '群管理-详细',
        notCache: true,
        isEdit:true,
        icon: 'md-home'
      },
      component: () => import('@/view/home/create'),
      children: [
        {
          path: '/group/create_nav1',
          name: 'group_create_nav1',
          meta: {
            hideInMenu: true,
            title: '群管理',
            notCache: true,
            icon: 'md-home'
          },
          component: () => import('@/view/home/create_nav1')
        },
      ]
    },
    {
      path: '/group/:id',
      name: 'group_edit',
      redirect: '/group/:id/edit_nav1',
      meta: {
        hideInMenu: true,
        title: '群管理-详细',
        notCache: true,
        isEdit:true,
        icon: 'md-home'
      },
      component: () => import('@/view/home/edit'),
      children: [
        {
          path: '/group/:id/edit_nav1',
          name: 'group_edit_nav1',
          meta: {
            hideInMenu: true,
            title: '群管理',
            notCache: true,
            icon: 'md-home'
          },
          component: () => import('@/view/home/edit_nav1')
        },
      ]
    },
    {
      path: '/group/:id/qrcode',
      name: 'qrcode',
      meta: {
        hideInMenu: true,
        title: '二维码管理',
        notCache: true,
        parent:"group",
        icon: 'md-home'
      },
      component: () => import('@/view/qrcode/index')
    },
    {
      path: '/group/:id/qrcode/add',
      name: 'qrcode_add',
      meta: {
        hideInMenu: true,
        title: '二维码添加',
        notCache: true,
        icon: 'md-home'
      },
      component: () => import('@/view/qrcode/create')
    },
    {
      path: '/setting',
      name: 'setting',
      meta: {
        hideInMenu: true,
        title: '设置',
        notCache: true,
        icon: 'md-home'
      },
      component: () => import('@/view/setting/setting')
    },
  ]
}];
