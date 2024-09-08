import { Routes } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { ToDoComponent } from './to-do/to-do.component';
import { LoginComponent } from './login/login.component';
import { RegisterComponent } from './register/register.component';

export const routes: Routes = [
    { path: '', component: HomeComponent },
    { path: 'to-do', component: ToDoComponent },
    { path: 'login', component: LoginComponent },
    { path: 'register', component: RegisterComponent },
];
