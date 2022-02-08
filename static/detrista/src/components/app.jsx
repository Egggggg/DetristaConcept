import { NavLink, Outlet } from "react-router-dom";

export default function App() {
	return (
		<div>
			<nav id="menu">
				<NavLink to="/">Home</NavLink>
			</nav>
			<Outlet />
		</div>
	);
}
