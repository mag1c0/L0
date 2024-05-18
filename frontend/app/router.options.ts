import type {RouterConfig} from '@nuxt/schema'

export default <RouterConfig>{
    mode: "history",
    base:
        "/",
    linkActiveClass:
        "nuxt-link-active",
    linkExactActiveClass:
        "nuxt-link-exact-active",
    strict: true,
    routes: (_routes) => [
        {
            path: "/",
            component: () => import('~/components/vPages/vHome.vue').then(r => r.default || r),
            name: "index"
        }
    ]
}