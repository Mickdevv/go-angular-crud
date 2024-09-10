
import { Component } from '@angular/core';
import { ApiService } from '../api.service'
import { CommonModule } from '@angular/common'
import { CardModule } from 'primeng/card';
import { CheckboxModule } from 'primeng/checkbox';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [CommonModule, CardModule, CheckboxModule, FormsModule],
  templateUrl: './home.component.html',
  styleUrl: './home.component.scss'
})
export class HomeComponent {

  data: any
  checked = false

  constructor(private apiService: ApiService) {}
  
  ngOnInit(): void {
    this.fetchData()
  }

  fetchData() {
    this.apiService.getData().subscribe({
      next: response => {
        this.data = response
      },
      error: err => {
        console.error("Error fetching data ", err)
      },
    })
  }
}
