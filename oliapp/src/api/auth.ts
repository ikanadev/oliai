import { api } from "./mande";

export async function signUp(data: {
	firstName: string;
	lastName: string;
	email: string;
	password: string;
}) {
	return await api.post("/signup", data);
}
