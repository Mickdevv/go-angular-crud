import { Component, effect, inject } from '@angular/core';
import { AuthService } from '../services/auth.service'
import { ReactiveFormsModule, FormControl, FormGroup, Validators } from '@angular/forms';
import { UserWithToken, UserLoginRequest } from '../models/user.model'
import { Store } from '@ngrx/store';
import { login, logout } from '../state/user/user.actions';
import { selectUserLoading, selectUserSuccess } from '../state/user/user.selectors';
import { ProgressSpinnerModule } from 'primeng/progressspinner';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [ReactiveFormsModule, ProgressSpinnerModule, CommonModule],
  templateUrl: './login.component.html',
  styleUrl: './login.component.scss'
})
export class LoginComponent {

  private readonly store = inject(Store);

  loading = this.store.selectSignal(selectUserLoading)
  loginSuccess = this.store.selectSignal(selectUserSuccess)
  constructor(private authService: AuthService, private router: Router) {
    effect(() => {
      if (this.loginSuccess()) {
        this.router.navigate(['/'])

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
    // this.authService.login(user).subscribe({
    //   next: response => {
    //     console.log('Login success : ', response)
    //   },
    //   error: err => {
    //     console.error("Error logging in : ", err)
    //   },
    // })
  }
}
