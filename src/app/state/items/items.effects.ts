import { inject, Injectable } from "@angular/core";
import { Actions, createEffect, ofType } from "@ngrx/effects";
import { ApiService } from "../../services/api.service";
import { fetchItems } from "./items.actions";
import { catchError, map, mergeMap, of } from "rxjs";
import { Item } from "../../models/todo.model";

@Injectable()
export class ItemsEffects {
    private readonly actions = inject(Actions);
    private readonly itemsService = inject(ApiService);


    getItems$ = createEffect(() =>
        this.actions.pipe(
            ofType(fetchItems.submit),
            mergeMap(() =>
                this.itemsService.getData().pipe(
                    // On success, dispatch the fetchItems.success action
                    map((items: Item[]) => fetchItems.success({ items: items })),

                    // On error, dispatch the submitItemFailure action
                    catchError((error) => of(fetchItems.error({ error })))
                )
            )
        )
    )
}