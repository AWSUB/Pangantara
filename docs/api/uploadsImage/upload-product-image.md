### 12.2 Upload Product Image
**Definisi:** Upload foto produk. Maksimal 5MB, format JPG/PNG/WEBP.

**PATCH** `/upload/product/:id/image`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | multipart/form-data | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| id | UUID (path) | ✅ | ID produk |

**Form Data:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| file | file | ✅ | JPG, PNG, atau WEBP, maks 5MB |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Foto berhasil diupload |
| `400` | Tipe file tidak valid atau ukuran melebihi batas |
| `401` | Token tidak valid |
| `500` | Server error |

**Contoh Response (200):**
```json
{
    "success": true,
    "message": "Product image uploaded successfully",
    "data": {
        "file_url": "/uploads/images/products/uuid.jpg"
    }
}