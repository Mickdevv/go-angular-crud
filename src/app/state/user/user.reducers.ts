import { createReducer, on } from "@ngrx/store";
import { userInitialState } from "./user.state";
import { state } from "@angular/animations";
import { login } from "./user.actions";

export const userReducer = createReducer(
    userInitialState,

    on(login.submit, (state) => ({
        ...state,
        username: '',
        access: '',
        refresh: '',
        loading: true
    })),
    on(login.success, (state, { user }) => ({
        ...state,
        username: user.username,
        access: user.access,
        refresh: user.refresh,
        loading: false
    })),
    on(login.error, (state, { error }) => ({
        ...state,
        username: '',
        access: '',
        refresh: '',
        loading: false,
        error: error
    })),
)