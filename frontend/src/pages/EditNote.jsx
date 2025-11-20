import { useState, useEffect } from "react";
import { useParams, useNavigate } from "react-router-dom";
import axios from "axios";

export default function EditNote() {
  const { id } = useParams();
  const navigate = useNavigate();
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const [image, setImage] = useState(null); // file yang diupload
  const [preview, setPreview] = useState(null); // preview gambar lama / baru
  const token = localStorage.getItem("token");

  useEffect(() => {
    if (!token) return;

    axios
      .get(`http://localhost:8080/api/notes/${id}`, {
        headers: { Authorization: `Bearer ${token}` },
      })
      .then(res => {
        setTitle(res.data.title);
        setContent(res.data.content);
        if (res.data.image_url) setPreview(`http://localhost:8080/${res.data.image_url}`);
      })
      .catch(err => console.log(err.response?.data || err.message));
  }, [id, token]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (!token) return;

    const formData = new FormData();
    formData.append("title", title);
    formData.append("content", content);
    if (image) formData.append("image", image);

    try {
      await axios.put(
        `http://localhost:8080/api/notes/${id}`,
        formData,
        {
          headers: { 
            Authorization: `Bearer ${token}`,
            "Content-Type": "multipart/form-data"
          },
        }
      );
      navigate("/notes");
    } catch (err) {
      console.log(err.response?.data || err.message);
      alert("Gagal update note");
    }
  };

  const handleImageChange = (e) => {
    const file = e.target.files[0];
    if (file) {
      setImage(file);
      setPreview(URL.createObjectURL(file)); // preview langsung
    }
  };

  return (
    <div className="app-container">
      <div className="note-edit-card">
        <h2 className="note-edit-title">Edit Note</h2>
        <form onSubmit={handleSubmit} className="note-edit-form">
          <input
            type="text"
            placeholder="Title"
            value={title}
            onChange={e => setTitle(e.target.value)}
            required
          />
          <textarea
            placeholder="Content"
            value={content}
            onChange={e => setContent(e.target.value)}
            required
          />
          <input
            type="file"
            accept="image/*"
            onChange={handleImageChange}
          />
          {preview && (
            <img
              src={preview}
              alt="Preview"
              style={{ width: "100%", maxHeight: "200px", objectFit: "cover", marginTop: "10px", borderRadius: "5px" }}
            />
          )}
          <button type="submit" className="btn btn-primary">
            Update Note
          </button>
        </form>
      </div>
    </div>
  );
}
