export const state = () => ({
  locales: ['es', 'en'],
  locale: 'es',
  farmingData: {}
})

export const mutations = {
  SET_LANG(state, locale) {
    if (state.locales.includes(locale)) {
      state.locale = locale
    }
  },

  setFarmingData(state, fmData){
    state.farmingData = fmData
  }
}
