import type { RouteDefinition } from "@solidjs/router";
import { lazy } from "solid-js";

export const routes: RouteDefinition[] = [
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
