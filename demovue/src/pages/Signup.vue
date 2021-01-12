<template>
    <div class="container">
    <input name="username" type="text"  placeholder="请输入用户名" v-model="username">
    <br>
    <input name="pwd" type="text" placeholder="请输入密码" v-model="pwd">
    <v-link href="/signsuc"
    v-on:click="alert">注册</v-link>
    </div>

</template>

<script>
  import VLink from '../components/VLink.vue'
  export default {
    components: {
      VLink
    },methods: {
      alert() {
          alert("hey")
  }}}
</script>