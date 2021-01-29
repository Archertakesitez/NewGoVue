
<template>

  <div id="app">
  <!-- <main-layout> -->
    <!-- <h1>{{ username }}</h1> -->
    <h1>{{ message }}</h1>
  
    <!-- <a @click="getclick">click me</a> -->
    <!-- <p>{{username}}</p> -->
    <input name="username" type="text"  placeholder="请输入用户名" v-model="username">
    <br>
    <input name="pwd" type="text" placeholder="请输入密码" v-model="pwd">
    <br>
    
    <a @click="postReq">登录</a>
    <br>

    <router-link to="/signup">无账号请注册</router-link>
    <br>
    
    <router-view></router-view>
    


    
 
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
    changeOne(data){
        //const vs = this;
        if(this.message=="登录成功！跳转中..."){
            //this.message="nijuedene"
            this.$router.push({path: '/count', query: {username: this.username, count:data.count}});
        }
         console.log("hello")
    },
    
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
        this.message=result.data.message
        console.log(result.data.message)
        var data={
        "count":result.data.count
      }
        this.changeOne(data);
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
  text-decoration: none;
  cursor:pointer;


  color: #42b983;
}
.router-link-active {    
  text-decoration: none;
  
}
</style>
