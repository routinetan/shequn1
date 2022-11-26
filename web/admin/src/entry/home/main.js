import Vue from 'vue';
import iView from 'iview';
import VueRouter from 'vue-router';
import Request from '../../api/request.js'
import app from './App';
import store from '../../store';
import router from '../../routes';
import './index.less';

Vue.prototype.Request = Request
Vue.use(VueRouter);
Vue.use(iView);

window._vueInstance = new Vue({
  el: '#app',
  components: { app },
  template: '<app/>',
  store,
  router
});
