import { createRouter, createWebHistory } from "vue-router";

export const router = createRouter({
	history: createWebHistory(),
	routes: [
		{
			path: "/",
			component: () => import("./home/Home.vue"),
			name: "home",
		},
		{
			path: "/signup",
			component: () => import("./signup/SignUp.vue"),
			name: "signup",
		},
		{
			path: "/signin",
			component: () => import("./signin/SignIn.vue"),
			name: "signin",
		},
	],
});
