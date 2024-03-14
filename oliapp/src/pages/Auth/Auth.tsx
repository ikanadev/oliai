import { getToken } from "@/utils";
import { type ParentProps, onMount } from "solid-js";
import { useNavigate } from "@solidjs/router";

export default function Auth(props: ParentProps) {
	const navigate = useNavigate();
	console.log("render Auth");
	onMount(() => {
		const token = getToken();
		if (token) {
			navigate("/home", { replace: true });
		}
	});

	return (
		<div class="min-h-screen flex items-start md:items-center">
			<div class="flex-1 hidden md:block">
				<img src="/chatbot.svg" alt="chatbot" class="w-full object-fill" />
			</div>
			<div class="flex-1">
				<div class="flex-1 flex justify-center p-4">
					<div class="w-full sm:w-[400px] max-w-full mt-20 md:mt-0">
						{props.children}
					</div>
				</div>
			</div>
		</div>
	);
}
