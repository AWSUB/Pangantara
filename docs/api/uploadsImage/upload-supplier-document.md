### 12.1 Upload Supplier Document
**Definisi:** Upload dokumen verifikasi supplier (NIB, Halal, atau Other). Maksimal 5MB, format PDF/JPG/PNG.

**PATCH** `/upload/supplier/:id/document`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | multipart/form-data | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| id | UUID (path) | ✅ | ID supplier |

**Form Data:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| document_type | string | ✅ | `nib`, `halal`, atau `other` |
| file | file | ✅ | PDF, JPG, atau PNG, maks 5MB |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Dokumen berhasil diupload |
| `400` | Tipe file tidak valid, ukuran melebihi batas, atau document_type salah |
| `401` | Token tidak valid |
| `500` | Server error |

**Contoh Response (200):**
```json
{
    "success": true,
    "message": "Document uploaded successfully",
    "data": {
        "document_type": "nib",
        "file_url": "/uploads/documents/supplier/uuid.pdf"
    }
}