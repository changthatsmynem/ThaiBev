export interface Product {
  id: number;
  code: string;
  barcode: string;
  created_at: string;
}

export interface ProductRequest {
  code: string;
}

export interface ApiResponse<T> {
  status: string;
  message: string;
  data: T;
}