x# Go E-Commerce (Mini Project)

This is a step-by-step learning project where I build a simple **e-commerce system** using **Golang**.  
The goal is to learn Go while gradually adding features like products, cart, checkout, and later a web API.

---

## 📂 Project Structure
```
go-ecom/
├── phase_1/
│   ├── main.go         # CLI app entrypoint
│   └── products.json   # Sample products
├── phase_2/            # (coming soon) Database integration
├── phase_3/            # (coming soon) REST API
├── phase_4/            # (coming soon) Frontend integration
├── phase_5/            # (coming soon) Cloud deployment
└── README.md
```


---

## 🛒 Features by Phase

### ✅ Phase 1: CLI Mini Shop
- Load products from `products.json`
- Basic in-memory cart
- Commands:
  - `list`
  - `add <id> <qty>`
  - `remove <id>`
  - `checkout`
  - `exit`

### 🔜 Phase 2: Database Storage
- Replace in-memory cart with **SQLite/Postgres**
- Persist products and orders

### 🔜 Phase 3: REST API
- Expose endpoints using **Go net/http**
- `GET /products`, `POST /cart`, `POST /checkout`

### 🔜 Phase 4: Frontend Integration
- Simple frontend with **React/Next.js**
- Connect to backend API

### 🔜 Phase 5: Cloud Deployment
- Dockerize the app
- Deploy on **AWS / GCP / Azure**
- CI/CD pipeline

---

## 🚀 How to Run (Phase 1)
```bash
cd phase_1
go run main.go
