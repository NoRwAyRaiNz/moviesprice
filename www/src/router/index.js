import Vue from 'vue'
import Router from 'vue-router'
import HomePage from '@/views/homePage/index'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      component: HomePage
    }
  ]
})
