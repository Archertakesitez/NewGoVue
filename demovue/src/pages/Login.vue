
<template>

  <div id="app">
  <!-- <main-layout> -->
    <!-- <h1>{{ username }}</h1> -->
    <h1>{{ message }}</h1>
    <router-link to="/count">Go to Count</router-link>
    <!-- <a @click="getclick">click me</a> -->
    <!-- <p>{{username}}</p> -->
    <input name="username" type="text"  placeholder="请输入用户名" v-model="username">
    <br>
    <input name="pwd" type="text" placeholder="请输入密码" v-model="pwd">
    <br>
    
    <a href="/count?id={{">登录</a>
    <br>

    <a href="/signup">无账号请注册</a>
    <br>
    <a @click="postReq">跳转 </a> 
     
    


    
 
  <!-- </main-layout> -->
  </div>
</template>

<script>
//import MainLayout from '../Layouts/App.vue'
const axios = require('axios');

export default {
  components:{
    //MainLayout
    
  },
  name: 'Login',
  mounted() {
  },
  //el: '#hello',
  data() {
    return {
      username:"",
      pwd: "",
      message:"请登录"
    }
  },
  methods: {
   // getclick(){
      //axios
       // .get('http://localhost:8080/')
       // .then(response =>{
       //   this.username = response.data.Username
        //  this.count = response.data.Count
       // } )
   // },
   requestSent:function(data){
     axios({
       method:"POST",
       url:"http://localhost:8080/",
       data: data,
        headers: {
                        "content-type": "text/plain"
                    }
     }).then(result => {
        //this.count=result.data
        this.message=result.data
     }
     )
     console.log(data)
   },
    postReq: function(){
      var data={
        "username":this.username,
        "pwd":this.pwd
      }
      
      this.requestSent(data);
      
    }
  },
  props: {
    msg: String
  }
}


</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
