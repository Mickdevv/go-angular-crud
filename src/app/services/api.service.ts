import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { Item } from '../models/todo.model';

@Injectable({
  providedIn: 'root'
})
export class ApiService {

  api_URL = "http://localhost:3000/api/items/"

  constructor(private http: HttpClient) { }

  getItems(): Observable<any> {
    return this.http.get<Item>(this.api_URL, { withCredentials: true })
  }

  getItemById(id: string): Observable<any> {
    return this.http.get<Item>(this.api_URL + id + "/", { withCredentials: true })
  }

  deleteItem(id: number): Observable<any> {
    // return of(1);
    return this.http.delete<Number>(this.api_URL + id + "/", { withCredentials: true });
  }
}
