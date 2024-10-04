import { createSelector } from "@ngrx/store";
import { AppState } from "../app.state";
import { create } from "domain";

export const selectItemsState = (state: AppState) =>
    state.toDoList

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