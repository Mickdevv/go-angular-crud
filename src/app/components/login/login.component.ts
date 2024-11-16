import { Component, effect, inject } from '@angular/core';
import { ReactiveFormsModule, FormControl, FormGroup, Validators } from '@angular/forms';
import { UserLoginRequest } from '../../models/user.model'
import { Store } from '@ngrx/store';
import { login, logout } from '../../state/user/user.actions';
import { selectUserError, selectUserLoading, selectUserSuccess } from '../../state/user/user.selectors';
import { ProgressSpinnerModule } from 'primeng/progressspinner';
import { ButtonModule } from 'primeng/button';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { MessageService } from 'primeng/api';
import { RippleModule } from 'primeng/ripple';
import { MessagesModule } from 'primeng/messages';
import { Message } from 'primeng/api';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [MessagesModule, RouterModule, ReactiveFormsModule, ProgressSpinnerModule, CommonModule, ButtonModule, RippleModule],
  templateUrl: './login.component.html',
  styleUrl: './login.component.scss',
  providers: [MessageService]
})
export class LoginComponent {

  private readonly store = inject(Store);

  messages: Message[] = [];
  loading = this.store.selectSignal(selectUserLoading)
  loginSuccess = this.store.selectSignal(selectUserSuccess)
  loginError = this.store.selectSignal(selectUserError)
  constructor() {
    effect(() => {
      const err = this.loginError()
      if (err) {
        console.warn(err)
        this.messages = [
          { severity: 'error', detail: err.error.error },
        ];
      }
    })
  }

  loginForm = new FormGroup({
    username: new FormControl('', [Validators.required, Validators.minLength(3)]),  // username must be at least 3 characters
    password: new FormControl('', [Validators.required, Validators.minLength(5)])   // password must be at least 6 characters
  });

  onSubmit() {
    if (this.loginForm.invalid) {
      return;
    }

    const user: UserLoginRequest = {
      username: this.loginForm.value.username ?? '',  // Default to empty string if null
      password: this.loginForm.value.password ?? ''
    };

    console.log(this.loginForm.value)
    this.store.dispatch(login.submit({ user }));
  }
}
