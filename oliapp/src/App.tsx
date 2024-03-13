import type { ParentProps } from "solid-js";
import { AppMessages } from "@/components";
import { AppStateProvider } from "@/store";

function App(props: ParentProps) {
	return (
		<AppStateProvider>
			{props.children}
			<AppMessages />
		</AppStateProvider>
	)
}

export default App;
