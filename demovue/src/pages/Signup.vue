<template>
    <div class="container">
    <p>{{ message }}</p>
    <input name="username" type="text"  placeholder="请输入用户名" v-model="username">
    <br>
    <input name="pwd" type="text" placeholder="请输入密码" v-model="pwd">
    <p @click="saveUser"><button type="button">注册</button></p> 
    </div>

</template>

<script>
  
  const axios = require('axios');
  export default {
    components: {
      
    },
  name: 'Signup',
  mounted() {
  },
  //el: '#hello',
  data() {
    return {
      message:"请注册",  
      username:"",
      pwd: ""
    }
  },methods: {
    requestSent(data){
     axios({
       method:"POST",
       url:"http://localhost:8080/signup",
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
      saveUser(){
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