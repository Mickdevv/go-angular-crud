import { Component, inject } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { HeaderComponent } from './components/header/header.component'
import { FooterComponent } from './components/footer/footer.component';
import { Store } from '@ngrx/store';
import { CookieService } from 'ngx-cookie-service';
import { jwtDecode } from "jwt-decode";
import { login } from './state/user/user.actions';
import { UserWithToken } from './models/user.model';
@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, HeaderComponent, FooterComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent {
  private readonly store = inject(Store)
  constructor(private cookieService: CookieService) { }
  title = 'frontend';

  ngOnInit() {
    if (this.cookieService.check('jwt_token')) {
      const initialToken = this.cookieService.get('jwt_token')
      const decodedToken = jwtDecode(initialToken) as any
      const user: UserWithToken = { username: decodedToken.username, access: initialToken, refresh: initialToken }
      this.store.dispatch(login.success({ user }))
    }
  }
}
