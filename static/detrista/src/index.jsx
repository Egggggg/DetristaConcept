import React from "react";
import ReactDOM from "react-dom";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import App from "./components/app";
import Home from "./routes/home";
import Game from "./routes/game";

ReactDOM.render(
	<BrowserRouter>
		<Routes>
			<Route path="/" element={<App />}>
				<Route index element={<Home />} />
				<Route path="/play/:slug" element={<Game />} />
			</Route>
		</Routes>
	</BrowserRouter>,
	document.getElementById("root")
);
