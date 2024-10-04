import { createReducer, on } from "@ngrx/store";
import { itemsInitialState } from "./items.state";
import { addItem, fetchItems } from "./items.actions";

export const itemsReducer = createReducer(
    itemsInitialState,
    on(fetchItems.submit, (state) => ({
        ...state,
        items: [],
        loading: true
    })),
    on(fetchItems.success, (state, { items }) => ({
        ...state,
        items: items,
        loading: false
    })),
    on(fetchItems.error, (state, { error }) => ({
        ...state,
        error: error,
        loading: false
    })),

    on(addItem.submit, (state, { item }) => ({
        ...state,
        items: [...state.items, item],
        loading: true
    })),
    on(addItem.success, (state, { item }) => ({
        ...state,
        items: [...state.items, item],
        loading: false
    })),
    on(addItem.error, (state, { error }) => ({
        ...state,
        error: error,
        loading: false
    })),
)