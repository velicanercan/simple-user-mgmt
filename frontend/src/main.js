import { createApp } from 'vue';
import App from './App.vue';
import axios from 'axios';
import 'bootstrap/dist/css/bootstrap.css';
import { createRouter, createWebHistory } from 'vue-router';
import UserList from './components/pages/UserList';
import UserCreate from './components/pages/UserCreate';
import UserEdit from './components/pages/UserEdit';
import UserShow from './components/pages/UserShow';
import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css';
import { useVuelidate } from '@vuelidate/core';
axios.defaults.baseURL = process.env.VUE_APP_API_URL

const router = createRouter({
  history: createWebHistory(),
  routes: [ // define routes and title for each route
    { path: '/', component: UserList, meta: { title: 'User Management' } },
    { path: '/create', component: UserCreate, meta: { title: 'New User' } },
    { path: '/edit/:id', component: UserEdit, meta: { title: 'Edit User' } },
    { path: '/show/:id', component: UserShow, meta: { title: 'User Details' } },
  ],
});
router.beforeEach((to, from, next) => {
    document.title = to.meta.title || 'User Management';
    next();
});
const app = createApp(App)
app.use(router)
app.use(ElementPlus)
app.use(useVuelidate)
app.mount('#app')