import { QueryClient } from "@tanstack/solid-query";

export const queryClient = new QueryClient({
	defaultOptions: {
		queries: {
			staleTime: 2,
			retry: false,
		},
	},
});
