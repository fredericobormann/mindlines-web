// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-11-01',
  devtools: { enabled: true },

  runtimeConfig: {
    public: {
      backendUrl: '/api',
    },
  },

  colorMode: {
    preference: 'light',
  },

  vue: {
    propsDestructure: true,
  },

  nitro: {
    routeRules: {
      '/api/**': {
        headers: {
          'Content-Type': 'application/json; charset=utf-8',
        },
      }
    }
  },

  modules: ['@nuxt/ui', '@hebilicious/vue-query-nuxt']
})
