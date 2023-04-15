<script setup lang="ts">

</script>

<template>
  <footer>
    <div class="inner">
      <a href="/" class="element">お問い合わせ</a>
      <a href="/" class="element">利用規約</a>
      <a href="/" class="element">個人情報のお取り扱い</a>
    </div>
  </footer>
</template>

<style scoped>
footer {
  position: fixed;
  z-index: 100;
  left:50%;
  bottom:0;
  transform: translate( -50%, 0);
  width:100%;
  background-color: green;
  margin:0 auto;
  text-align: center;
  height: 100px;
}
.inner {
  margin : 40px;
  vertical-align: center;
}
.inner .element {
  color: #fff;
  font-weight: 600;
  padding: 1rem;
}

@media (min-width: 1024px) {
}
</style>