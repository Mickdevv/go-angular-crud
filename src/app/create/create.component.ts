import { CommonModule } from '@angular/common';
import { Component, inject, Signal } from '@angular/core';
import { FormGroup, FormControl, ReactiveFormsModule } from '@angular/forms';
import { Store } from '@ngrx/store';
import { Item } from '../models/todo.model';
import { addItem } from '../state/items/items.actions';
import { UserWithToken } from '../models/user.model';
import { selectUserToken } from '../state/user/user.selectors';
import { Router } from '@angular/router';

@Component({
  selector: 'app-create',
  standalone: true,
  imports: [ReactiveFormsModule, CommonModule],
  templateUrl: './create.component.html',
  styleUrl: './create.component.scss'
})
export class CreateComponent {
  private readonly store = inject(Store);

  constructor(private router: Router) {
    // if (!this.userToken()) {
    //   this.router.navigate(['/login'])
    // }
  }

  userToken: Signal<any> = this.store.selectSignal(selectUserToken)

  createItem = new FormGroup({
    title: new FormControl(''),
    description: new FormControl(''),
    done: new FormControl(false),
  })

  submitItem() {
    const item: Item = {
      title: this.createItem.value.title || '',
      description: this.createItem.value.description || '',
      done: this.createItem.value.done || false
    }
    this.store.dispatch(addItem.submit({ item: item }))
  }
}
