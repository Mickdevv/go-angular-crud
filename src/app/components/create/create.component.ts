import { CommonModule } from '@angular/common';
import { Component, inject, Signal } from '@angular/core';
import { FormGroup, FormControl, ReactiveFormsModule, Validators } from '@angular/forms';
import { Store } from '@ngrx/store';
import { Item } from '../../models/todo.model';
import { addItem } from '../../state/items/items.actions';
import { UserWithToken } from '../../models/user.model';
import { selectUserSuccess, selectUserToken } from '../../state/user/user.selectors';
import { Router } from '@angular/router';
import { ButtonModule } from 'primeng/button';


@Component({
  selector: 'app-create',
  standalone: true,
  imports: [ReactiveFormsModule, CommonModule, ButtonModule],
  templateUrl: './create.component.html',
  styleUrl: './create.component.scss'
})
export class CreateComponent {
  private readonly store = inject(Store);

  constructor(private router: Router) {
    if (!this.loginSuccess()) {
      console.warn('Routing from create to login')
      this.router.navigate(['/login'])
    }
  }

  loginSuccess: Signal<any> = this.store.selectSignal(selectUserSuccess)

  createItem = new FormGroup({
    title: new FormControl('', [Validators.required, Validators.minLength(3)]),
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
