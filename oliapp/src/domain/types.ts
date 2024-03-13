export type ApiError = {
	error: string;
	message: string;
};

export type JSXEvent<Ev, El> = Ev & { currentTarget: El, target: Element };
