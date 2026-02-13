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

### ทางเลือกเพิ่ม: รัน API ที่เครื่องหลักแล้วเผยให้คนอื่นเรียกได้

กรณีรัน API ที่เครื่องหลักเครื่องเดียว แล้วต้องการให้คนอื่นใน workshop เรียก API ตัวเดียวกันได้ (คนละเครือข่ายหรือคนละที่) ให้ใช้เครื่องมือ tunnel จากอินเทอร์เน็ตเข้ามาที่ localhost เช่น **ngrok** หรือ **Cloudflare Tunnel**

| เครื่องมือ | วิธีใช้ |
|------------|---------|
| **ngrok** | ดาวน์โหลดจาก [ngrok.com/download](https://ngrok.com/download) แล้วรันคำสั่งด้านล่าง (ให้ API รันอยู่ที่พอร์ต 1323 ก่อน) — จะได้ URL แบบ `https://xxxx.ngrok-free.app` ใช้เป็น base ต่อท้าย `/api/1.0` ใน Postman / Frontend |
| **Cloudflare Tunnel** | ติดตั้ง `cloudflared` จาก [developers.cloudflare.com/cloudflare-one/connections/connect-apps/install-and-setup/installation](https://developers.cloudflare.com/cloudflare-one/connections/connect-apps/install-and-setup/installation) แล้วรันคำสั่งด้านล่าง — จะได้ URL แบบ `https://xxxx.trycloudflare.com` ใช้เป็น base ต่อท้าย `/api/1.0` เช่นกัน |

**คำสั่ง ngrok:**

```bash
ngrok http 1323
```

**คำสั่ง Cloudflare Tunnel:**

```bash
cloudflared tunnel --url http://localhost:1323
```

**หมายเหตุ:** คนที่รัน API ต้องเปิด tunnel ไว้ตลอดช่วง workshop และแจก URL ที่ได้ให้คนอื่นไปตั้งใน Postman (ตัวแปร `baseUrl`) หรือใน Frontend

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

#### ตัวอย่าง Request Body (จาก Postman)

**POST** `/api/1.0/collections/:collectionId/items` — สร้าง features (GeoJSON FeatureCollection รองรับ `flow_meter` | `step_test` | `dma_boundary` โครงสร้าง properties ต่างกันตาม collection):

```json
{
  "type": "FeatureCollection",
  "features": [
    {
      "type": "Feature",
      "properties": {
        "dma_id": 21,
        "dmaname": "ท่อส่ง ชนบท-โนนข่า",
        "dmano": "DMA11",
        "globalid": "01K3J57FB1VV3XX19MK67P12B1",
        "id": "68ad23c575dc2456f0d93ee6",
        "loggerid": 7606,
        "mmno": null,
        "pwacode": "5521012",
        "recorddate": "2025-11-10T11:49:18.664000+00:00",
        "remark": "TEST"
      },
      "geometry": {
        "coordinates": [
          [
            [102.57820975342231, 16.04552445775165],
            [102.57805644131889, 16.04554015944908],
            [102.57757436735893, 16.045562363136266],
            [102.57686246353228, 16.04549428794188],
            [102.57396075361486, 16.04519003440831],
            [102.57236570141683, 16.045012989995676],
            [102.57019396093997, 16.044799375253938],
            [102.56946391970166, 16.044789893055103],
            [102.5688646890991, 16.0448932074534],
            [102.56851766322505, 16.04500122333182],
            [102.56612044106285, 16.046064008618355],
            [102.56312826907411, 16.04741503503872],
            [102.55301582031628, 16.051972174455653],
            [102.54772792297285, 16.059886776670513],
            [102.54263742066607, 16.05664142466835],
            [102.54248934680268, 16.056479548561793],
            [102.5424650099133, 16.056451278060962],
            [102.54264786832131, 16.056354258051165],
            [102.5392344101586, 16.05269297741011],
            [102.53717934005329, 16.047544323713577],
            [102.53654523655057, 16.045880213288655],
            [102.53649979968199, 16.043123213619594],
            [102.53678722621308, 16.033303676600568],
            [102.53551425167859, 16.030542226495246],
            [102.53062634959448, 16.024329675937505],
            [102.54199093349912, 16.028016021316603],
            [102.53978494493636, 16.01974755885649],
            [102.54624474435549, 16.02330582120574],
            [102.55108344127302, 16.028439776308712],
            [102.56226757166355, 16.027412920348727],
            [102.57636949744595, 16.037232128358003],
            [102.5773585141714, 16.042657277622858],
            [102.57817686755506, 16.045369530366546],
            [102.57820975342231, 16.04552445775165]
          ]
        ],
        "type": "Polygon"
      }
    }
  ]
}
```

**PUT** `/api/1.0/collections/:collectionId/items` — อัปเดต features (แต่ละ feature ต้องมี `id`):

```json
{
  "type": "FeatureCollection",
  "features": [
    {
      "id": "698bee12dc2e44e83166766c",
      "type": "Feature",
      "geometry": {
        "type": "Polygon",
        "coordinates": [
          [
            [102.60638901365023, 16.0604179560052],
            [102.60623570154723, 16.0604336577026],
            [102.60575362758723, 16.0604558613898],
            [102.60504172376024, 16.0603877861954],
            [102.60214001384323, 16.0600835326619],
            [102.60054496164523, 16.0599064882492],
            [102.59837322116823, 16.0596928735075],
            [102.59764317993023, 16.0596833913086],
            [102.59704394932723, 16.0597867057069],
            [102.59669692345324, 16.0598947215854],
            [102.59429970129123, 16.0609575068719],
            [102.59130752930223, 16.0623085332923],
            [102.58119508054423, 16.066865672709202],
            [102.57633233356923, 16.0690603662612],
            [102.57081668089423, 16.071534922921902],
            [102.57066860703124, 16.0713730468153],
            [102.57064427014123, 16.0713447763145],
            [102.57082712854923, 16.0712477563047],
            [102.56741367038723, 16.0675864756637],
            [102.56535860028123, 16.0624378219671],
            [102.57056381418423, 16.0607127134989],
            [102.57400773299423, 16.0590128614696],
            [102.56758114906023, 16.046043280652],
            [102.56369351190723, 16.0454357247488],
            [102.55880560982322, 16.039223174191],
            [102.55925800092523, 16.0329660646731],
            [102.56796420516423, 16.03464105711],
            [102.57442400458423, 16.038199319459302],
            [102.57926270150124, 16.0433332745623],
            [102.58491987709623, 16.0508873755007],
            [102.60454875767424, 16.0521256266115],
            [102.60553777440023, 16.0575507758764],
            [102.60635612778323, 16.0602630286201],
            [102.60638901365023, 16.0604179560052]
          ]
        ]
      },
      "properties": {
        "_createdAt": "2026-02-11T02:48:50.184Z",
        "_createdBy": "698bea19e78db6ef27c5ed1a",
        "_createdat": "2025-08-26T03:02:29.985Z",
        "_createdby": "65792edcd715b5ab12310280",
        "_id": "698bee12dc2e44e83166766c",
        "_updatedAt": "2026-02-11T02:48:50.184Z",
        "_updatedBy": "698bea19e78db6ef27c5ed1a",
        "_updatedat": "2025-11-10T11:56:06.552Z",
        "_updatedby": "65792edcd715b5ab12310280",
        "dma_id": 21,
        "dmaname": "ท่อส่ง ชนบท-โนนข่า",
        "dmano": "DMA11",
        "globalid": "01K3J57FB1VV3XX19MK67P12B1",
        "id": "68ad23c575dc2456f0d93ee6",
        "loggerid": 7606,
        "mmno": null,
        "pwacode": "5521012",
        "recorddate": "2025-11-10T11:49:18.664000+00:00",
        "remark": "โรงกรองทุ่งมน"
      }
    }
  ]
}
```

**DELETE** `/api/1.0/collections/:collectionId/items` — ลบ features (ส่งเฉพาะ `id` ของแต่ละ feature):

```json
{
  "type": "FeatureCollection",
  "features": [
    {
      "id": "698c5563de33b2d04b955991"
    }
  ]
}
```
