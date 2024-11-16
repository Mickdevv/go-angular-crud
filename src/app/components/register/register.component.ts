import { Component, effect, inject } from '@angular/core';
import { ReactiveFormsModule, FormControl, FormGroup, Validators } from '@angular/forms';
import { AuthService } from '../../services/auth.service';
import { UserWithToken, UserRegisterRequest } from '../../models/user.model';
import { ButtonModule } from 'primeng/button';
import { RouterModule, Router } from '@angular/router';
import { selectUserError, selectUserLoading, selectUserSuccess } from '../../state/user/user.selectors';
import { Store } from '@ngrx/store';
import { ProgressSpinnerModule } from 'primeng/progressspinner';
import { CommonModule } from '@angular/common';
import { register } from '../../state/user/user.actions';
import { Message, MessageService } from 'primeng/api';
import { MessagesModule } from 'primeng/messages';


@Component({
  selector: 'app-register',
  standalone: true,
  imports: [MessagesModule, ReactiveFormsModule, ButtonModule, RouterModule, ProgressSpinnerModule, CommonModule],
  templateUrl: './register.component.html',
  styleUrl: './register.component.scss',
  providers: [MessageService]
})
export class RegisterComponent {

  constructor(private router: Router, private authService: AuthService) {
    effect(() => {
      const err = this.error()
      if (err) {
        console.warn(err)
        this.messages = [
          { severity: 'error', detail: err.error.error },
        ];
      }
    })
  }

  private readonly store = inject(Store)


  messages: Message[] = [];
  loading = this.store.selectSignal(selectUserLoading)
  loginSuccess = this.store.selectSignal(selectUserSuccess)
  error = this.store.selectSignal(selectUserError)

  registerForm = new FormGroup({
    username: new FormControl('', [Validators.required, Validators.minLength(3)]),  // username must be at least 3 characters
    password1: new FormControl('', [Validators.required, Validators.minLength(6)]),   // password must be at least 6 characters
    password2: new FormControl('', [Validators.required, Validators.minLength(6)])   // password must be at least 6 characters
  });

  onSubmit() {
    if (this.registerForm.invalid) {
      return;
    }

    const user: UserRegisterRequest = {
      username: this.registerForm.value.username ?? '',  // Default to empty string if null
      password1: this.registerForm.value.password1 ?? '',
      password2: this.registerForm.value.password2 ?? ''
    };

    this.store.dispatch(register.submit(user));
  }
}
