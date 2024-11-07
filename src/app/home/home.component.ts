
import { Component } from '@angular/core';
import { ApiService } from '../services/api.service'
import { CommonModule } from '@angular/common'
import { CardModule } from 'primeng/card';
import { CheckboxModule } from 'primeng/checkbox';
import { FormsModule } from '@angular/forms';
import { fetchItems } from '../state/items/items.actions';
import { Store } from '@ngrx/store';
import { ButtonModule } from 'primeng/button';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [CommonModule, CardModule, CheckboxModule, FormsModule, ButtonModule],
  templateUrl: './home.component.html',
  styleUrl: './home.component.scss'
})
export class HomeComponent {

  data: any
  checked = false

  constructor(private apiService: ApiService, private store: Store) { }

  ngOnInit(): void {
    this.fetchData()
    this.store.dispatch(fetchItems.submit());
  }

  fetchData() {
    this.apiService.getData().subscribe({
      next: response => {
        this.data = response
      },
      error: err => {
        console.error("Error fetching data ", err)
      },
    })
  }
}
