import { Component, effect, inject } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Store } from '@ngrx/store';
import { selectItemsLoading, selectItemsError, selectItem } from '../state/items/items.selectors';
import { fetchItem } from '../state/items/items.actions';
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-edit',
  standalone: true,
  imports: [ReactiveFormsModule, CommonModule],
  templateUrl: './edit.component.html',
  styleUrl: './edit.component.scss'
})
export class EditComponent {
  private readonly store = inject(Store);

  data: any
  checked = false

  selectedItem = this.store.selectSignal(selectItem)
  itemsLoading = this.store.selectSignal(selectItemsLoading)
  itemsError = this.store.selectSignal(selectItemsError)

  editItem = new FormGroup({
    title: new FormControl(''),
    description: new FormControl(''),
    done: new FormControl(false),
  })

  constructor(private route: ActivatedRoute) {
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
}