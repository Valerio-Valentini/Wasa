import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import GroupView from '../views/GroupView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/home', redirect: "/"},
		{path: '/', component: HomeView},
		{path: '/login', component: LoginView},
		{path: '/createGroup', component: GroupView},
		{ path: "/:catchAll(.*)", component: HomeView},
	]
})

export default router
