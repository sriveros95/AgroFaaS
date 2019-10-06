<template>
  <v-app>
    <nuxt />
    <core-footer />
  </v-app>
</template>

<script>
export default {
  components: {
    CoreFooter: () => import('@/components/core/Footer'),
  }
}
</script>
