import { mande  } from "mande";

export const api = mande("http://0.0.0.0:4000");

export function setToken(token: string) {
	api.options.headers.Authorization = `Bearer ${token}`;
}

export function clearToken() {
	api.options.headers.Authorization = null;
}
