// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    ssr: true,
    app: {
        head: {
            title: "WB Tech | L0",
            meta: [
                {charset: "utf-8"},
                {name: "viewport", content: "width=device-width, initial-scale=1, user-scalable=no"},
                {"http-equiv": "Content-Type", content: "text/html; charset=UTF-8"},
            ],
            noscript: [
                {textContent: 'JavaScript is required'}
            ]
        }
    },
    css: [
        '~/assets/css/normalize.css',
        '~/assets/css/main.css',
    ],
    devtools: {enabled: false},
    telemetry: false,
    modules: [
        '@nuxtjs/google-fonts',
    ],
    googleFonts: {
        families: {
            Montserrat: [400, 500, 600, 700]
        },
        prefetch: false,
        preconnect: true,
        preload: true,
        download: true,
        base64: false,
        overwriting: true // this flag
    }
})