import { Switch, Match, Show } from "solid-js";
import { Alert as TAlert } from "terracotta";
import { InfoRounded, WarningTriangle, ErrorRounded, CheckCircle, Close } from "@/icons";

type Props = {
	message: string;
	variant?: "info" | "success" | "warning" | "error";
	onClose?: VoidFunction;
}

export default function Alert(props: Props) {
	return (
		<TAlert
			class="alert shadow-lg pl-4 pr-2 py-2"
			classList={{
				"alert-info": props.variant === "info",
				"alert-success": props.variant === "success",
				"alert-warning": props.variant === "warning",
				"alert-error": props.variant === "error",
			}}
		>
			<Switch fallback={<InfoRounded class="text-primary text-2xl" />}>
				<Match when={props.variant === "success"}>
					<CheckCircle class="text-2xl" />
				</Match>
				<Match when={props.variant === "info"}>
					<InfoRounded class="text-2xl" />
				</Match>
				<Match when={props.variant === "warning"}>
					<WarningTriangle class="text-2xl" />
				</Match>
				<Match when={props.variant === "error"}>
					<ErrorRounded class="text-2xl" />
				</Match>
			</Switch>
			{props.message}
			<Show when={props.onClose}>
				{(onClose) => (
					<button type="button" onClick={onClose()} class="btn btn-circle btn-ghost">
						<Close class="text-2xl" />
					</button>
				)}
			</Show>
		</TAlert>
	);
}
