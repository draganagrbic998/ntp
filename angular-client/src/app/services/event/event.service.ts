import { HttpClient, HttpParams, HttpResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, of, Subject } from 'rxjs';
import { catchError, map } from 'rxjs/operators';
import { SMALL_PAGE_SIZE } from 'src/app/constants/pagination';
import { Event } from 'src/app/models/event';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class EventService {

  constructor(
    private http: HttpClient
  ) { }

  private refreshData: Subject<null> = new Subject();
  refreshData$: Observable<null> = this.refreshData.asObservable();
  selectedEvent: Event;

  getAll(page: number, productId: number): Observable<HttpResponse<Event[]>>{
    const params = new HttpParams().set('page', page + '').set('size', SMALL_PAGE_SIZE + '').set('product', productId + '');
    return this.http.get<Event[]>(`${environment.eventsApi}`, {observe: 'response', params}).pipe(
      catchError(() => of(null))
    );
  }

  save(event: Event): Observable<Event>{
    if (!event.ID){
      return this.http.post<Event>(environment.eventsApi, event).pipe(
        catchError(() => of(null))
      );
    }
    return this.http.put<Event>(`${environment.eventsApi}/${event.ID}`, event).pipe(
      catchError(() => of(null))
    );
  }

  delete(id: number): Observable<boolean>{
    return this.http.delete<null>(`${environment.eventsApi}/${id}`).pipe(
      map(() => true),
      catchError(() => of(false))
    );
  }

  announceRefreshData(): void{
    this.refreshData.next();
  }

}
