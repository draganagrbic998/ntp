import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { Login } from 'src/app/models/login';
import { User } from 'src/app/models/user';
import { catchError, map } from 'rxjs/operators';
import { environment } from 'src/environments/environment';
import { Registration } from 'src/app/models/registration';
import { Profile } from 'src/app/models/profile';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor(
    private http: HttpClient
  ) { }

  login(login: Login): Observable<User>{
    return this.http.post<User>(`${environment.usersApi}/login`, login).pipe(
      catchError(() => of(null))
    );
  }

  register(register: Registration): Observable<boolean>{
    return this.http.post<null>(`${environment.usersApi}/register`, register).pipe(
      map(() => true),
      catchError(() => of(false))
    );
  }

  save(profile: Profile): Observable<User>{
    return this.http.post<User>(`${environment.usersApi}/update-profile`, profile).pipe(
      catchError(() => of(null))
    );
  }

}
