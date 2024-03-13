export function emailValidator(email: string) {
	const regex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
	if (!regex.test(email)) {
		return "Ingresa un correo válido";
	}
	return "";
}

export function nonEmptyValidator(value: string) {
	if (value.trim().length === 0) {
		return "Este campo es requerido";
	}
	return "";
}

export function minLenValidator(len: number) {
	return (value: string) => {
		if (value.trim().length < len) {
			return `Debe tener al menos ${len} caracteres`;
		}
		return "";
	}
}
