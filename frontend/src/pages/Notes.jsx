import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import axios from "axios";

export default function Notes() {
  const [notes, setNotes] = useState([]);
  const token = localStorage.getItem("token");

  const fetchNotes = () => {
    if (!token) return;

    axios.get("http://localhost:8080/api/notes", {
      headers: { Authorization: `Bearer ${token}` },
    })
    .then(res => setNotes(res.data))
    .catch(err => console.log(err.response?.data || err.message));
  };

  useEffect(() => {
    fetchNotes();
  }, [token]);

  const handleDelete = (id) => {
    if (!window.confirm("Yakin mau hapus note ini?")) return;

    axios.delete(`http://localhost:8080/api/notes/${id}`, {
      headers: { Authorization: `Bearer ${token}` },
    })
    .then(() => fetchNotes())
    .catch(err => console.log(err.response?.data || err.message));
  };

  if (!token) return <div className="app-container">Login dulu untuk lihat notes</div>;
  if (!notes.length) return <div className="app-container">Loading...</div>;

  return (
    <div className="page-container">
      <div className="notes-container">
        {notes.map(note => (
          <div key={note.id} className="note-card">
            {note.image_url && (
            <img 
              src={`http://localhost:8080/${note.image_url}`} 
              alt={note.title} 
              style={{ width: "100%", height: "150px", objectFit: "cover", borderRadius: "5px" }}
            />
          )}
            <h3>{note.title}</h3>
            <p>{note.content}</p>
            <div className="note-buttons">
              <Link to={`/notes/detail/${note.id}`} className="btn btn-sm btn-info">Detail</Link>
              <Link to={`/notes/update/${note.id}`} className="btn btn-sm btn-warning">Edit</Link>
              <button className="btn btn-sm btn-error" onClick={() => handleDelete(note.id)}>Delete</button>
            </div>
          </div>
        ))}
      </div>
    </div>
  );

}
