import { inject, Injectable } from "@angular/core";
import { Actions, createEffect, ofType } from "@ngrx/effects";
import { ApiService } from "../../services/api.service";
import { fetchItems, deleteItem, addItem, updateItem } from "./items.actions";
import { catchError, delay, exhaustMap, map, mergeMap, of, switchMap, tap } from "rxjs";
import { Item } from "../../models/todo.model";

@Injectable()
export class ItemsEffects {
    private readonly actions = inject(Actions);
    private readonly itemsService = inject(ApiService);

    getItems = createEffect(() =>
        this.actions.pipe(
            ofType(fetchItems.submit),
            switchMap(() =>
                this.itemsService.getItems().pipe(
                    delay(2000),
                    // On success, dispatch the fetchItems.success action
                    tap(items => console.log('Fetched items:', items)),
                    map((items: Item[]) => fetchItems.success({ items: items })),

                    // On error, dispatch the submitItemFailure action
                    catchError((error) => of(fetchItems.error({ error })))
                )
            )
        )
    )

    deleteItem = createEffect(() =>
        this.actions.pipe(
            ofType(deleteItem.submit),
            exhaustMap(({ id }) => // Destructure to get `id` from action
                this.itemsService.deleteItem(id).pipe(
                    delay(2000), // Optional delay to simulate loading or timing

                    // On success, dispatch `deleteItem.success` with the deleted item's ID
                    map((id) => deleteItem.success({ id })),

                    // On error, dispatch `fetchItems.error` with error details
                    catchError((error) => of(deleteItem.error({ error })))
                )
            )
        )
    );

    addItem = createEffect(() =>
        this.actions.pipe(
            ofType(addItem.submit),
            exhaustMap(({ item }) => // Destructure to get `id` from action
                this.itemsService.addItem(item).pipe(
                    delay(2000), // Optional delay to simulate loading or timing

                    // On success, dispatch `deleteItem.success` with the deleted item's ID
                    map((id) => addItem.success({ item })),

                    // On error, dispatch `fetchItems.error` with error details
                    catchError((error) => of(addItem.error({ error })))
                )
            )
        )
    );

    updateItem = createEffect(() =>
        this.actions.pipe(
            ofType(updateItem.submit),
            exhaustMap(({ item }) => // Destructure to get `id` from action
                this.itemsService.addItem(item).pipe(
                    delay(2000), // Optional delay to simulate loading or timing

                    // On success, dispatch `deleteItem.success` with the deleted item's ID
                    map((id) => updateItem.success({ item })),

                    // On error, dispatch `fetchItems.error` with error details
                    catchError((error) => of(updateItem.error({ error })))
                )
            )
        )
    );
}