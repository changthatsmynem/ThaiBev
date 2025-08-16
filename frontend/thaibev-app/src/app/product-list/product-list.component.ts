import { Component, inject, OnInit, ChangeDetectorRef } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { HttpClient } from '@angular/common/http';
import { Product, ApiResponse } from '../models/product.model';
import { maybe } from '../utils/maybe';

@Component({
  selector: 'app-product-list',
  standalone: true,
  imports: [FormsModule, CommonModule],
  templateUrl: './product-list.component.html',
  styleUrl: './product-list.css',
})
export class ProductListComponent implements OnInit {
  private readonly http = inject(HttpClient);
  private readonly cdr = inject(ChangeDetectorRef);
  private readonly apiUrl = 'http://localhost:8080/api/products';

  newCode = '';
  products: Product[] = [];
  showModal = false;
  modalMessage = '';
  selectedProduct: Product | null = null;
  isDuplicate = false;
  errorMessage = '';

  ngOnInit() {
    this.getProduct();
  }

  getProduct() {
    this.http.get<ApiResponse<Product[]>>(`${this.apiUrl}`).subscribe({
      next: (response) => {
        this.products = maybe(response?.data).getOrElse([]);
        this.cdr.detectChanges();
      },
      error: (err) => {
        console.error('Error fetching products:', err);
      }
    });
  }

  addProduct() {
    const code = this.newCode?.trim() || '';
    this.isDuplicate = false;
    this.errorMessage = '';

    if (code.length >= 16) {
      this.http.post<any>(`${this.apiUrl}`, { code }).subscribe({
        next: () => {
          this.newCode = '';
          this.getProduct();
        },
        error: (err) => {
          if (err.error?.status?.includes('error')) {
            this.isDuplicate = true;
            this.errorMessage = 'รหัสสินค้านี้มีอยู่แล้ว';
            this.cdr.markForCheck();
            this.cdr.detectChanges();
          } else {
            console.error('Error creating product:', err);
          }
        }
      });
    }
  }

  formatInput(event: any) {
    let value = event.target.value.replace(/[^A-Za-z0-9]/g, '').toUpperCase();
    if (value.length > 16) value = value.substring(0, 16);

    const formatted = value.replace(/(\w{4})(?=\w)/g, '$1-');
    this.newCode = formatted;
    event.target.value = formatted;

    this.isDuplicate = false;
    this.errorMessage = '';
  }

  openModal(product: Product) {
    this.selectedProduct = product;
    this.showModal = true;
    this.modalMessage = `ต้องการลบข้อมูล รหัสสินค้า ${product.code} หรือไม่?`;
  }

  closeModal() {
    this.showModal = false;
    this.selectedProduct = null;
  }

  confirmDelete() {
    if (this.selectedProduct) {
      this.removeProduct(this.selectedProduct.id);
    }
  }

  removeProduct(id: number) {
    this.http.delete(`${this.apiUrl}/${id}`).subscribe({
      next: () => {
        this.showModal = false;
        this.getProduct();
      },
      error: (err) => {
        console.error('Error removing product:', err);
      }
    });
  }
}
