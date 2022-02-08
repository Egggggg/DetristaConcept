import { useParams } from "react-router-dom";

export default function Game() {
	const slug = useParams().slug;

	return <div> {slug} !!</div>;
}
