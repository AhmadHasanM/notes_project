# 📝 Notes Project  
A simple full-stack Notes Application built with **Go (Golang)**, **React (Vite)**, **PostgreSQL**, and **Docker**.

---

## 🚀 Tech Stack

### **Backend – Golang**
- Go Fiber framework
- PostgreSQL (pgxpool)
- JWT Authentication
- MVC folder structure
- Dockerized backend

### **Frontend – React (Vite)**
- React + Vite
- Axios for API requests
- TailwindCSS for UI
- Protected Route system
- Dockerized frontend

### **Database**
- PostgreSQL  
- Automatic migration (on start)

### **Deployment / DevOps**
- Docker & Docker Compose

---

## 📂 Project Structure

### **Backend Folder Structure**
Backend berada di folder `/backend`:

backend/
│── config/
│ ├── config.go # Load environment variables (.env)
│ ├── database.go # PostgreSQL connection (pgxpool)
│ └── jwt.go # JWT token generator + validator
│
│── controllers/
│ ├── auth_controller.go
│ └── notes_controller.go
│
│── middlewares/
│ └── auth_middleware.go
│
│── models/
│ ├── notes.go
│ └── users.go
│
│── routes/
│ └── routes.go # API route mapping
│
│── main.go # App entry point
│── go.mod
│── go.sum
│── .env

markdown
Copy code

### **Frontend Folder Structure**
Frontend berada di `/frontend`:

- `src/api` → `api.js` untuk request Axios  
- `src/components` → `Navbar.jsx`, `ProtectedRoute.jsx`  
- `src/pages` → Login, Register, Notes, Detail, CRUD  
- Konfigurasi: `vite.config.js`, `tailwind.config.js`

---

## 🔐 Environment Variables

### **Backend `.env`**
PORT=8080
DB_URL=postgres://admin:admin123@db:5432/notesdb
JWT_SECRET=your-secret-key

markdown
Copy code

### **Frontend `.env`**
VITE_API_URL=http://localhost:8080

yaml
Copy code

---

## 🐳 Run with Docker Compose

Pastikan berada di folder `/infra`:

docker-compose build
docker-compose up

yaml
Copy code

Services:
- Backend → `http://localhost:8080`
- Frontend → `http://localhost:5173`
- PostgreSQL → port 5432

---

## ▶️ Run Locally (tanpa Docker)

### Backend
cd backend
go mod tidy
go run main.go

shell
Copy code

### Frontend
cd frontend
npm install
npm run dev

yaml
Copy code

---

## 📌 API Endpoints (Ringkas)

### **Auth**
| Method | Route | Description |
|--------|--------|-------------|
| POST | `/auth/register` | Register user |
| POST | `/auth/login` | Login & get JWT |

### **Notes**
| Method | Route | Description |
|--------|--------|-------------|
| GET | `/notes` | List all notes |
| POST | `/notes` | Create note |
| GET | `/notes/:id` | Detail note |
| PUT | `/notes/:id` | Update note |
| DELETE | `/notes/:id` | Delete note |

---

## 🧑‍💻 Author
**Ahmad Hasan**  
Full-Stack Developer  

---

Jika kamu ingin versi:
✨ Lebih profesional  
✨ Dengan badge GitHub (build passing, technologies, license, dll)  
✨ Dengan screenshot UI  

Tinggal bilang!

<img width="1894" height="969" alt="image" src="https://github.com/user-attachments/assets/dbe52f35-de83-4737-829f-008452f3aaf4" />
<img width="1892" height="972" alt="image" src="https://github.com/user-attachments/assets/d248cb5f-2053-4bf8-afac-e54252c6aa1f" />
<img width="1919" height="973" alt="image" src="https://github.com/user-attachments/assets/32fc24b9-b8ee-4c8f-aa17-8fe5e136e738" />
<img width="590" height="953" alt="image" src="https://github.com/user-attachments/assets/0a9f010a-e344-460c-9c59-fe964d38df9a" />
<img width="1884" height="867" alt="image" src="https://github.com/user-attachments/assets/9afd86a3-96c7-4c19-b0ef-fe7dc2c829f0" />
<img width="1912" height="762" alt="image" src="https://github.com/user-attachments/assets/6cc84e68-5b1f-4f89-8499-0debb1753878" />


