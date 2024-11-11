import { createReducer, on } from "@ngrx/store";
import { userInitialState } from "./user.state";
import { state } from "@angular/animations";
import { login, logout } from "./user.actions";

export const userReducer = createReducer(
    userInitialState,

    on(login.submit, (state) => ({
        ...state,
        username: '',
        access: '',
        refresh: '',
        loading: true,
        success: false
    })),
    on(login.success, (state, { user }) => ({
        ...state,
        username: user.username,
        access: user.access,
        refresh: user.refresh,
        success: true,
        loading: false
    })),
    on(login.error, (state, { error }) => ({
        ...state,
        username: '',
        access: '',
        refresh: '',
        loading: false,
        success: false,
        error: error
    })),

    on(logout.submit, (state) => ({
        ...state,
        userInitialState
    })),
    on(logout.resetSuccess, (state) => ({
        ...state,
        success: false
    }))
)