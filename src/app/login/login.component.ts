import { Component } from '@angular/core';
import { AuthService } from '../services/auth.service'
import { ReactiveFormsModule, FormControl, FormGroup } from '@angular/forms';
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
    username: new FormControl(''),
    password: new FormControl('')
  });

  onSubmit() {
    user: User = {
      username: '',
      password: ''
    };
    console.log(this.loginForm.value)
    this.authService.login(this.loginForm.value)
  }
}
