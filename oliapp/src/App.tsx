import { queryClient } from "@/api/queryClient";
import { AppMessages } from "@/components";
import { AppStateProvider } from "@/store";
import { QueryClientProvider } from "@tanstack/solid-query";
import type { ParentProps } from "solid-js";

function App(props: ParentProps) {
	return (
		<AppStateProvider>
			<QueryClientProvider client={queryClient}>
				{props.children}
			</QueryClientProvider>
			<AppMessages />
		</AppStateProvider>
	)
}

export default App;
