export function GET(endpoint) {
	return new Promise((resolve, reject) => {
		fetch(endpoint, {
			method: "GET"
		})
			.then((res) => res.text())
			.then((text) => resolve(text))
			.catch((err) => {
				throw err;
			});
	});
}
