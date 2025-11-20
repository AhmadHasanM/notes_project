import axios from "axios";

const api = axios.create({
  baseURL: "http://notes_backend:8080/api", // ganti localhost jadi nama service Docker
  withCredentials: true,
});

api.interceptors.request.use((config) => {
  const token = localStorage.getItem("token");
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export default api;
