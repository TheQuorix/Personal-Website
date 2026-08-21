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
    public: {
      apiBase: 'http://localhost:8080'
    }
  },
  nitro: {
    preset: 'node-server'
  }
})