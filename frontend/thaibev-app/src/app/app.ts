import { Component, signal } from '@angular/core';
import { ProductListComponent } from './product-list/product-list.component';

@Component({
  selector: 'app-root',
  imports: [ProductListComponent],
  template: '<app-product-list></app-product-list>',
  styleUrl: './app.css'
})
export class App {
  protected readonly title = signal('thaibev-app');
}
