import { createActionGroup, emptyProps, props } from "@ngrx/store";
import { UserWithToken, UserLoginRequest } from "../../models/user.model";

export const login = createActionGroup({
    source: 'user login',
    events: {
        submit: props<{ user: UserLoginRequest }>(),
        success: props<{ user: UserWithToken }>(),
        error: props<{ error: any }>()
    }
})

export const register = createActionGroup({
    source: 'user register',
    events: {
        submit: props<{ username: string, password1: string, password2: string }>()
    }
})

export const logout = createActionGroup({
    source: 'user logout',
    events: {
        submit: emptyProps(),
        success: emptyProps(),
        error: props<{ error: any }>()
    }
})