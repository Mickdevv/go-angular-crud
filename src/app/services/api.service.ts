import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
@Injectable({
  providedIn: 'root'
})
export class ApiService {

  api_URL = "http://localhost:3000/api/items"

  constructor(private http: HttpClient) {}

  getData(): Observable<any> {
    	return this.http.get<any>(this.api_URL)
  }
}
