import { Component, inject } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-product-list',
  standalone: true,
  imports: [FormsModule, CommonModule],
  templateUrl: './product-list.component.html',
  styleUrl: './product-list.css',
})
export class ProductListComponent {
  private readonly http = inject(HttpClient);
  private readonly apiUrl = 'http://localhost:8080/api/v1';

  newCode = '';
  products: { code: string }[] = [];
  showModal = false;
  modalMessage = '';

  getProduct() {
    this.http.get<{ code: string }[]>(`${this.apiUrl}/data`).subscribe({
      next: (data) => {
        this.products = data;
      },
      error: (err) => {
        console.error('Error fetching products:', err);
      }
    });
  }

  addProduct() {
    if (!this.newCode?.trim()) {
      return;
    }

    if (this.newCode.length === 19) {
      const productData = { code: this.newCode };

      this.http.post(`${this.apiUrl}/data`, productData).subscribe({
        next: () => {
          this.products.push(productData);
          this.newCode = '';
        },
        error: (err) => {
          console.error('Error adding product:', err);
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
  }

  closeModal() {
    this.showModal = false;
  }

  removeProduct(index: number) {
    if (index < 0 || index >= this.products.length) {
      console.error('Invalid index for product removal');
      return;
    }
    this.http.delete(`${this.apiUrl}/data/${index}`).subscribe({
      next: () => {
        this.products.splice(index, 1);
      },
      error: (err) => {
        console.error('Error removing product:', err);
      }
    });
    this.modalMessage = 'ต้องการลบข้อมูลรหัสสินค้า ' + this.products[index].code + ' หรือไม่?';
    this.showModal = true;
  }
}
