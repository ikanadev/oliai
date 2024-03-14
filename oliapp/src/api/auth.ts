import type { User } from "@/domain";
import { api } from "./mande";

export async function signUp(data: {
	firstName: string;
	lastName: string;
	email: string;
	password: string;
}) {
	return await api.post<void>("/signup", data);
}

export async function signIn(data: { email: string; password: string }) {
	return await api.post<{ token: string, user: User }>("/signin", data);
}

export async function profile() {
	return await api.get<{ token: string, user: User }>("/api/profile");
}
