import { HttpClient, HttpParams, HttpResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, of, Subject } from 'rxjs';
import { catchError, map } from 'rxjs/operators';
import { SMALL_PAGE_SIZE } from 'src/app/constants/pagination';
import { Advertisement } from 'src/app/models/ad';
import { environment } from 'src/environments/environment';
import { addSyntheticLeadingComment } from 'typescript';

@Injectable({
  providedIn: 'root'
})
export class AdService {

  constructor(
    private http: HttpClient
  ) { }

  private refreshData: Subject<null> = new Subject();
  refreshData$: Observable<null> = this.refreshData.asObservable();
  private searchData: Subject<string> = new Subject();
  searchData$: Observable<string> = this.searchData.asObservable();
  private listToggle: Subject<null> = new Subject();
  listToggle$ = this.listToggle.asObservable();
  selectedAd: Advertisement;
  myProducts = false;

  getAll(page: number, search: string): Observable<HttpResponse<Advertisement[]>>{
    const params = new HttpParams().set('page', page + '').set('size', SMALL_PAGE_SIZE + '').set('search', search);
    const url = this.myProducts ? `${environment.adsApi}-my` : environment.adsApi;
    return this.http.get<Advertisement[]>(url, {observe: 'response', params}).pipe(
      catchError(() => of(null))
    );
  }

  getOne(id: number): Observable<Advertisement>{
    return this.http.get<Advertisement>(`${environment.adsApi}/${id}`).pipe(
      catchError(() => of(null))
    );
  }

  save(ad: Advertisement): Observable<Advertisement>{
    if (!ad.ID){
      return this.http.post<Advertisement>(environment.adsApi, ad).pipe(
        catchError(() => of(null))
      );
    }
    return this.http.put<Advertisement>(`${environment.adsApi}/${ad.ID}`, ad).pipe(
      catchError(() => of(null))
    );
  }

  delete(id: number): Observable<boolean>{
    return this.http.delete<null>(`${environment.adsApi}/${id}`).pipe(
      map(() => true),
      catchError(() => of(false))
    );
  }

  announceRefreshData(): void{
    this.refreshData.next();
  }

  announceSearchData(search: string): void{
    this.searchData.next(search);
  }

  announceListToggle(): void{
    this.myProducts = !this.myProducts;
    this.listToggle.next();
  }

}
