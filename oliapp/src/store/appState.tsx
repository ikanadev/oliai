import type { ApiError, AppMessage, User } from "@/domain";
import { MessageType } from "@/domain";
import { API_ERRORS_MAP, isApiError } from "@/utils";
import { type ParentProps, createContext, createUniqueId, useContext } from "solid-js";
import { createStore, produce } from "solid-js/store";

const defaultUser: User = { id: "", firstName: "", lastName: "", email: "" };

type AppStoreActions = {
	addMessage: (message: string, type: MessageType) => void;
	addErrorMessage: (message: string) => void;
	addWarningMessage: (message: string) => void;
	addInfoMessage: (message: string) => void;
	addSuccessMessage: (message: string) => void;
	deleteMessage: (id: string) => void;
	// biome-ignore lint/suspicious/noExplicitAny: is checking a type
	handleApiError: (err: any, cb?: (message: string) => void) => void;
	setUser: (user: User) => void;
	clearAppState: () => void;
};

type AppData = { user: User };

type AppStore = {
	messages: AppMessage[];
	appData: AppData;
} & AppStoreActions;

export const AppStateContext = createContext<AppStore>({
	messages: [],
	appData: { user: defaultUser },
	addMessage: () => { },
	addErrorMessage: () => { },
	addWarningMessage: () => { },
	addInfoMessage: () => { },
	addSuccessMessage: () => { },
	deleteMessage: () => { },
	handleApiError: () => { },
	setUser: () => { },
	clearAppState: () => { },
});

export function AppStateProvider(props: ParentProps) {
	const [messages, setMessages] = createStore<AppMessage[]>([]);
	const [appData, setAppData] = createStore<AppData>({ user: defaultUser });

	const addMessage = (message: string, type: MessageType) => {
		setMessages(produce((old) => {
			old.push({ id: createUniqueId(), message, type });
		}));
	};

	const addErrorMessage = (message: string) => addMessage(message, MessageType.ERROR);
	const addWarningMessage = (message: string) => addMessage(message, MessageType.WARNING);
	const addInfoMessage = (message: string) => addMessage(message, MessageType.INFO);
	const addSuccessMessage = (message: string) => addMessage(message, MessageType.SUCCESS);

	const deleteMessage = (id: string) => {
		setMessages(produce((old) => {
			const index = old.findIndex((message) => message.id === id);
			if (index >= 0) {
				old.splice(index, 1);
			}
		}));
	};

	// biome-ignore lint/suspicious/noExplicitAny: is checking a type
	const handleApiError = (err: any, cb?: (message: string) => void) => {
		if (err && isApiError(err.body)) {
			const errMsg = (err.body as ApiError).message;
			// handle generic errors
			if (errMsg === "Unauthorized") {
				addErrorMessage("No tiene permisos para acceder a este recurso");
				return;
			}
			cb?.(API_ERRORS_MAP[errMsg] ?? errMsg);
			return;
		}
		addErrorMessage("Ha ocurrido un error inesperado");
		console.error(err);
	}

	const setUser = (user: User) => {
		setAppData({ user });
	};

	const clearAppState = () => {
		setMessages([]);
		setAppData({ user: defaultUser });
	};

	const store = {
		messages,
		appData,
		addMessage,
		deleteMessage,
		addErrorMessage,
		addWarningMessage,
		addInfoMessage,
		addSuccessMessage,
		handleApiError,
		setUser,
		clearAppState,
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
