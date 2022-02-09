function GET(endpoint) {
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

document.addEventListener("DOMContentLoaded", () => {
	function createGame() {
		GET("/create").then((res) => {
			window.location.href = `/play/${res}`;
		});
	}

	function joinGame() {
		const gameCode = document.getElementById("gamecode").value;
		window.location.href = `/play/${gameCode}`;
	}

	const createGameButton = document.getElementById("creategame");
	const joinGameButton = document.getElementById("joingame");

	createGameButton.onclick = createGame;
	joinGameButton.onclick = joinGame;
});
