import Router from 'vue-router'
//注册：
Vue.use(Router)
//实例化
export default new Router({
  routes: [
    {
      path: '/count',
      component:Count
    } 
     ]
})