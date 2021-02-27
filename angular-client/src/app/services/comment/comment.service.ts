import { HttpClient, HttpParams, HttpResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, of, Subject } from 'rxjs';
import { catchError, map } from 'rxjs/operators';
import { SMALL_PAGE_SIZE } from 'src/app/constants/pagination';
import { Comment } from 'src/app/models/comment';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class CommentService {

  constructor(
    private http: HttpClient
  ) { }

  private refreshData: Subject<null> = new Subject();
  refreshData$: Observable<null> = this.refreshData.asObservable();

  getAll(page: number, productId: number): Observable<HttpResponse<Comment[]>>{
    const params = new HttpParams().set('page', page + '').set('size', SMALL_PAGE_SIZE + '').set('product', productId + '');
    return this.http.get<Comment[]>(environment.commentsApi, {observe: 'response', params}).pipe(
      catchError(() => of(null))
    );
  }

  save(comment: Comment): Observable<Comment>{
    return this.http.post<Comment>(environment.commentsApi, comment).pipe(
      catchError(() => of(null))
    );
  }

  delete(id: number): Observable<boolean>{
    return this.http.delete<null>(`${environment.commentsApi}/${id}`).pipe(
      map(() => true),
      catchError(() => of(false))
    );
  }

  like(id: number, dislike: boolean): Observable<boolean>{
    return this.http.get<boolean>(`${environment.commentsApi}/${id}/like?dislike=${dislike}`).pipe(
      map(() => true),
      catchError(() => of(false))
    );
  }

  replies(id: number): Observable<Comment[]>{
    return this.http.get<Comment[]>(`${environment.commentsApi}/${id}/replies`).pipe(
      catchError(() => of([]))
    );
  }

  announceRefreshData(): void{
    this.refreshData.next();
  }

}
