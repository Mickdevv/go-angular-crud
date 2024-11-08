import { createActionGroup, emptyProps, props } from "@ngrx/store";
import { Item } from "../../models/todo.model";

export const fetchItems = createActionGroup({
    source: 'Items',
    events: {
        submit: emptyProps(),
        success: props<{ items: Item[] }>(),
        error: props<{ error: any }>(),
    }
})

export const deleteItem = createActionGroup({
    source: 'Delete item',
    events: {
        submit: props<{ id: number }>(),
        success: props<{ id: number }>(),
        error: props<{ error: any }>(),
    }
})

// export const addItem = createActionGroup({
//     source: 'Items',
//     events: {
//         submit: props<{ item: Item }>(),
//         success: props<{ item: Item }>(),
//         error: props<{ error: any }>(),
//     }
// })

// export const updateItem = createActionGroup({
//     source: 'Items',
//     events: {
//         submit: props<{ item: Item }>(),
//         success: props<{ item: Item }>(),
//         error: props<{ error: any }>(),
//     }
// })

// export const removeItem = createActionGroup({
//     source: 'Items',
//     events: {
//         submit: props<{ item: Item }>(),
//         success: props<{ item: Item }>(),
//         error: props<{ error: any }>(),
//     }
// })