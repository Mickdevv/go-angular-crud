import { Component, OnInit } from '@angular/core';
import { MenuItem } from 'primeng/api';
import { MenubarModule } from 'primeng/menubar';
import { BadgeModule } from 'primeng/badge';
import { AvatarModule } from 'primeng/avatar';
import { InputTextModule } from 'primeng/inputtext';
import { CommonModule } from '@angular/common';
import { RippleModule } from 'primeng/ripple';

@Component({
  selector: 'app-header',
  standalone: true,
  imports: [MenubarModule, BadgeModule, AvatarModule, InputTextModule, RippleModule, CommonModule],
  templateUrl: './header.component.html',
  styleUrl: './header.component.scss'
})
export class HeaderComponent {
  items: MenuItem[] | undefined;
  itemsEnd: MenuItem[] | undefined;

  ngOnInit() {
      this.items = [
          {
              label: 'Home',
              icon: 'pi pi-home',
              routerLink: ['/']
          },
          {
              label: 'To do',
              icon: 'pi pi-search',
              items: [
                  {
                      label: 'Create',
                      icon: 'pi pi-bolt',
                      shortcut: '⌘+S',
                      routerLink: ['/create']
                  },
                  {
                      label: 'Update',
                      icon: 'pi pi-server',
                      shortcut: '⌘+B',
                      routerLink: ['/create']
                  },
                  {
                      label: 'Delete',
                      icon: 'pi pi-pencil',
                      shortcut: '⌘+U',
                      routerLink: ['/create']
                  },
                  {
                      separator: true
                  },
                  {
                      label: 'Templates',
                      icon: 'pi pi-palette',
                      items: [
                          {
                              label: 'Item1',
                              icon: 'pi pi-palette',
                              badge: '2',
                              routerLink: ['/create']
                          },
                          {
                              label: 'Item2',
                              icon: 'pi pi-palette',
                              badge: '3',
                              routerLink: ['/create']
                          }
                      ]
                  }
              ]
          },
          {
            label: 'Login',
            icon: 'pi pi-user',
            routerLink: ['/login']
        },
      ];
    }
}