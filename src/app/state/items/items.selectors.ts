import { createSelector } from "@ngrx/store";
import { AppState } from "../app.state";

export const selectItemsState = (state: AppState) =>
    state.items

export const selectItems = createSelector(
    selectItemsState,
    (s) => s.items
)

export const selectItemsLoading = createSelector(
    selectItemsState,
    (s) => s.loading
)
export const selectItemsError = createSelector(
    selectItemsState,
    (s) => s.error
)