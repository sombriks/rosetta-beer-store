const baseUrl = import.meta.env.VITE_API_URL;

export const listBeers = ({ search = '', page = 1, pageSize = 10 }) => {
	const url = `${baseUrl}/beer/list?search=${search}&page=${page}&pageSize=${pageSize}`;
	return fetch(url).then((ret) => ret.json());
};
