import Vue from 'vue';
import Router from 'vue-router';

import Home from './components/Home.vue';
import Result from './components/Result.vue';


Vue.use(Router);

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  linkActiveClass: 'active',
  linkExactActiveClass: 'exact-active',
  scrollBehavior() {
    return { x: 0, y: 0 };
  },
  routes: [
    {
      path: '/',
      redirect: '/home',
    },
    {
      path: '/home',
      name: 'home',
      component: Home,
    },
    {
      path: '/result/::Pid',
      name: 'result',
      component: Result,
    }      
  ],
});
