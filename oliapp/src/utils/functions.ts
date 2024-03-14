import type { ApiError } from "@/domain";

// biome-ignore lint/suspicious/noExplicitAny: is checking a type
export function isApiError(obj: any): obj is ApiError {
	return (obj && typeof obj === 'object' && obj.message && typeof obj.message === 'string');
}

const TOKEN_KEY = "token";
export function saveToken(token: string) {
	localStorage.setItem(TOKEN_KEY, token);
}
export function getToken() {
	return localStorage.getItem(TOKEN_KEY);
}
export function removeToken() {
	localStorage.removeItem(TOKEN_KEY);
}
