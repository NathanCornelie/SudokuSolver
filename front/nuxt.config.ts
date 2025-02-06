// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: "2024-04-03",
  app: {
    baseURL: '/solver/'
  },
  devtools: { enabled: true },
  modules: ["vuetify-nuxt-module", "nuxt-mdi"],
  appConfig: {
    //apiURL: "http://141.98.153.77/api/solver",
    apiURL: "http://localhost:8081/api/solver",
  },
});
