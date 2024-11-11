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

  addItem(item: Item): Observable<any> {
    return this.http.post<Item>(this.api_URL, item, { withCredentials: true });
  }

  getItemById(id: number): Observable<any> {
    return this.http.get<Item>(`${this.api_URL}${id}/`, { withCredentials: true })
  }

  updateItem(item: Item): Observable<any> {
    return this.http.put<Item>(`${this.api_URL}${item.id}/`, item, { withCredentials: true });
  }

  deleteItem(id: number): Observable<any> {
    return this.http.delete<Number>(`${this.api_URL}${id}/`, { withCredentials: true });
  }
}
