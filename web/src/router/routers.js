

export default [
    {
        path: '/login',
        name: 'login',
        meta: {
            title: 'Login',
            hideInMenu: true
        },
        component: () => import('view/login/login.vue')
    },
    {
        path: '/',
        name: '_home',
        redirect: '/home',
        component: () => import('components/main'),
        meta: {
            hideInMenu: true,
            notCache: true
        },
        children: [
            {
                path: '/home',
                name: 'home',
                meta: {
                    hideInMenu: true,
                    title: 'Главная страница',
                    notCache: true,
                    icon: 'md-home'
                },
                component: () => import('view/single-page/home')
            }
        ]
    },
    {
        path: '/statistics',
        name: 'statistics',
        component: () => import('components/main'),
        meta: {
            hideInBread: true
        },
        children: [
            {
                path: 'statistics',
                name: 'statistics',
                meta: {
                    icon: '_qq',
                    title: 'Statistics'
                },
                component: () => import('view/statistics/statistics.vue')
            }
        ]
    },
    {
        path: '/dsp',
        name: 'dsp',
        component: () => import('components/main'),
        meta: {
            hideInBread: true
        },
        children: [
            {
                path: 'dsp',
                name: 'dsp',
                meta: {
                    icon: '_qq',
                    title: 'DSP'
                },
                component: () => import('view/dsp/sspTable.vue')
            }
        ]
    },
    {
        path: '/companies',
        name: 'companies',
        component: () => import('components/main'),
        meta: {
            hideInBread: true
        },
        children: [
            {
                path: '',
                name: 'companies',
                meta: {
                    icon: '_qq',
                    title: 'Companies'
                },
                component: () => import('view/statistics/statistics.vue')
            }
        ]
    },
]