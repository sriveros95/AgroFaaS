<template>
  <div class="bg">
    <section class="float-bottom">
      <v-layout column wrap class="my-12" align-center>
        <v-flex xs12 sm4 class="my-4">
          <div class="text-center">
            <h1 class="mb-2 display-1 text-center">{{ $t('home.title') }}</h1>
            <h2 class="headline">{{ $t('home.sub-title') }}</h2>
          </div>
        </v-flex>
      </v-layout>
    </section>
    <section class="overHead">
      <v-layout column wrap align-center>
        <v-tooltip bottom>
          <template v-slot:activator="{ on }">
            <span class="theMoon" v-on="on">{{ $store.state.farmingData.moonCode | moonEmojize }}</span>
          </template>
          <span>{{ $store.state.farmingData.moonString }}</span>
        </v-tooltip>
        <h3 v-if="$store.state.farmingData.plantingTime">It's Planting Time!</h3>
        <v-tooltip top>
          <template v-slot:activator="{ on }">
            <span class="theEarth" v-on="on">{{ $store.state.farmingData.what2Plant | earthEmojize }}</span>
          </template>
          <span>{{ $store.state.farmingData.farmingString }}</span>
        </v-tooltip>
        <p class="text-center"  v-if="$store.state.farmingData.plantingTime">{{ $store.state.farmingData.what2Plant }}</p>
      </v-layout>
    </section>
  </div>
</template>
<style scoped>
* {
  color: white;
}
.float-bottom {
  position: relative;
  bottom: -49vh;
}
.overHead{
  position: absolute;
  top: 5%;
  left: 0;
  right: 0;
}
.theMoon{
  font-size: 12vh
}
.theEarth{
  font-size: 6vh
}
.bg {
  position: absolute;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  background-color: black; /* For browsers that do not support gradients */
  background-image: linear-gradient(
    #4444ff,
    #222299,
    black
  ); /* (Top -> Bottom) */
}
</style>
<script>
const axios = require("axios");
var emoji = require("node-emoji");
export default {
  async fetch({ store, params }) {
    let { data } = await axios.get("https://sriveros95.o6s.io/moon-farmer");
    store.commit("setFarmingData", data);
    document.body.style.backgroundAttachment = "fixed"; 
  },
  filters: {
    moonEmojize(moonCode) {
      switch (moonCode) {
        case "NW":
          return emoji.get('new_moon')
        case "XG":
          return emoji.get('waxing_gibbous_moon')
          break;
        case "FQ":
          return emoji.get('first_quarter_moon')
        case "XC":
          return emoji.get('waxing_crescent_moon')
        case "FL":
          return emoji.get('full_moon')
        case "NG":
          return emoji.get('waning_gibbous_moon')
        case "LQ":
          return emoji.get('last_quarter_moon')
        case "NC":
          return emoji.get('waning_crescent_moon')
        default:
          return emoji.get('coffee')
      }
    },
    earthEmojize(what2Plant) {
      switch (what2Plant) {
        case "Above ground annuals, especially Leaf plants also Cereals, Herbs, Cucumbers":
          return emoji.emojify(':seedling: :ear_of_rice: :herb: :cucumber:')
        case "Below ground plants, especially Root plants, Plant trees, shrubs and perennials.":
          return emoji.emojify(':evergreen_tree: :potato: :deciduous_tree: :carrot:"')
        default:
          return ""
      }
    }
  }
};
</script>