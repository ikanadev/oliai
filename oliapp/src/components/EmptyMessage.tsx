import { mergeProps } from "solid-js";
import { InfoRounded } from "@/icons";

export default function EmptyMessage(props: { message?: string }) {
	const mergedProps = mergeProps({ message: "No hay resultados." }, props);
	return (
		<div class="flex justify-center py-10 items-center gap-2 opacity-60">
			<InfoRounded class="text-2xl" />
			<p class="italic">{mergedProps.message}</p>
		</div>
	);
}
