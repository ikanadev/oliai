import type { ApiError } from "@/domain";

// biome-ignore lint/suspicious/noExplicitAny: is checking a type
export function isApiError(obj: any): obj is ApiError {
	return (obj && typeof obj === 'object' && obj.message && typeof obj.message === 'string');
}
