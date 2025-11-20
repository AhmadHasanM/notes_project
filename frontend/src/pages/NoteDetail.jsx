import { useEffect, useState } from "react";
import { useParams, Link } from "react-router-dom";
import axios from "axios";

export default function NoteDetail() {
  const { id } = useParams();
  const [note, setNote] = useState(null);
  const token = localStorage.getItem("token");

  useEffect(() => {
    if (!id || !token) return;

    axios
      .get(`http://localhost:8080/api/notes/${id}`, {
        headers: { Authorization: `Bearer ${token}` },
      })
      .then(res => setNote(res.data))
      .catch(err => console.log(err.response?.data || err.message));
  }, [id, token]);

  if (!token) return <div className="auth-container">Login dulu</div>;
  if (!note) return <div className="auth-container">Loading...</div>;

  return (
    <div className="app-container">
      <div className="note-detail-card">
        <h2 className="note-detail-title">{note.title}</h2>
        {note.image_url && (
          <img
            src={`http://localhost:8080/${note.image_url}`}
            alt={note.title}
            className="note-detail-image"
          />
        )}
        <p className="note-detail-content">{note.content}</p>
        <p className="note-detail-meta">
          <strong>Created at:</strong>{" "}
          {new Date(note.created_at).toLocaleString()}
        </p>
        <p className="note-detail-meta">
          <strong>Updated at:</strong>{" "}
          {new Date(note.updated_at).toLocaleString()}
        </p>
        <Link to="/notes" className="btn btn-primary back-btn">
          Back to Notes
        </Link>
      </div>
    </div>
  );
}
