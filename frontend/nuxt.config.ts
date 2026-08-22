// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2025-07-15',
  devtools: { enabled: true },
  modules: [
    '@unocss/nuxt',
    '@nuxt/fonts'
  ],
  components: [
    {
      path: '~/components',
      pathPrefix: false,
    },
  ],
  fonts: {
    provider: 'google',
  },
  app: {
    head: {
      link: [
        { rel: 'icon', type: 'image/webp', href: '/avatar.webp'}
      ]
    }
  },
  runtimeConfig: {
    apiBase: process.env.NUXT_API_BASE || 'http://backend:8080',
    public: {
      apiBase: process.env.NUXT_PUBLIC_API_BASE || ''
    }
  },
  nitro: {
    preset: 'node-server'
  }
})