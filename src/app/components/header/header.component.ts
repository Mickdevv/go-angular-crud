import { Component, effect, inject, OnInit, Signal } from '@angular/core';
import { MenuItem } from 'primeng/api';
import { MenubarModule } from 'primeng/menubar';
import { BadgeModule } from 'primeng/badge';
import { AvatarModule } from 'primeng/avatar';
import { InputTextModule } from 'primeng/inputtext';
import { CommonModule } from '@angular/common';
import { RippleModule } from 'primeng/ripple';
import { logout } from '../../state/user/user.actions';
import { Store } from '@ngrx/store';
import { Router } from '@angular/router';
import { selectUserSuccess } from '../../state/user/user.selectors';

@Component({
    selector: 'app-header',
    standalone: true,
    imports: [MenubarModule, BadgeModule, AvatarModule, InputTextModule, RippleModule, CommonModule],
    templateUrl: './header.component.html',
    styleUrl: './header.component.scss'
})
export class HeaderComponent {

    items: MenuItem[] | undefined = [
        {
            label: 'Login',
            icon: 'pi pi-user',
            routerLink: ['/login']
        }, {
            label: 'Register',
            icon: 'pi pi-user',
            routerLink: ['/register']
        }
    ]

    private readonly store = inject(Store)
    loggedIn: Signal<any> = this.store.selectSignal(selectUserSuccess)

    constructor(private router: Router) {
        effect(() => {
            if (!this.loggedIn()) {
                this.items = [
                    {
                        label: 'Login',
                        icon: 'pi pi-user',
                        routerLink: ['/login']
                    }, {
                        label: 'Register',
                        icon: 'pi pi-user',
                        routerLink: ['/register']
                    }
                ]
            } else {
                this.items = [
                    {
                        label: 'Logout',
                        icon: 'pi pi-user',
                        command: () => this.logout()
                    },
                    {
                        label: 'List',
                        icon: 'pi pi-home',
                        routerLink: ['/']
                    },
                    {
                        label: 'Create',
                        icon: 'pi pi-bolt',
                        routerLink: ['/create']
                    }
                ]
            }
        })
    }

    logout() {
        console.warn('logout triggered')
        this.store.dispatch(logout.submit())
        this.router.navigate(['/login'])
    }
}