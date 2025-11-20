import { Link, useNavigate } from "react-router-dom";

export default function Navbar() {
  const navigate = useNavigate();

  const logout = () => {
    localStorage.removeItem("token");
    navigate("/login");
  };

  return (
    <div className="navbar bg-base-100 shadow">
      <div className="flex-1">
        <Link to="/notes" className="btn btn-ghost normal-case text-xl">
          Notes App
        </Link>
      </div>

      <div className="flex-none gap-2">
        <Link to="/notes/create-note" className="btn btn-primary">+ Note</Link>

        <button className="btn btn-error" onClick={logout}>
          Logout
        </button>
      </div>
    </div>
  );
}
