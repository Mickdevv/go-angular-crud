import { Injectable } from '@angular/core';
import { UserWithToken, UserLoginRequest, UserRegisterRequest } from '../models/user.model'
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  base_url = "http://localhost:3000/api"

  constructor(private http: HttpClient) { }

  login(user: UserLoginRequest): Observable<UserWithToken> {
    return this.http.post<any>(`${this.base_url}/login/`, user, { withCredentials: true });
  }

  register(user: UserRegisterRequest): Observable<any> {
    return this.http.post<any>(`${this.base_url}/register/`, user, { withCredentials: true });
  }
}
