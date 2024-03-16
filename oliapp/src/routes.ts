import type { RouteDefinition } from "@solidjs/router";
import { lazy } from "solid-js";

export const routes: RouteDefinition[] = [
	{
		path: "/",
		component: lazy(() => import("./pages/User/User")),
		children: [
			{
				path: "/",
				component: lazy(() => import("./pages/User/Index")),
			},
			{
				path: "/home",
				component: lazy(() => import("./pages/User/Home/Home")),
			},
			{
				path: "/companies",
				component: lazy(() => import("./pages/User/Companies/Companies")),
			},
			{
				path: "/companies/add",
				component: lazy(() => import("./pages/User/AddCompany/AddCompany")),
			},
		]
	},
	{
		path: "/auth",
		component: lazy(() => import("./pages/Auth/Auth")),
		children: [
			{
				path: "signin",
				component: lazy(() => import("./pages/Auth/SignIn/SignIn")),
			},
			{
				path: "signup",
				component: lazy(() => import("./pages/Auth/SignUp/SignUp")),
			},
		],
	},
	{
		path: "/*404",
		component: lazy(() => import("./pages/NotFound/NotFound")),
	},
];
