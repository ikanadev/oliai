import type { ParentProps } from "solid-js";

export default function Auth(props: ParentProps) {
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
