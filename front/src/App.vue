<script setup>
import HeaderComp from "./components/header/HeaderComp.vue";
</script>

<template>
  <div class="contents">
    <HeaderComp />
    <router-view />
  </div>
</template>

<style lang="scss">
  @import "@/assets/styles/style.scss";
  .contents {
    height: 100vh;
    background-color: #f5f5f5;
  }
</style>