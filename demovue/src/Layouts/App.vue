
<template>
  <div class="container">
    <ul>
      <li>
        <v-link href="/count">登录</v-link>
        <br>
        <v-link href="/signup">无账号请注册</v-link>
      </li>
    </ul>

  <router-view></router-view>
  </div>
</template>

<script>
  import VLink from '../components/VLink.vue'

  export default {
    components: {
      VLink
    }
  }
</script>

<style scoped>
  .container {
    max-width: 600px;
    margin: 0 auto;
    padding: 15px 30px;
    background: #f9f7f5;
  }
</style>