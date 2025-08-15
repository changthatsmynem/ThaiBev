# Barcode Generation Guide

## Dependencies Required
Add to `go.mod`:
```bash
go get github.com/boombuler/barcode
```

## Project Structure
```
backend/thaibev-api/
├── models/
│   ├── product.go      # Product struct with barcode field
│   └── response.go     # Existing response models
├── services/
│   └── barcode.go      # Barcode generation service
└── handlers/
    └── api.go          # Updated API handlers
```

## API Endpoints

### POST /api/v1/data
**Request:**
```json
{
  "code": "1234-5678-9012-3456"
}
```

**Response:**
```json
{
  "status": "success",
  "message": "Product created with barcode",
  "data": {
    "barcode": "1234-5678-9012-3456",
    "image": "iVBORw0KGgoAAAANSUhEUgAA..." // base64 encoded PNG
  }
}
```

## Usage in Frontend
```typescript
// Display barcode image
<img [src]="'data:image/png;base64,' + product.image" alt="Barcode">
```

## Barcode Types Supported
- Code128 (default)
- Code39, EAN, QR codes available via different imports