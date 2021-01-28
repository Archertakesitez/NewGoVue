import Vue from 'vue'
//import App from './Layouts/App.vue'
///import routes from './routes'
import VueRouter from 'vue-router'
import Login from './pages/Login.vue'
import Count from './pages/Count.vue'
import Signup from './pages/Signup.vue'
Vue.use(VueRouter)
/** const app = new Vue({
  el: '#app',
  data: {
    currentRoute: window.location.pathname
  },
  computed: {
    ViewComponent () {
      const matchingView = routes[this.currentRoute]
      return matchingView
        ? require('./pages/' + matchingView + '.vue').default
        : require('./pages/404.vue').default
    }
  },
  render (h) {
    console.log("view",this.ViewComponent)
    return h(this.ViewComponent)
  }
})

window.addEventListener('popstate', () => {
  app.currentRoute = window.location.pathname
})
**/
const routes = [
  { path: '/', component: Login },
  { path: '/signup', component: Signup },
  { path: '/count', component: Count }
]
const router = new VueRouter({
  //mode: 'history',
  routes // short for `routes: routes`
})
new Vue({
  el: '#app',
  router,
  render: h => h(Login)
}).$mount('#app')
/**window.addEventListener('popstate', () => {
  app.currentRoute = window.location.pathname
})
**/