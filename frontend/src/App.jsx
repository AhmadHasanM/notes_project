import { BrowserRouter, Routes, Route, useLocation } from "react-router-dom";
import Login from "./pages/Login";
import Register from "./pages/Register";
import Notes from "./pages/Notes";
import CreateNote from "./pages/CreateNote";
import NoteDetail from "./pages/NoteDetail";
import EditNote from "./pages/EditNote";
import Navbar from "./components/Navbar";

function AppWrapper() {
  const location = useLocation();
  const hideNavbar = location.pathname === "/login" || location.pathname === "/register";

  return (
    <>
      {!hideNavbar && <Navbar />}
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route path="/notes" element={<Notes />} />
        <Route path="/notes/create-note" element={<CreateNote />} />
        <Route path="/notes/detail/:id" element={<NoteDetail />} />
        <Route path="/notes/update/:id" element={<EditNote />} />
      </Routes>
    </>
  );
}

export default function App() {
  return (
    <BrowserRouter>
      <AppWrapper />
    </BrowserRouter>
  );
}
