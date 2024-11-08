
import { Component, effect, inject } from '@angular/core';
import { ApiService } from '../services/api.service'
import { CommonModule } from '@angular/common'
import { CardModule } from 'primeng/card';
import { CheckboxModule } from 'primeng/checkbox';
import { FormsModule } from '@angular/forms';
import { fetchItems, deleteItem } from '../state/items/items.actions';
import { Store } from '@ngrx/store';
import { ButtonModule } from 'primeng/button';
import { ProgressSpinnerModule } from 'primeng/progressspinner';
import { selectItems, selectItemsError, selectItemsLoading } from '../state/items/items.selectors';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [CommonModule, CardModule, CheckboxModule, FormsModule, ButtonModule, ProgressSpinnerModule],
  templateUrl: './home.component.html',
  styleUrl: './home.component.scss'
})
export class HomeComponent {
  private readonly store = inject(Store);

  data: any
  checked = false

  items = this.store.selectSignal(selectItems)
  itemsLoading = this.store.selectSignal(selectItemsLoading)
  itemsError = this.store.selectSignal(selectItemsError)

  constructor(private apiService: ApiService) {
    effect(() => {
      console.warn(this.items())
    })
  }

  ngOnInit(): void {
    this.store.dispatch(fetchItems.submit());
  }


  deleteItem(id: number) {
    this.store.dispatch(deleteItem.submit({ id: id }))
  }
}
