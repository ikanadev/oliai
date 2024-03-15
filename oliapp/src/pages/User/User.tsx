import { type ParentProps, Show, onMount } from "solid-js";
import { useAppState } from "@/store";
import { removeToken, getToken, saveToken } from "@/utils";
import { useNavigate } from "@solidjs/router";
import { profile } from "@/api";
import { setToken, clearToken } from "@/api/mande";
import Layout from "./Layout";

export default function User(props: ParentProps) {
	const navigate = useNavigate();
	const { appData, setUser, addErrorMessage, clearAppState } = useAppState();

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
		<>
			<Show
				when={appData.user.id.length > 0}
				fallback={
					<div class="flex flex-col items-center py-20">
						<span class="loading loading-ring loading-lg" />
					</div>
				}
			>
				<Layout>
					{props.children}
				</Layout>
			</Show>
		</>
	);
}
