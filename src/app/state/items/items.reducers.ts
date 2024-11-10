import { createReducer, on } from "@ngrx/store";
import { itemsInitialState } from "./items.state";
import { fetchItems, deleteItem, addItem, fetchItem } from "./items.actions";

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

    on(fetchItem.submit, (state) => ({
        ...state,
        selectedItem: undefined,
        loading: true
    })),
    on(fetchItem.success, (state, { item: item }) => ({
        ...state,
        selectedItem: item,
        loading: false
    })),
    on(fetchItem.error, (state, { error }) => ({
        ...state,
        error: error,
        loading: false
    })),

    on(addItem.submit, (state) => ({
        ...state,
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

    on(deleteItem.submit, (state) => ({
        ...state,
        loading: true
    })),
    on(deleteItem.success, (state, { id }) => ({
        ...state,
        items: [...state.items.filter((item) => item.id != Number(id))],
        loading: false
    })),
    on(deleteItem.error, (state, { error }) => ({
        ...state,
        error: error,
        loading: false
    })),
)