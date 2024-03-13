import { Switch, Match, Show } from "solid-js";
import { Alert as TAlert } from "terracotta";
import { InfoRounded, WarningTriangle, ErrorRounded, CheckCircle, Close } from "@/icons";
import { MessageType } from "@/domain";

type Props = {
	message: string;
	variant?: MessageType;
	onClose?: VoidFunction;
}

export default function Alert(props: Props) {
	return (
		<TAlert
			class="alert shadow-lg pl-4 pr-2 py-1.5 gap-1 sm:gap-4"
			classList={{
				"alert-info": props.variant === MessageType.INFO,
				"alert-success": props.variant === MessageType.SUCCESS,
				"alert-warning": props.variant === MessageType.WARNING,
				"alert-error": props.variant === MessageType.ERROR,
			}}
		>
			<Switch fallback={<InfoRounded class="text-primary text-2xl" />}>
				<Match when={props.variant === MessageType.SUCCESS}>
					<CheckCircle class="text-2xl" />
				</Match>
				<Match when={props.variant === MessageType.INFO}>
					<InfoRounded class="text-2xl" />
				</Match>
				<Match when={props.variant === MessageType.WARNING}>
					<WarningTriangle class="text-2xl" />
				</Match>
				<Match when={props.variant === MessageType.ERROR}>
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
