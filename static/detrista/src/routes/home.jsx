import { useState } from "react";
import { NavLink, Navigate } from "react-router-dom";
import { GET } from "../func";

function gameCodeChange(setCode) {
	return (e) => {
		setCode(e.target.value);
	};
}

function createGame(setCode, setRedirect) {
	return (e) => {
		GET("/create").then((res) => {
			setCode(res);
			setRedirect(true);
		});
	};
}

export default function Home() {
	const [code, setCode] = useState("");
	const [redirect, setRedirect] = useState(false);

	return (
		<div>
			{redirect && <Navigate to={`/play/${code}`} />}
			<button id="create-game" onClick={createGame(setCode, setRedirect)}>
				Create Game
			</button>
			<br />
			<input
				type="text"
				id="gamecode"
				placeholder="Game Code..."
				onChange={gameCodeChange(setCode)}
				value={code}
			/>
			<NavLink to={`/play/${code}`}>Join Game</NavLink>
		</div>
	);
}
