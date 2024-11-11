import { Component, inject, OnInit } from '@angular/core';
import { MenuItem } from 'primeng/api';
import { MenubarModule } from 'primeng/menubar';
import { BadgeModule } from 'primeng/badge';
import { AvatarModule } from 'primeng/avatar';
import { InputTextModule } from 'primeng/inputtext';
import { CommonModule } from '@angular/common';
import { RippleModule } from 'primeng/ripple';
import { logout } from '../state/user/user.actions';
import { Store } from '@ngrx/store';
import { Router } from '@angular/router';

@Component({
    selector: 'app-header',
    standalone: true,
    imports: [MenubarModule, BadgeModule, AvatarModule, InputTextModule, RippleModule, CommonModule],
    templateUrl: './header.component.html',
    styleUrl: './header.component.scss'
})
export class HeaderComponent {

    private readonly store = inject(Store)

    constructor(private router: Router) { }

    items: MenuItem[] | undefined;
    itemsEnd: MenuItem[] | undefined;

    ngOnInit() {
        this.items = [
            {
                label: 'List',
                icon: 'pi pi-home',
                routerLink: ['/']
            },
            {
                label: 'Create',
                icon: 'pi pi-bolt',
                routerLink: ['/create']
            },
            // {
            //     label: 'To do',
            //     icon: 'pi pi-search',
            //     items: [
            //         {
            //             label: 'Create',
            //             icon: 'pi pi-bolt',
            //             shortcut: '⌘+S',
            //             routerLink: ['/create']
            //         },
            //         {
            //             label: 'Update',
            //             icon: 'pi pi-server',
            //             shortcut: '⌘+B',
            //             routerLink: ['/create']
            //         },
            //         {
            //             label: 'Delete',
            //             icon: 'pi pi-pencil',
            //             shortcut: '⌘+U',
            //             routerLink: ['/create']
            //         },
            //         {
            //             separator: true
            //         },
            //         {
            //             label: 'Templates',
            //             icon: 'pi pi-palette',
            //             items: [
            //                 {
            //                     label: 'Item1',
            //                     icon: 'pi pi-palette',
            //                     badge: '2',
            //                     routerLink: ['/create']
            //                 },
            //                 {
            //                     label: 'Item2',
            //                     icon: 'pi pi-palette',
            //                     badge: '3',
            //                     routerLink: ['/create']
            //                 }
            //             ]
            //         }
            //     ]
            // },
            {
                label: 'Login',
                icon: 'pi pi-user',
                routerLink: ['/login']
            },
        ];
    }

    logout() {
        this.store.dispatch(logout.submit())

        this.router.navigate(['/login'])
    }
}