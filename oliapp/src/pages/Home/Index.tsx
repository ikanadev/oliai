import { type ParentProps, Show, onMount } from "solid-js";
import { useAppState } from "@/store";
import { removeToken, getToken, saveToken } from "@/utils";
import { useNavigate } from "@solidjs/router";
import { profile } from "@/api";
import { setToken, clearToken } from "@/api/mande";

export default function Index(props: ParentProps) {
	const navigate = useNavigate();
	const { appData, setUser, addErrorMessage, clearAppState } = useAppState();

	const logout = () => {
		removeToken();
		clearToken();
		clearAppState();
		navigate("/auth/signin", { replace: true });
	};

	onMount(() => {
		if (appData.user.id.length > 0) {
			return;
		}
		const token = getToken();
		if (!token) {
			navigate("/auth/signin", { replace: true });
			return;
		}
		setToken(token);
		profile().then((data) => {
			saveToken(data.token);
			setUser(data.user);
		}).catch((err) => {
			console.error(err);
			removeToken();
			clearToken();
			clearAppState();
			addErrorMessage("Error restaurando sesión, por favor vuelve a iniciar sesión");
			navigate("/auth/signin", { replace: true });
		});
	});
	return (
		<div class="flex justify-center">
			<h1 class="my-20">INDEX</h1>
			<button type="button" onClick={logout}>logout</button>
			<Show when={appData.user.id.length > 0} fallback={<p>Loading...</p>}>
				{props.children}
			</Show>
		</div>
	);
}
