import { Component } from '@angular/core';
import { ReactiveFormsModule, FormControl, FormGroup, Validators } from '@angular/forms';
import { AuthService } from '../../services/auth.service';
import { UserWithToken, UserRegisterRequest } from '../../models/user.model';
import { ButtonModule } from 'primeng/button';
import { RouterModule, Router } from '@angular/router';

@Component({
  selector: 'app-register',
  standalone: true,
  imports: [ReactiveFormsModule, ButtonModule, RouterModule],
  templateUrl: './register.component.html',
  styleUrl: './register.component.scss'
})
export class RegisterComponent {
  constructor(private router: Router, private authService: AuthService) {

  }

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

    this.authService.register(user).subscribe({
      next: response => {
        console.log('Register success : ', response)
      },
      error: err => {
        console.error("Error logging in : ", err)
      },
    })
  }
}
