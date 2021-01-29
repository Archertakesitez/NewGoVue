<template>
<div id="app">
    <p>您的浏览次数是：{{ message }}</p>
    
</div>
</template>

<script>
  export default {
    components: {
      
    },
  name: 'Count',
  mounted() {
      this.getCount()
  },
  //el: '#hello',
  data() {
    return {
        message:"n"
    }
  },methods: {
    getCount(){
        this.message=this.$route.query.count
    }
  },
  props: {
    msg: String
  }
}
</script>