import { Component, effect, inject } from '@angular/core';
import { AuthService } from '../../services/auth.service'
import { ReactiveFormsModule, FormControl, FormGroup, Validators } from '@angular/forms';
import { UserWithToken, UserLoginRequest } from '../../models/user.model'
import { Store } from '@ngrx/store';
import { login, logout } from '../../state/user/user.actions';
import { selectUserError, selectUserLoading, selectUserSuccess } from '../../state/user/user.selectors';
import { ProgressSpinnerModule } from 'primeng/progressspinner';
import { ButtonModule } from 'primeng/button';
import { CommonModule } from '@angular/common';
import { RouterModule, Router } from '@angular/router';
import { MessageService } from 'primeng/api';
import { ToastModule } from 'primeng/toast';
import { RippleModule } from 'primeng/ripple';
import { MessagesModule } from 'primeng/messages';
import { Message } from 'primeng/api';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [MessagesModule, RouterModule, ReactiveFormsModule, ProgressSpinnerModule, CommonModule, ButtonModule, ToastModule, RippleModule],
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
  constructor(private messageService: MessageService) {
    effect(() => {
      const err = this.loginError()
      if (err) {
        console.warn("Login error")
        this.messages = [
          { severity: 'error', detail: 'Error logging in' },
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
