import { Component } from '@angular/core';
import { AuthService } from '../services/auth.service'
import { ReactiveFormsModule, FormControl, FormGroup, Validators } from '@angular/forms';
import { User } from '../models/user.model'

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [ReactiveFormsModule],
  templateUrl: './login.component.html',
  styleUrl: './login.component.scss'
})
export class LoginComponent {

  constructor(private authService: AuthService) {}

  loginForm = new FormGroup({
    username: new FormControl('', [Validators.required, Validators.minLength(3)]),  // username must be at least 3 characters
    password: new FormControl('', [Validators.required, Validators.minLength(5)])   // password must be at least 6 characters
  });

  onSubmit() {
    if (this.loginForm.invalid) {
      return;
    }
  
    const user: User = {
      username: this.loginForm.value.username ?? '',  // Default to empty string if null
      password: this.loginForm.value.password ?? ''
    };
    
    console.log(this.loginForm.value)
    this.authService.login(user).subscribe({
      next: response => {
        console.log('Login success : ', response)
      },
      error: err => {
        console.error("Error logging in : ", err)
      },
    })
  }
}
