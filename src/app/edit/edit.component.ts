import { Component, effect, inject, Signal } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Store } from '@ngrx/store';
import { selectItemsLoading, selectItemsError, selectItem } from '../state/items/items.selectors';
import { fetchItem, updateItem } from '../state/items/items.actions';
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { Item } from '../models/todo.model';
import { ApiService } from '../services/api.service';
import { selectUserToken } from '../state/user/user.selectors';

@Component({
  selector: 'app-edit',
  standalone: true,
  imports: [ReactiveFormsModule, CommonModule],
  templateUrl: './edit.component.html',
  styleUrl: './edit.component.scss'
})
export class EditComponent {
  private readonly store = inject(Store);

  userToken: Signal<any> = this.store.selectSignal(selectUserToken)
  selectedItem = this.store.selectSignal(selectItem)
  itemsLoading = this.store.selectSignal(selectItemsLoading)
  itemsError = this.store.selectSignal(selectItemsError)

  editItem = new FormGroup({
    title: new FormControl(''),
    description: new FormControl(''),
    done: new FormControl(false),
  })

  constructor(private route: ActivatedRoute, private itemsService: ApiService, private router: Router) {
    // if (!this.userToken()) {
    //   this.router.navigate(['/login'])
    // }
    effect(() => {
      const item = this.selectedItem()
      if (item) {
        this.editItem.patchValue({
          title: item.title,
          description: item.description,
          done: item.done
        })
      }
    })
  }

  ngOnInit() {

    const itemId = Number(this.route.snapshot.paramMap.get('id'));
    this.store.dispatch(fetchItem.submit({ id: itemId }))
  }

  submitUpdate() {
    if (this.selectedItem() != undefined) {
      const updatedItem: Item = {
        title: this.editItem.value.title || '',
        description: this.editItem.value.description || '',
        done: this.editItem.value.done || false,
        ownerId: this.selectedItem()!.ownerId,
        id: this.selectedItem()!.id,
      }
      this.store.dispatch(updateItem.submit({ item: updatedItem }))
    }
  }
}
