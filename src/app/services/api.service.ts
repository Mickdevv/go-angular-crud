import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Item } from '../models/todo.model';

@Injectable({
  providedIn: 'root'
})
export class ApiService {

  api_URL = "http://localhost:3000/api/items"

  constructor(private http: HttpClient) { }

  getData(): Observable<any> {
    return this.http.get<Item>(this.api_URL, { withCredentials: true })
  }
}
