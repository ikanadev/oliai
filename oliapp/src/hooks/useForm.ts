import { type Accessor, type JSX, createSignal } from "solid-js";
import type { JSXEvent } from "@/domain";

type FormFieldArg = {
	initialValue?: string;
	validators: ((value: string) => string)[];
}
type FormField = {
	value: Accessor<string>;
	onInput: JSX.InputEventHandlerUnion<HTMLInputElement, InputEvent>;
	error: Accessor<string>;
};

function getFormFields<T>(fields: { [key in keyof T]: FormFieldArg }): { [key in keyof T]: FormField } {
	const keys = Object.keys(fields) as (keyof T)[];
	return keys.reduce((acc, key) => {
		const [value, setValue] = createSignal(fields[key].initialValue ?? "");
		const [error, setError] = createSignal("");
		acc[key] = {
			value,
			error,
			onInput: (event: JSXEvent<InputEvent, HTMLInputElement>) => {
				const newValue = event.currentTarget.value;
				const validators = fields[key].validators;
				let error = "";
				for (const validator of validators) {
					error = validator(newValue);
					if (error.length > 0) {
						break;
					}
				}
				setError(error);
				setValue(newValue);
			}
		};
		return acc;
	}, {} as { [key in keyof T]: FormField });
}

export function useForm<T extends object>(fields: { [key in keyof T]: FormFieldArg }) {
	const form = getFormFields(fields);
	const isValid = () => {
		const keys = Object.keys(fields) as (keyof T)[];
		for (const key of keys) {
			const error = form[key].error();
			if (error.length > 0) {
				return false;
			}
		}
		return true;
	};

	return { form, isValid };
}
