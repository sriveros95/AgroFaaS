<template>
  <v-footer
    height="88"
    class="justify-center elevation-4 tFoot"
  >
    <social-media large />
  </v-footer>
</template>

<style scoped>
.tFoot{
  background-color:rgba(0,0,0,0.5)!important;
  opacity: 0.65;
  border-color: transparent!important;
  position: absolute;
  bottom: 0;
  right: 0;
  left: 0;
}
</style>

<script>
export default {
  components: {
    SocialMedia: () => import('@/components/SocialMedia')
  }
}
</script>
