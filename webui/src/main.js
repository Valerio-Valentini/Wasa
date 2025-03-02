import {createApp, reactive} from 'vue'
import App from './App.vue'
import router from './router'
import axios from './services/axios.js';
import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'
import User from './components/User.vue'
import ChatBox from './components/ChatBox.vue';
import Message from './components/Message.vue';

import './assets/dashboard.css'
import './assets/main.css'

const app = createApp(App)
app.config.globalProperties.$axios = axios;
app.component("ErrorMsg", ErrorMsg);
app.component("LoadingSpinner", LoadingSpinner);
app.component("UserItem", User);
app.component("ChatBox", ChatBox);
app.component("Message", Message)
app.use(router)
app.mount('#app')
