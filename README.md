# workshop-pwa-api

API สำหรับ Workshop PWA

## สิ่งที่ควรเตรียมสำหรับ Workshop

- **Postman** — สำหรับลองเรียก API ใน workshop  
  - ดาวน์โหลดได้ที่ [postman.com/downloads](https://www.postman.com/downloads/) (รองรับ Windows และ macOS)  
  - หรือใช้ Postman for Web ได้จาก [postman.com](https://www.postman.com/)  
  - หลังติดตั้งแล้ว ให้ **Import** ไฟล์ Postman collection เข้า Postman เพื่อโหลดชุด request ที่ใช้ใน workshop  
  **ชื่อไฟล์ที่ให้ import:** `PWA Training Workshop.postman_collection.json` (อยู่ในโฟลเดอร์โปรเจกต์)

## การรัน API บนเครื่องตัวเอง

มี 2 ทางเลือก — **วิธีที่ 1 (Docker)** ไม่ต้องติดตั้ง Go / **วิธีที่ 2 (Go)** ต้องติดตั้ง Go

---

### วิธีที่ 1: รันด้วย Docker (ไม่ต้องติดตั้ง Go)

เหมาะกับคนที่ไม่อยากติดตั้ง Go แต่มี **Docker Desktop** อยู่แล้ว (ดาวน์โหลดได้ที่ [docker.com/products/docker-desktop](https://www.docker.com/products/docker-desktop/) — รองรับ Windows และ macOS)

**สำคัญ:** ก่อนรันคำสั่ง `docker build` / `docker run` ต้อง**เปิด Docker Desktop ให้รันอยู่ก่อน** (ถ้าไม่เปิดจะขึ้นข้อความว่า "Cannot connect to the Docker daemon")

1. **ตั้งค่า environment**
   - คัดลอก `.env.example` เป็น `.env`
   - แก้ค่าใน `.env` ตามต้องการ (พอร์ต, Vallaris URL, API Key)

2. **Build และรัน**

   ```bash
   docker build -t workshop-pwa-api .
   docker run --env-file .env -p 1323:1323 workshop-pwa-api
   ```

   API จะรันที่ **http://localhost:1323**  
   (หยุด container: กด **Ctrl+C** ที่เทอร์มินัลที่รัน `docker run` อยู่)

---

### วิธีที่ 2: รันด้วย Go (ต้องติดตั้ง Go)

1. **ติดตั้ง Go** (ถ้ายังไม่มี)

   โปรเจกต์นี้รันได้ทั้งบน **Windows** และ **macOS** (รวมถึง Linux) ให้ติดตั้งตามระบบปฏิบัติการที่ใช้อยู่:

   | ระบบปฏิบัติการ | วิธีติดตั้ง |
   |----------------|-------------|
   | **Windows** | ดาวน์โหลด installer (.msi) จาก [go.dev/dl](https://go.dev/dl/) แล้วดับเบิลคลิกรันติดตั้ง |
   | **macOS** | ใช้ Homebrew: `brew install go` หรือดาวน์โหลด installer (.pkg) จาก [go.dev/dl](https://go.dev/dl/) |
   | **Linux** | ดาวน์โหลดจาก [go.dev/dl](https://go.dev/dl/) หรือใช้ package manager (เช่น Ubuntu: `sudo apt install golang-go`) |

   เปิด **Terminal** (macOS/Linux) หรือ **Command Prompt / PowerShell** (Windows) แล้วรัน `go version` เพื่อตรวจสอบ

2. **ตั้งค่า environment**
   - คัดลอก `.env.example` เป็น `.env`
   - แก้ค่าใน `.env` ตามต้องการ (พอร์ต, Vallaris URL, API Key)

3. **รัน API**

   ```bash
   go run .
   ```

   API จะรันที่ **http://localhost:1323** (หรือตามพอร์ตที่ตั้งใน `API_API_PORT`)

---

**การเชื่อมต่อจาก Frontend**
   - ตั้ง base URL ของ API เป็น: **`http://localhost:1323/api/1.0`**
   - ถ้า Frontend รันบนเครื่องอื่น ให้ใช้ IP ของเครื่องที่รัน API แทน `localhost`  
     ตัวอย่าง: `http://192.168.1.100:1323/api/1.0`

### Endpoints

| Method | Path                                                  | คำอธิบาย              |
| ------ | ----------------------------------------------------- | --------------------- |
| GET    | `/api/1.0/collections`                                | ดึงรายการ collections |
| GET    | `/api/1.0/collections/:collectionId`                  | ดึง collection ตาม ID |
| POST   | `/api/1.0/collections/:collectionId/items`            | สร้าง features        |
| GET    | `/api/1.0/collections/:collectionId/items`            | ดึง features          |
| GET    | `/api/1.0/collections/:collectionId/items/:featureId` | ดึง feature ตาม ID    |
| PUT    | `/api/1.0/collections/:collectionId/items`            | อัปเดต features       |
| DELETE | `/api/1.0/collections/:collectionId/items`            | ลบ features           |

**ค่า `collectionId` ที่รองรับ** (สำหรับ POST / PUT / DELETE ด้านบน):

| collectionId   | คำอธิบาย   |
|----------------|------------|
| `flow_meter`   | Flow Meter |
| `step_test`    | Step Test  |
| `dma_boundary` | DMA Boundary |
