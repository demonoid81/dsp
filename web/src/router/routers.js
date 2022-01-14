

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
]