export type TimeData = {
	createdAt: string,
	updatedAt: string,
	archivedAt: string | null,
	deletedAt: string | null,
};

export type Company = {
	id: string,
	name: string,
	logoUrl: string,
};

export type CompanyWithTimeData = Company & TimeData;

export type User = {
	id: string;
	firstName: string;
	lastName: string;
	email: string;
};
