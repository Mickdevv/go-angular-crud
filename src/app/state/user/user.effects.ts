import { inject, Injectable } from "@angular/core";
import { Actions, createEffect, ofType } from "@ngrx/effects";
import { ApiService } from "../../services/api.service";
import { AuthService } from "../../services/auth.service";
import { login, logout, register } from "./user.actions";
import { catchError, delay, map, of, switchMap, tap } from "rxjs";
import { Router } from "@angular/router";
import { CookieService } from 'ngx-cookie-service';

@Injectable()
export class UserEffects {
    private readonly actions = inject(Actions);
    private readonly authService = inject(AuthService);
    private readonly cookieService = inject(CookieService)
    private readonly router = inject(Router)

    login = createEffect(() =>
        this.actions.pipe(
            ofType(login.submit),
            switchMap(({ user }) =>
                this.authService.login(user).pipe(
                    delay(1000),
                    map((userWithToken) => {
                        this.cookieService.set('jwt_token', userWithToken.access);
                        return (login.success({ user: userWithToken }))
                    }
                    ),
                    catchError((error) => of(login.error({ error }))),
                    tap((action) => {
                        if (action.type === login.success.type) {
                            console.warn('Routing from login to home');
                            this.router.navigate(['/']);
                        }
                    }) // Redirect after successful login
                )
            )
        )
    )

    register = createEffect(() =>
        this.actions.pipe(
            ofType(register.submit),
            switchMap(({ ...user }) =>
                this.authService.register(user).pipe(
                    delay(1000),
                    map((userWithToken) => {
                        this.cookieService.set('jwt_token', userWithToken.access);
                        return (login.success({ user: userWithToken }))
                    }

                    ),
                    catchError((error) => of(login.error({ error }))),
                    tap(() => {
                        console.warn('Routing from register to home')
                        this.router.navigate(['/'])
                    }) // Redirect after successful login
                )
            )
        )
    )

    logout = createEffect(() =>
        this.actions.pipe(
            ofType(logout.submit),
            map(() => {
                this.cookieService.delete('jwt_token'); // Clear JWT token on logout
                console.warn('JWT token deleted');
                this.router.navigate(['/login']); // Redirect to login page after logout
                return (logout.success())
            }),
            catchError((error) => of(logout.error({ error }))),

        ),
    )
}