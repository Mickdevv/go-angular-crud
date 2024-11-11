
import { Component, effect, inject, Signal } from '@angular/core';
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
import { Router } from '@angular/router';
import { selectUserToken } from '../state/user/user.selectors';

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

  userToken: Signal<any> = this.store.selectSignal(selectUserToken)
  items = this.store.selectSignal(selectItems)
  itemsLoading = this.store.selectSignal(selectItemsLoading)
  itemsError = this.store.selectSignal(selectItemsError)

  constructor(private router: Router) {
    // if (!this.userToken().access) {
    //   this.router.navigate(['/login']);
    // }
    effect(() => {
      console.warn(this.items())
    })
  }

  ngOnInit(): void {
    this.store.dispatch(fetchItems.submit());
  }

  editItem(id: number) {
    this.router.navigate(['/edit', id]);
  }

  deleteItem(id: number) {
    this.store.dispatch(deleteItem.submit({ id: id }))
  }
}
