import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import GroupView from '../views/GroupView.vue'
import NickView from '../views/NickView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/home', redirect: "/"},
		{path: '/', component: HomeView},
		{path: '/login', component: LoginView},
		{path: '/createGroup', component: GroupView},
		{ path: "/:catchAll(.*)", component: HomeView},
		{path: '/changeNick', component: NickView},
	]
})

export default router
