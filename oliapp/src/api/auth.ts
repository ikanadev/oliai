import { api } from "./mande";

export async function signUp(data: {
	firstName: string;
	lastName: string;
	email: string;
	password: string;
}) {
	return await api.post<number>("/signup", data);
}

export async function signIn(data: { email: string; password: string }) {
	return await api.post("/signin", data);
}
