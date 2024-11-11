import { createSelector } from "@ngrx/store";
import { AppState } from "../app.state";
import { UserState } from "./user.state";

export const selectUserState = (state: AppState) =>
    state.user

export const selectUsername = createSelector(
    selectUserState,
    (s) => s.username
)
export const selectUserToken = createSelector(
    selectUserState,
    (s) => { s.access, s.refresh, s.username }
)
export const selectUserLoading = createSelector(
    selectUserState,
    (s) => s.loading
)
export const selectUserError = createSelector(
    selectUserState,
    (s) => s.error
)