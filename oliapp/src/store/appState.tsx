import type { AppMessage, MessageType } from "@/domain";
import { createContext, createUniqueId, useContext, type ParentProps } from "solid-js";
import { createStore, produce } from "solid-js/store";

type AppStoreActions = {
	addMessage: (message: string, type: MessageType) => void;
	deleteMessage: (id: string) => void;
};

type AppStore = {
	messages: AppMessage[];
} & AppStoreActions;

export const AppStateContext = createContext<AppStore>({
	messages: [],
	addMessage: () => { },
	deleteMessage: () => { },
});

export function AppStateProvider(props: ParentProps) {
	const [messages, setMessages] = createStore<AppMessage[]>([]);

	const addMessage = (message: string, type: MessageType) => {
		setMessages(produce((old) => {
			old.push({ id: createUniqueId(), message, type });
		}));
	};

	const deleteMessage = (id: string) => {
		setMessages(produce((old) => {
			const index = old.findIndex((message) => message.id === id);
			if (index >= 0) {
				old.splice(index, 1);
			}
		}));
	};

	const store = {
		messages,
		addMessage,
		deleteMessage,
	};

	return (
		<AppStateContext.Provider value={store}>
			{props.children}
		</AppStateContext.Provider>
	);
}

export function useAppState() {
	return useContext(AppStateContext);
}
