import { inject, Injectable } from "@angular/core";
import { Actions, createEffect, ofType } from "@ngrx/effects";
import { ApiService } from "../../services/api.service";
import { AuthService } from "../../services/auth.service";
import { login } from "./user.actions";
import { catchError, delay, map, of, switchMap, tap } from "rxjs";

@Injectable()
export class UserEffects {
    private readonly actions = inject(Actions);
    private readonly authService = inject(AuthService);

    login = createEffect(() =>
        this.actions.pipe(
            ofType(login.submit),
            switchMap(({ user }) =>
                this.authService.login(user).pipe(
                    delay(1000),
                    map((userWithToken) =>
                        login.success({ user: userWithToken })
                    ),
                    catchError((error) => of(login.error({ error })))
                )
            )
        )
    )
}