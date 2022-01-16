

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
        name: '_statistics',
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
        name: '_dsp',
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
        path: '/campaigns',
        name: '_campaigns',
        component: () => import('components/main'),
        meta: {
            hideInBread: true
        },
        children: [
            {
                path: '',
                name: 'Сampaigns',
                meta: {
                    icon: '_qq',
                    title: 'Сampaigns'
                },
                component: () => import('view/campaigns/campaigns.vue')
            }
        ]
    },
    {
        path: '/setting',
        name: 'Settings',
        component: () => import('components/main'),
        meta: {
            icon: 'logo-buffer',
            title: 'Settings'
        },
        children: [
            {
                path: 'users',
                name: 'users',
                meta: {
                    icon: '_qq',
                    title: 'Users'
                },
                component: () => import('view/statistics/statistics.vue')
            },
            {
                path: 'groups',
                name: 'groups',
                meta: {
                    icon: '_qq',
                    title: 'Groups'
                },
                component: () => import('view/statistics/statistics.vue')
            }
        ]
    },
]