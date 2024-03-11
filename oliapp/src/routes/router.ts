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
			path: "/auth",
			component: () => import("./auth/Auth.vue"),
			children: [
				{
					path: "signin",
					component: () => import("./auth/signin/SignIn.vue"),
					name: "signin",
				},
				{
					path: "signup",
					component: () => import("./auth/signup/SignUp.vue"),
					name: "signup",
				},
			]
		},
	],
});
